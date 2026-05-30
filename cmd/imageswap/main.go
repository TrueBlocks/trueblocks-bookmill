package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/TrueBlocks/trueblocks-art/packages/cli"
	"github.com/TrueBlocks/trueblocks-art/packages/docxzip"
)

var version = "dev"

func main() {
	app := cli.App{
		Name:        "imageswap",
		Description: "Swap [[IMG:...]] / [[R:...]] tag lines in .docx files with corresponding PNGs from an images directory.",
		Version:     version,
		ArgsUsage:   "<file.docx> [file.docx ...]",
		Flags: []cli.FlagDef{
			{Name: "images", Help: "path to images directory (contains slug subdirs); default: <cwd>/projects/math-books/images"},
			{Name: "slug", Help: "override slug (default: derived from docx filename)"},
			{Name: "dry-run", Help: "report what would change without modifying files", Default: false},
			{Name: "crop-borders", Help: "crop red borders from all PNGs in this directory (no docx needed)"},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

func run(c *cli.Context) error {
	imagesDir := c.String("images")
	slugFlag := c.String("slug")
	dryRun := c.Bool("dry-run")
	cropBorders := c.String("crop-borders")

	if cropBorders != "" {
		cropRedBordersInDir(cropBorders, dryRun)
		return nil
	}

	if len(c.Args) < 1 {
		return cli.NewUsageError(fmt.Errorf("at least one .docx file required (or use --crop-borders DIR)"))
	}

	if imagesDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("getting cwd: %w", err)
		}
		imagesDir = filepath.Join(cwd, "projects", "math-books", "images")
	}

	for _, docxPath := range c.Args {
		if err := swapImages(docxPath, imagesDir, slugFlag, dryRun); err != nil {
			log.Printf("ERROR %s: %v", docxPath, err)
		} else {
			log.Printf("OK %s", docxPath)
		}
	}
	return nil
}

func cropRedBordersInDir(dir string, dryRun bool) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("reading directory %s: %v", dir, err)
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(strings.ToLower(e.Name()), ".png") {
			continue
		}
		pngPath := filepath.Join(dir, e.Name())
		cropped, err := processRedBorder(pngPath, dryRun)
		if err != nil {
			log.Printf("  WARNING %s: %v", e.Name(), err)
		} else if !cropped {
			log.Printf("  skip %s (no red border)", e.Name())
		}
	}
}

// reTagLine matches both [[IMG:filename.png|caption]] and [[R:filename.png]]
var reTagLine = regexp.MustCompile(`\[\[(?:IMG:|R:)([^\]|]+?)(?:\|[^\]]*?)?\]\]`)

