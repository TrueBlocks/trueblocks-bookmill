package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/TrueBlocks/trueblocks-art/packages/docxzip"
)

func main() {
	imagesDir := flag.String("images", "", "path to images directory (contains slug subdirs)")
	slugFlag := flag.String("slug", "", "override slug (default: derived from docx filename)")
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalf("Usage: imageswap [--images DIR] <file.docx> [file2.docx ...]")
	}

	if *imagesDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("getting cwd: %v", err)
		}
		*imagesDir = filepath.Join(cwd, "projects", "math-books", "images")
	}

	for _, docxPath := range flag.Args() {
		if err := swapImages(docxPath, *imagesDir, *slugFlag); err != nil {
			log.Printf("ERROR %s: %v", docxPath, err)
		} else {
			log.Printf("OK %s", docxPath)
		}
	}
}

// reTagLine matches both [[IMG:filename.png|caption]] and [[R:filename.png]]
var reTagLine = regexp.MustCompile(`\[\[(?:IMG:|R:)([^\]|]+?)(?:\|[^\]]*?)?\]\]`)

func swapImages(docxPath, imagesDir, slugOverride string) error {
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
		log.Printf("  no image tags found in %s", docxPath)
		return nil
	}

	// mediaFiles maps zip path -> local PNG path for files to add/replace
	mediaFiles := map[string]string{}
	newRels := []string{}
	modified := false

	for _, tag := range tags {
		filename := tag[1]
		pngPath := filepath.Join(slugImgDir, filename)
		if _, err := os.Stat(pngPath); os.IsNotExist(err) {
			log.Printf("  WARNING: image not found: %s", pngPath)
			continue
		}

		// Normalize the tag paragraph to ImageTag style with vanish and [[R:...]] format
		docXML = normalizeTagParagraph(docXML, tag[0], filename)

		// Check if there's already a drawing paragraph right after the tag
		existingRID := findExistingImageRID(docXML, filename)
		if existingRID != "" {
			// Image paragraph exists — find its media target and replace the bytes
			target := findRelTarget(relsXML, existingRID)
			if target != "" {
				zipPath := "word/" + target
				mediaFiles[zipPath] = pngPath

				// Also update the image dimensions in the existing drawing XML
				cx, cy := pngDimensionsEMU(pngPath)
				docXML = updateExistingImageDimensions(docXML, existingRID, cx, cy)

				log.Printf("  updating %s <- %s", zipPath, filename)
				modified = true
			}
		} else {
			// No image paragraph — insert one
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
		log.Printf("  nothing to update in %s", docxPath)
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

	return docxzip.Rewrite(docxPath, docxPath, overrides)
}

// normalizeTagParagraph converts any tag paragraph to ImageTag style with vanish
// and normalizes [[IMG:file.png|caption]] to [[R:file.png]]
func normalizeTagParagraph(docXML, fullTag, filename string) string {
	// Find the <w:p> containing this tag
	tagIdx := strings.Index(docXML, fullTag)
	if tagIdx < 0 {
		return docXML
	}

	// Find enclosing paragraph
	pStart := strings.LastIndex(docXML[:tagIdx], "<w:p>")
	if pStart < 0 {
		pStart = strings.LastIndex(docXML[:tagIdx], "<w:p ")
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
func findExistingImageRID(docXML, filename string) string {
	tagText := "[[R:" + filename + "]]"
	tagIdx := strings.Index(docXML, tagText)
	if tagIdx < 0 {
		return ""
	}

	// Find end of current paragraph
	pEnd := strings.Index(docXML[tagIdx:], "</w:p>")
	if pEnd < 0 {
		return ""
	}
	afterTag := tagIdx + pEnd + len("</w:p>")

	// Look at next paragraph
	rest := docXML[afterTag:]
	rest = strings.TrimSpace(rest)

	// Check if next paragraph has a drawing
	nextPEnd := strings.Index(rest, "</w:p>")
	if nextPEnd < 0 {
		return ""
	}
	nextPara := rest[:nextPEnd+len("</w:p>")]

	if !strings.Contains(nextPara, "<w:drawing>") {
		return ""
	}

	blipRe := regexp.MustCompile(`r:embed="(rId\d+)"`)
	m := blipRe.FindStringSubmatch(nextPara)
	if m == nil {
		return ""
	}
	return m[1]
}

// updateExistingImageDimensions finds the drawing paragraph associated with rID
// and updates both wp:extent and a:ext cx/cy values to the correct dimensions.
func updateExistingImageDimensions(docXML, rID string, cx, cy int64) string {
	rIDAttr := fmt.Sprintf(`r:embed="%s"`, rID)
	rIDIdx := strings.Index(docXML, rIDAttr)
	if rIDIdx < 0 {
		return docXML
	}

	// Find the enclosing <w:drawing>...</w:drawing>
	drawStart := strings.LastIndex(docXML[:rIDIdx], "<w:drawing>")
	if drawStart < 0 {
		return docXML
	}
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

	return strings.Replace(docXML, oldDrawing, newDrawing, 1)
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

// pngDimensionsEMU reads a PNG and returns width/height in EMUs (English Metric Units).
// Enforces max width of 4.5" and max height of 6" for 6x9 trade paperback sizing.
func pngDimensionsEMU(pngPath string) (cx, cy int64) {
	const emuPerInch = 914400
	const maxWidthInches = 4.5
	const maxHeightInches = 6.0

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