func swapImages(docxPath, imagesDir, slugOverride string, dryRun bool) error {
	slug := slugOverride
	if slug == "" {
		slug = strings.TrimSuffix(filepath.Base(docxPath), ".docx")
	}
	slugImgDir := filepath.Join(imagesDir, slug)

	if _, err := os.Stat(slugImgDir); os.IsNotExist(err) {
		return fmt.Errorf("no images directory for slug %q at %s", slug, slugImgDir)
	}

	files, err := docxzip.ReadFiles(docxPath, docxzip.DocumentXML, docxzip.RelsXML, docxzip.ContentTypesXML)
	if err != nil {
		return fmt.Errorf("reading docx: %w", err)
	}

	docXML := string(files[docxzip.DocumentXML])
	relsXML := string(files[docxzip.RelsXML])
	contentTypes := string(files[docxzip.ContentTypesXML])

	// Collect existing media files to know what image numbers are taken
	entries, err := docxzip.ListEntries(docxPath)
	if err != nil {
		return fmt.Errorf("listing entries: %w", err)
	}
	existingMedia := map[string]bool{}
	for _, name := range entries {
		if strings.HasPrefix(name, "word/media/") {
			existingMedia[name] = true
		}
	}

	// Find next available rId and image number
	nextRID := findMaxRID(relsXML) + 1
	nextImgNum := findMaxImageNum(existingMedia) + 1

	// Find all tag paragraphs and process them
	tags := reTagLine.FindAllStringSubmatch(docXML, -1)
	if len(tags) == 0 {

	}

	// mediaFiles maps zip path -> local PNG path for files to add/replace
	mediaFiles := map[string]string{}
	newRels := []string{}
	modified := false

	for _, tag := range tags {
		filename := tag[1]
		pngPath := filepath.Join(slugImgDir, filename)

		// Prefer Urban sibling folder if it has the same file
		urbanDir := filepath.Join(filepath.Dir(slugImgDir), "Urban")
		urbanPath := filepath.Join(urbanDir, filename)
		if _, err := os.Stat(urbanPath); err == nil {
			log.Printf("  using Urban variant for %s", filename)
			pngPath = urbanPath
		}

		if _, err := os.Stat(pngPath); os.IsNotExist(err) {
			log.Printf("  WARNING: image not found: %s", pngPath)
			continue
		}

		if _, err := processRedBorder(pngPath, dryRun); err != nil {
			log.Printf("  WARNING: red border processing failed for %s: %v", filename, err)
		}

		// Normalize the tag paragraph to ImageTag style with vanish and [[R:...]] format
		docXML = normalizeTagParagraph(docXML, tag[0], filename)

		// Check if there's already a drawing paragraph right after the tag
		existingRID, existingDrawingPara := findExistingImageRID(docXML, filename)
		if existingRID != "" {
			// Image paragraph exists — find its media target and replace the bytes
			target := findRelTarget(relsXML, existingRID)
			if target != "" {
				zipPath := "word/" + target

				// Check if bytes would actually change
				newData, err := os.ReadFile(pngPath)
				if err != nil {
					log.Printf("  WARNING: cannot read %s: %v", pngPath, err)
					continue
				}
				existingData, err := docxzip.ReadFile(docxPath, zipPath)
				var bytesMatch bool
				if err == nil && len(existingData) == len(newData) {
					bytesMatch = true
					for i := range newData {
						if newData[i] != existingData[i] {
							bytesMatch = false
							break
						}
					}
				}

				// Check if dimensions would change
				cx, cy := pngDimensionsEMU(pngPath)
				oldDocXML := docXML
				docXML = updateExistingImageDimensions(docXML, existingDrawingPara, cx, cy)
				dimsChanged := docXML != oldDocXML

				if bytesMatch && !dimsChanged {
					continue
				}

				mediaFiles[zipPath] = pngPath
				log.Printf("  updating %s <- %s", zipPath, filename)
				modified = true
			}
		} else {
			// No image paragraph — insert one
			if dryRun {
				cx, cy := pngDimensionsEMU(pngPath)
				log.Printf("  would insert %s (cx=%d cy=%d)", filename, cx, cy)
				modified = true
				continue
			}

			rID := fmt.Sprintf("rId%d", nextRID)
			mediaName := fmt.Sprintf("image%d.png", nextImgNum)
			mediaZipPath := "word/media/" + mediaName
			nextRID++
			nextImgNum++

			// Read PNG dimensions for the drawing element
			cx, cy := pngDimensionsEMU(pngPath)

			imgParagraph := buildImageParagraph(rID, mediaName, filename, cx, cy, nextImgNum-1)
			docXML = insertParagraphAfterTag(docXML, filename, imgParagraph)

			mediaFiles[mediaZipPath] = pngPath
			newRels = append(newRels, fmt.Sprintf(
				`<Relationship Id="%s" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/image" Target="media/%s"/>`,
				rID, mediaName))

			log.Printf("  inserted %s <- %s", mediaZipPath, filename)
			modified = true
		}
	}

	if !modified {
		return nil
	}

	if dryRun {
		return nil
	}

	// Add new relationships
	if len(newRels) > 0 {
		insertion := strings.Join(newRels, "")
		relsXML = strings.Replace(relsXML, "</Relationships>", insertion+"</Relationships>", 1)
	}

	// Ensure PNG content type is registered
	if !strings.Contains(contentTypes, `Extension="png"`) {
		contentTypes = strings.Replace(contentTypes, "</Types>",
			`<Default Extension="png" ContentType="image/png"/></Types>`, 1)
	}

	// Guard: document.xml must not shrink significantly — that means text was lost.
	// Allow up to 1000 bytes of shrinkage to account for tag normalization ([[IMG:]] -> [[R:]])
	// and dimension updates which may reduce digit counts in numeric attributes.
	origDocLen := len(files[docxzip.DocumentXML])
	newDocLen := len(docXML)
	const shrinkThreshold = 1000
	if newDocLen < origDocLen-shrinkThreshold {
		return fmt.Errorf("document.xml shrank (%d -> %d bytes) — text was lost, aborting", origDocLen, newDocLen)
	}

	// Build overrides map for Rewrite
	overrides := map[string][]byte{
		docxzip.DocumentXML:     []byte(docXML),
		docxzip.RelsXML:         []byte(relsXML),
		docxzip.ContentTypesXML: []byte(contentTypes),
	}
	for zipPath, localPath := range mediaFiles {
		imgData, err := os.ReadFile(localPath)
		if err != nil {
			return fmt.Errorf("reading %s: %w", localPath, err)
		}
		overrides[zipPath] = imgData
	}

	backupDir := filepath.Join(os.Getenv("HOME"), ".local", "share", "trueblocks-backups")
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return fmt.Errorf("creating backup dir: %w", err)
	}
	bakPath := filepath.Join(backupDir, filepath.Base(docxPath)+".bak")
	srcData, err := os.ReadFile(docxPath)
	if err != nil {
		return fmt.Errorf("reading docx for backup: %w", err)
	}
	if err := os.WriteFile(bakPath, srcData, 0644); err != nil {
		return fmt.Errorf("writing backup: %w", err)
	}

	if err := docxzip.Rewrite(docxPath, docxPath, overrides); err != nil {
		os.Rename(bakPath, docxPath)
		return err
	}

	log.Printf("  backup saved: %s", bakPath)
	return nil
}

// normalizeTagParagraph converts any tag paragraph to ImageTag style with vanish
// and normalizes [[IMG:file.png|caption]] to [[R:file.png]]
func normalizeTagParagraph(docXML, fullTag, filename string) string {
	// Find the <w:p> containing this tag
	tagIdx := strings.Index(docXML, fullTag)
	if tagIdx < 0 {
		return docXML
	}

	// Find enclosing paragraph — take whichever opener is closest to the tag
	pStartPlain := strings.LastIndex(docXML[:tagIdx], "<w:p>")
	pStartAttr := strings.LastIndex(docXML[:tagIdx], "<w:p ")
	pStart := pStartPlain
	if pStartAttr > pStart {
		pStart = pStartAttr
	}
	if pStart < 0 {
		return docXML
	}
	pEnd := strings.Index(docXML[tagIdx:], "</w:p>")
	if pEnd < 0 {
		return docXML
	}
	pEnd = tagIdx + pEnd + len("</w:p>")

	oldPara := docXML[pStart:pEnd]

	// Build the normalized paragraph
	newPara := fmt.Sprintf(
		`<w:p><w:pPr><w:pStyle w:val="ImageTag"/></w:pPr><w:r><w:rPr><w:vanish/></w:rPr><w:t>[[R:%s]]</w:t></w:r></w:p>`,
		filename)

	return strings.Replace(docXML, oldPara, newPara, 1)
}

// findExistingImageRID checks if the paragraph immediately after the tag
// for the given filename contains a drawing with an embedded image.
// Returns the rId and the full drawing paragraph text (so callers can locate
// the exact paragraph even when the same rId appears multiple times).
func findExistingImageRID(docXML, filename string) (rID, drawingPara string) {
	tagText := "[[R:" + filename + "]]"
	tagIdx := strings.Index(docXML, tagText)
	if tagIdx < 0 {
		return "", ""
	}

	// Find end of current paragraph
	pEnd := strings.Index(docXML[tagIdx:], "</w:p>")
	if pEnd < 0 {
		return "", ""
	}
	afterTag := tagIdx + pEnd + len("</w:p>")

	// Look at next paragraph
	rest := docXML[afterTag:]
	rest = strings.TrimSpace(rest)

	// Check if next paragraph has a drawing
	nextPEnd := strings.Index(rest, "</w:p>")
	if nextPEnd < 0 {
		return "", ""
	}
	nextPara := rest[:nextPEnd+len("</w:p>")]

	if !strings.Contains(nextPara, "<w:drawing>") {
		return "", ""
	}

	blipRe := regexp.MustCompile(`r:embed="(rId\d+)"`)
	m := blipRe.FindStringSubmatch(nextPara)
	if m == nil {
		return "", ""
	}
	return m[1], nextPara
}

// updateExistingImageDimensions updates the wp:extent and a:ext cx/cy values within
// the given drawing paragraph and ensures the paragraph is centered.
// drawingPara must be the exact paragraph text as it appears in docXML.
func updateExistingImageDimensions(docXML, drawingPara string, cx, cy int64) string {
	// Locate the specific drawing paragraph by its exact text content.
	paraStart := strings.Index(docXML, drawingPara)
	if paraStart < 0 {
		return docXML
	}

	// Find the <w:drawing>...</w:drawing> within that paragraph.
	drawOffset := strings.Index(drawingPara, "<w:drawing>")
	if drawOffset < 0 {
		return docXML
	}
	drawStart := paraStart + drawOffset
	drawEnd := strings.Index(docXML[drawStart:], "</w:drawing>")
	if drawEnd < 0 {
		return docXML
	}
	drawEnd = drawStart + drawEnd + len("</w:drawing>")

	oldDrawing := docXML[drawStart:drawEnd]

	// Replace wp:extent cx/cy
	extentRe := regexp.MustCompile(`(<wp:extent\s+)cx="[0-9]+"(\s+)cy="[0-9]+"`)
	newDrawing := extentRe.ReplaceAllString(oldDrawing,
		fmt.Sprintf(`${1}cx="%d"${2}cy="%d"`, cx, cy))

	// Replace a:ext cx/cy (inside pic:spPr > a:xfrm)
	aExtRe := regexp.MustCompile(`(<a:ext\s+)cx="[0-9]+"(\s+)cy="[0-9]+"`)
	newDrawing = aExtRe.ReplaceAllString(newDrawing,
		fmt.Sprintf(`${1}cx="%d"${2}cy="%d"`, cx, cy))

	docXML = strings.Replace(docXML, oldDrawing, newDrawing, 1)

	// Ensure the enclosing paragraph is centered.
	// Re-locate the paragraph start from paraStart (anchored to the right paragraph).
	pStart := paraStart
	pEnd := strings.Index(docXML[pStart:], "</w:p>")
	if pEnd < 0 {
		return docXML
	}
	pEnd = pStart + pEnd + len("</w:p>")
	oldPara := docXML[pStart:pEnd]

	if strings.Contains(oldPara, `w:val="center"`) {
		return docXML
	}

	var newPara string
	if strings.Contains(oldPara, "<w:pPr>") {
		newPara = strings.Replace(oldPara, "<w:pPr>", `<w:pPr><w:jc w:val="center"/>`, 1)
	} else if idx := strings.Index(oldPara, "<w:pPr "); idx >= 0 {
		closeIdx := strings.Index(oldPara[idx:], ">")
		insertAt := idx + closeIdx + 1
		newPara = oldPara[:insertAt] + `<w:jc w:val="center"/>` + oldPara[insertAt:]
	} else {
		firstClose := strings.Index(oldPara, ">")
		newPara = oldPara[:firstClose+1] + `<w:pPr><w:jc w:val="center"/></w:pPr>` + oldPara[firstClose+1:]
	}

	return strings.Replace(docXML, oldPara, newPara, 1)
}

// findRelTarget finds the Target for a given rId in the rels XML
func findRelTarget(relsXML, rID string) string {
	pattern := `Id="` + regexp.QuoteMeta(rID) + `"[^>]*Target="([^"]+)"`
	re := regexp.MustCompile(pattern)
	m := re.FindStringSubmatch(relsXML)
	if m == nil {
		return ""
	}
	return m[1]
}

// insertParagraphAfterTag inserts imgParagraph right after the tag paragraph for filename
func insertParagraphAfterTag(docXML, filename, imgParagraph string) string {
	tagText := "[[R:" + filename + "]]"
	tagIdx := strings.Index(docXML, tagText)
	if tagIdx < 0 {
		return docXML
	}

	// Find end of the tag paragraph
	pEnd := strings.Index(docXML[tagIdx:], "</w:p>")
	if pEnd < 0 {
		return docXML
	}
	insertAt := tagIdx + pEnd + len("</w:p>")

	return docXML[:insertAt] + imgParagraph + docXML[insertAt:]
}

// buildImageParagraph creates the XML for a centered paragraph with an inline image
func buildImageParagraph(rID, mediaName, filename string, cx, cy int64, docPrID int) string {
	return fmt.Sprintf(
		`<w:p><w:pPr><w:jc w:val="center"/></w:pPr><w:r><w:drawing>`+
			`<wp:inline distT="0" distB="0" distL="0" distR="0">`+
			`<wp:extent cx="%d" cy="%d"/>`+
			`<wp:docPr id="%d" name="Image %d" descr="%s"/>`+
			`<a:graphic><a:graphicData uri="http://schemas.openxmlformats.org/drawingml/2006/picture">`+
			`<pic:pic xmlns:pic="http://schemas.openxmlformats.org/drawingml/2006/picture">`+
			`<pic:nvPicPr><pic:cNvPr id="%d" name="%s"/><pic:cNvPicPr/></pic:nvPicPr>`+
			`<pic:blipFill><a:blip r:embed="%s"/><a:stretch><a:fillRect/></a:stretch></pic:blipFill>`+
			`<pic:spPr><a:xfrm><a:off x="0" y="0"/><a:ext cx="%d" cy="%d"/></a:xfrm>`+
			`<a:prstGeom prst="rect"><a:avLst/></a:prstGeom></pic:spPr>`+
			`</pic:pic></a:graphicData></a:graphic></wp:inline></w:drawing></w:r></w:p>`,
		cx, cy, docPrID, docPrID, filename,
		docPrID, mediaName, rID, cx, cy)
}

// woodcutNames is the set of PNG filenames that must always render at exactly 4.5" x 4.5".
var woodcutNames = map[string]bool{
	"woodcut-open.png":  true,
	"woodcut-close.png": true,
}

// pngDimensionsEMU reads a PNG and returns width/height in EMUs (English Metric Units).
// Woodcut images are forced to exactly 4.5" x 4.5" regardless of pixel dimensions.
// All other images are capped at max width 4.5" and max height 6" for 6x9 trade paperback sizing.
func pngDimensionsEMU(pngPath string) (cx, cy int64) {
	const emuPerInch = 914400
	const woodcutInches = 4.5
	const maxWidthInches = 4.5
	const maxHeightInches = 6.0

	// Woodcuts are always exactly 4.5" x 4.5"
	basename := filepath.Base(pngPath)
	if woodcutNames[basename] {
		const emu45 = int64(woodcutInches * emuPerInch)
		return emu45, emu45
	}

	f, err := os.Open(pngPath)
	if err != nil {
		return 4114800, 2743200 // 4.5" x 3" default
	}
	defer f.Close()

	cfg, err := png.DecodeConfig(f)
	if err != nil {
		return 4114800, 2743200
	}

	w := float64(cfg.Width)
	h := float64(cfg.Height)

	wInch := w / 300.0
	hInch := h / 300.0

	if wInch > maxWidthInches {
		scale := maxWidthInches / wInch
		wInch *= scale
		hInch *= scale
	}
	if hInch > maxHeightInches {
		scale := maxHeightInches / hInch
		wInch *= scale
		hInch *= scale
	}

	cx = int64(wInch * emuPerInch)
	cy = int64(hInch * emuPerInch)
	return cx, cy
}

// findMaxRID finds the highest rId number in the rels XML
func findMaxRID(relsXML string) int {
	re := regexp.MustCompile(`Id="rId(\d+)"`)
	matches := re.FindAllStringSubmatch(relsXML, -1)
	max := 0
	for _, m := range matches {
		if n, err := strconv.Atoi(m[1]); err == nil && n > max {
			max = n
		}
	}
	return max
}

// findMaxImageNum finds the highest image number in existing media files
func findMaxImageNum(existing map[string]bool) int {
	re := regexp.MustCompile(`word/media/image(\d+)\.png`)
	max := 0
	for path := range existing {
		m := re.FindStringSubmatch(path)
		if m != nil {
			if n, err := strconv.Atoi(m[1]); err == nil && n > max {
				max = n
			}
		}
	}
	return max
}

func isRed(r, g, b uint32) bool {
	return r >= 50000 && g <= 20000 && b <= 15000
}

func processRedBorder(pngPath string, dryRun bool) (bool, error) {
	f, err := os.Open(pngPath)
	if err != nil {
		return false, err
	}

	img, err := png.Decode(f)
	f.Close()
	if err != nil {
		return false, err
	}

	bounds := img.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()

	type redRun struct {
		start, thickness int
	}

	// findFirstRedRun scans a sequence of pixels and returns the first consecutive
	// red run. coordFn(i) returns the (x,y) for step i. Returns (-1,0) if none.
	findFirstRedRun := func(n int, coordFn func(int) (int, int)) redRun {
		run := redRun{-1, 0}
		for i := 0; i < n; i++ {
			x, y := coordFn(i)
			r, g, b, _ := img.At(x, y).RGBA()
			if isRed(r, g, b) {
				if run.start < 0 {
					run.start = i
				}
				run.thickness++
			} else if run.start >= 0 {
				break
			}
		}
		return run
	}

	// Step 1: Scan each row from top to find the first row with a horizontal red run
	topRun := redRun{-1, 0}
	topRowRedMinX, topRowRedMaxX := -1, -1
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		run := findFirstRedRun(w, func(i int) (int, int) { return bounds.Min.X + i, y })
		if run.thickness >= 2 {
			topRun = redRun{start: y, thickness: 0}
			topRowRedMinX = bounds.Min.X + run.start
			topRowRedMaxX = bounds.Min.X + run.start + run.thickness - 1
			// Count how many consecutive rows have red at this same x
			for yy := y; yy < bounds.Max.Y; yy++ {
				r, g, b, _ := img.At(topRowRedMinX+run.thickness/2, yy).RGBA()
				if isRed(r, g, b) {
					topRun.thickness++
				} else {
					break
				}
			}
			break
		}
	}
	if topRun.thickness < 2 || topRowRedMinX < 0 {
		return false, nil
	}

	// Use the midpoint of the top red line's x-range to scan vertically
	boxMidX := (topRowRedMinX + topRowRedMaxX) / 2

	// Step 2: Scan from bottom up at boxMidX to find bottom red line
	bottomRun := findFirstRedRun(h, func(i int) (int, int) { return boxMidX, bounds.Max.Y - 1 - i })
	if bottomRun.thickness < 2 {
		return false, nil
	}
	bottomY := bounds.Max.Y - 1 - bottomRun.start

	// Use the midpoint of top and bottom to scan horizontally
	boxMidY := (topRun.start + bottomY) / 2

	// Step 3: Scan from left at boxMidY to find left red line
	leftRun := findFirstRedRun(w, func(i int) (int, int) { return bounds.Min.X + i, boxMidY })
	if leftRun.thickness < 2 {
		return false, nil
	}
	leftX := bounds.Min.X + leftRun.start

	// Step 4: Scan from right at boxMidY to find right red line
	rightRun := findFirstRedRun(w, func(i int) (int, int) { return bounds.Max.X - 1 - i, boxMidY })
	if rightRun.thickness < 2 {
		return false, nil
	}
	rightX := bounds.Max.X - 1 - rightRun.start

	pad := 4
	cropLeft := leftX + leftRun.thickness + pad
	cropTop := topRun.start + topRun.thickness + pad
	cropRight := rightX - pad
	cropBottom := bottomY - pad

	cropRect := image.Rect(cropLeft, cropTop, cropRight+1, cropBottom+1)
	if cropRect.Empty() {
		return false, nil
	}

	if cropRect.Dx() < w/4 || cropRect.Dy() < h/4 {
		return false, nil
	}

	if dryRun {
		log.Printf("  would crop red border from source (%dx%d -> %dx%d) [L=%d R=%d T=%d B=%d]",
			bounds.Dx(), bounds.Dy(), cropRect.Dx(), cropRect.Dy(), leftRun.thickness, rightRun.thickness, topRun.thickness, bottomRun.thickness)
		return true, nil
	}

	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	si, ok := img.(subImager)
	if !ok {
		return false, fmt.Errorf("image type does not support SubImage")
	}

	croppedImg := si.SubImage(cropRect)

	var buf bytes.Buffer
	if err := png.Encode(&buf, croppedImg); err != nil {
		return false, fmt.Errorf("encoding cropped PNG: %w", err)
	}

	if err := os.WriteFile(pngPath, buf.Bytes(), 0644); err != nil {
		return false, fmt.Errorf("writing cropped PNG: %w", err)
	}

	log.Printf("  cropped red border from source (%dx%d -> %dx%d) [L=%d R=%d T=%d B=%d]",
		bounds.Dx(), bounds.Dy(), cropRect.Dx(), cropRect.Dy(), leftRun.thickness, rightRun.thickness, topRun.thickness, bottomRun.thickness)

	return true, nil
}
