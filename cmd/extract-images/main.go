package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/TrueBlocks/trueblocks-art/packages/cli"
	"gopkg.in/yaml.v3"
)

var version = "dev"

type ImageEntry struct {
	File    string `yaml:"file"`
	Page    int    `yaml:"page"`
	Seq     int    `yaml:"seq"`
	Width   int    `yaml:"width"`
	Height  int    `yaml:"height"`
	HadBlue bool   `yaml:"had_blue,omitempty"`
	Rotated bool   `yaml:"rotated,omitempty"`
}

type Manifest struct {
	Book   string       `yaml:"book"`
	Images []ImageEntry `yaml:"images"`
}

func main() {
	app := cli.App{
		Name:        "extract-images",
		Description: "Extract images from annotated PDF pages. Detects red rectangles (image regions) and blue rectangles (erase regions), crops and cleans images.",
		Version:     version,
		Flags: []cli.FlagDef{
			{Name: "pdf", Help: "path to the annotated PDF file", Default: ""},
			{Name: "output-dir", Help: "directory for extracted images (default: ./extract-images/)", Default: ""},
			{Name: "dpi", Help: "DPI for rendering PDF pages (default: 300)", Default: 300},
			{Name: "start-page", Help: "start from this page (default: 1)", Default: 1},
			{Name: "end-page", Help: "stop at this page (default: all)", Default: 0},
			{Name: "red-threshold", Help: "minimum red channel value for red detection (default: 180)", Default: 180},
			{Name: "blue-threshold", Help: "minimum blue channel value for blue detection (default: 180)", Default: 180},
		},
		Run: run,
	}
	cli.Exit(app.Main())
}

func run(c *cli.Context) error {
	pdfPath := c.String("pdf")
	if pdfPath == "" {
		return fmt.Errorf("--pdf is required")
	}

	absPDF, err := filepath.Abs(pdfPath)
	if err != nil {
		return fmt.Errorf("resolving pdf path: %w", err)
	}
	if _, err := os.Stat(absPDF); err != nil {
		return fmt.Errorf("pdf not found: %s", absPDF)
	}

	outputDir := c.String("output-dir")
	if outputDir == "" {
		outputDir = "./extract-images"
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("creating output dir: %w", err)
	}

	dpi := c.Int("dpi")
	startPage := c.Int("start-page")
	endPage := c.Int("end-page")
	redThresh := c.Int("red-threshold")
	blueThresh := c.Int("blue-threshold")

	totalPages := countPages(absPDF)
	if totalPages == 0 {
		return fmt.Errorf("could not determine page count")
	}
	if endPage == 0 || endPage > totalPages {
		endPage = totalPages
	}

	slug := slugFromFilename(filepath.Base(absPDF))

	manifest := Manifest{Book: filepath.Base(absPDF)}
	globalSeq := 0

	for page := startPage; page <= endPage; page++ {
		fmt.Fprintf(os.Stderr, "[%d/%d] Scanning page %d...\n", page-startPage+1, endPage-startPage+1, page)

		img, err := renderPageToImage(absPDF, page, dpi)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  warning: could not render page %d: %v\n", page, err)
			continue
		}

		redBoxes := findRedBoxes(img, redThresh)
		if len(redBoxes) == 0 {
			continue
		}

		fmt.Fprintf(os.Stderr, "  found %d red box(es)\n", len(redBoxes))

		for seq, box := range redBoxes {
			globalSeq++

			rotated := false
			expandedBox := image.Rect(box.Min.X-40, box.Min.Y-40, box.Max.X+40, box.Max.Y+40)
			pageBounds := img.Bounds()
			if expandedBox.Min.X < pageBounds.Min.X {
				expandedBox.Min.X = pageBounds.Min.X
			}
			if expandedBox.Min.Y < pageBounds.Min.Y {
				expandedBox.Min.Y = pageBounds.Min.Y
			}
			if expandedBox.Max.X > pageBounds.Max.X {
				expandedBox.Max.X = pageBounds.Max.X
			}
			if expandedBox.Max.Y > pageBounds.Max.Y {
				expandedBox.Max.Y = pageBounds.Max.Y
			}
			if hasGreenCircle(img, expandedBox) {
				fmt.Fprintf(os.Stderr, "  green circle detected, will rotate 90° CCW\n")
				eraseGreenPixels(img, expandedBox)
				rotated = true
			}

			cropped := cropImage(img, box)

			hadBlue := false
			blueBoxes := findBlueBoxes(cropped, blueThresh)
			if len(blueBoxes) > 0 {
				hadBlue = true
				for _, bb := range blueBoxes {
					eraseRegion(cropped, bb)
				}
			}

			if rotated {
				cropped = rotateCCW(cropped)
			}

			filename := fmt.Sprintf("p%03d-%d.png", page, seq+1)
			outPath := filepath.Join(outputDir, filename)
			if err := savePNG(cropped, outPath); err != nil {
				fmt.Fprintf(os.Stderr, "  warning: could not save %s: %v\n", filename, err)
				continue
			}

			bounds := cropped.Bounds()
			manifest.Images = append(manifest.Images, ImageEntry{
				File:    filename,
				Page:    page,
				Seq:     globalSeq,
				Width:   bounds.Dx(),
				Height:  bounds.Dy(),
				HadBlue: hadBlue,
				Rotated: rotated,
			})

			_ = slug
		}
	}

	manifestPath := filepath.Join(outputDir, "manifest.yaml")
	data, err := yaml.Marshal(manifest)
	if err != nil {
		return fmt.Errorf("marshaling manifest: %w", err)
	}
	if err := os.WriteFile(manifestPath, data, 0644); err != nil {
		return fmt.Errorf("writing manifest: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Extracted %d images, manifest at %s\n", len(manifest.Images), manifestPath)
	return nil
}

func countPages(pdfPath string) int {
	cmd := exec.Command("pdfinfo", pdfPath)
	out, err := cmd.Output()
	if err != nil {
		return 0
	}
	for _, line := range strings.Split(string(out), "\n") {
		if strings.HasPrefix(line, "Pages:") {
			var pages int
			fmt.Sscanf(strings.TrimPrefix(line, "Pages:"), "%d", &pages)
			return pages
		}
	}
	return 0
}

func slugFromFilename(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))
	parts := strings.SplitN(name, " - ", 3)
	if len(parts) == 3 {
		return slugify(parts[2])
	}
	return slugify(name)
}

func slugify(s string) string {
	s = strings.ToLower(s)
	var result []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			result = append(result, c)
		} else if c == ' ' || c == '-' || c == '_' {
			if len(result) > 0 && result[len(result)-1] != '-' {
				result = append(result, '-')
			}
		}
	}
	return strings.Trim(string(result), "-")
}

func renderPageToImage(pdfPath string, pageNum, dpi int) (*image.NRGBA, error) {
	tmpPrefix := filepath.Join(os.TempDir(), fmt.Sprintf("extract-img-p%d", pageNum))
	tmpFile := tmpPrefix + ".png"
	defer os.Remove(tmpFile)

	cmd := exec.Command("pdftoppm",
		"-f", fmt.Sprintf("%d", pageNum),
		"-l", fmt.Sprintf("%d", pageNum),
		"-r", fmt.Sprintf("%d", dpi),
		"-png",
		"-singlefile",
		pdfPath,
		tmpPrefix,
	)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("pdftoppm failed: %w", err)
	}

	f, err := os.Open(tmpFile)
	if err != nil {
		return nil, fmt.Errorf("opening rendered page: %w", err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("decoding png: %w", err)
	}

	bounds := img.Bounds()
	nrgba := image.NewNRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			nrgba.Set(x, y, img.At(x, y))
		}
	}
	return nrgba, nil
}

func isRedPixel(c color.Color, threshold int) bool {
	r, g, b, _ := c.RGBA()
	r8, g8, b8 := r>>8, g>>8, b>>8
	return int(r8) > threshold && int(g8) < 100 && int(b8) < 100
}

func isBluePixel(c color.Color, threshold int) bool {
	r, g, b, _ := c.RGBA()
	r8, g8, b8 := r>>8, g>>8, b>>8
	return int(b8) > threshold && int(r8) < 100 && int(g8) < 100
}

func findRedBoxes(img *image.NRGBA, threshold int) []image.Rectangle {
	return findColoredBoxes(img, threshold, isRedPixel)
}

func findBlueBoxes(img *image.NRGBA, threshold int) []image.Rectangle {
	return findColoredBoxes(img, threshold, isBluePixel)
}

func findColoredBoxes(img *image.NRGBA, threshold int, isMatch func(color.Color, int) bool) []image.Rectangle {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}

	var boxes []image.Rectangle

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			ly, lx := y-bounds.Min.Y, x-bounds.Min.X
			if visited[ly][lx] {
				continue
			}
			if !isMatch(img.At(x, y), threshold) {
				continue
			}

			minX, minY, maxX, maxY := x, y, x, y
			queue := []image.Point{{x, y}}
			visited[ly][lx] = true
			count := 0

			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]
				count++

				if p.X < minX {
					minX = p.X
				}
				if p.Y < minY {
					minY = p.Y
				}
				if p.X > maxX {
					maxX = p.X
				}
				if p.Y > maxY {
					maxY = p.Y
				}

				for _, d := range []image.Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}} {
					nx, ny := p.X+d.X, p.Y+d.Y
					nlx, nly := nx-bounds.Min.X, ny-bounds.Min.Y
					if nx < bounds.Min.X || nx >= bounds.Max.X || ny < bounds.Min.Y || ny >= bounds.Max.Y {
						continue
					}
					if visited[nly][nlx] {
						continue
					}
					if isMatch(img.At(nx, ny), threshold) {
						visited[nly][nlx] = true
						queue = append(queue, image.Point{nx, ny})
					}
				}
			}

			bw, bh := maxX-minX, maxY-minY
			if bw < 50 || bh < 50 {
				continue
			}

			area := bw * bh
			fillRatio := float64(count) / float64(area)
			if fillRatio > 0.5 {
				continue
			}

			boxes = append(boxes, image.Rect(minX, minY, maxX+1, maxY+1))
		}
	}

	sort.Slice(boxes, func(i, j int) bool {
		if boxes[i].Min.Y != boxes[j].Min.Y {
			return boxes[i].Min.Y < boxes[j].Min.Y
		}
		return boxes[i].Min.X < boxes[j].Min.X
	})

	return mergeOverlapping(boxes)
}

func mergeOverlapping(boxes []image.Rectangle) []image.Rectangle {
	if len(boxes) <= 1 {
		return boxes
	}

	merged := true
	for merged {
		merged = false
		var result []image.Rectangle
		used := make([]bool, len(boxes))

		for i := 0; i < len(boxes); i++ {
			if used[i] {
				continue
			}
			box := boxes[i]
			for j := i + 1; j < len(boxes); j++ {
				if used[j] {
					continue
				}
				if box.Overlaps(boxes[j]) || touchesOrNear(box, boxes[j], 10) {
					box = box.Union(boxes[j])
					used[j] = true
					merged = true
				}
			}
			result = append(result, box)
		}
		boxes = result
	}

	return boxes
}

func touchesOrNear(a, b image.Rectangle, dist int) bool {
	expanded := image.Rect(a.Min.X-dist, a.Min.Y-dist, a.Max.X+dist, a.Max.Y+dist)
	return expanded.Overlaps(b)
}

func cropImage(img *image.NRGBA, box image.Rectangle) *image.NRGBA {
	bounds := img.Bounds()
	crop := box
	if crop.Min.X < bounds.Min.X {
		crop.Min.X = bounds.Min.X
	}
	if crop.Min.Y < bounds.Min.Y {
		crop.Min.Y = bounds.Min.Y
	}
	if crop.Max.X > bounds.Max.X {
		crop.Max.X = bounds.Max.X
	}
	if crop.Max.Y > bounds.Max.Y {
		crop.Max.Y = bounds.Max.Y
	}

	crop = trimRedEdges(img, crop)

	w, h := crop.Dx(), crop.Dy()
	result := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			result.Set(x, y, img.At(crop.Min.X+x, crop.Min.Y+y))
		}
	}
	return result
}

func trimRedEdges(img *image.NRGBA, box image.Rectangle) image.Rectangle {
	maxTrim := 30
	minRedPerSlice := 5

	isReddish := func(c color.Color) bool {
		r, g, b, _ := c.RGBA()
		r8, g8, b8 := int(r>>8), int(g>>8), int(b>>8)
		return r8 > 120 && r8 > g8+30 && r8 > b8+30
	}

	// Trim left: find the rightmost column (from the left edge) that has
	// significant red, then crop just past it
	lastRedCol := box.Min.X
	for col := box.Min.X; col < box.Min.X+maxTrim && col < box.Max.X; col++ {
		redCount := 0
		for row := box.Min.Y; row < box.Max.Y; row++ {
			if isReddish(img.At(col, row)) {
				redCount++
			}
		}
		if redCount >= minRedPerSlice {
			lastRedCol = col
		}
	}
	box.Min.X = lastRedCol + 1

	// Trim right
	lastRedCol = box.Max.X - 1
	for col := box.Max.X - 1; col > box.Max.X-1-maxTrim && col > box.Min.X; col-- {
		redCount := 0
		for row := box.Min.Y; row < box.Max.Y; row++ {
			if isReddish(img.At(col, row)) {
				redCount++
			}
		}
		if redCount >= minRedPerSlice {
			lastRedCol = col
		}
	}
	box.Max.X = lastRedCol

	// Trim top
	lastRedRow := box.Min.Y
	for row := box.Min.Y; row < box.Min.Y+maxTrim && row < box.Max.Y; row++ {
		redCount := 0
		for col := box.Min.X; col < box.Max.X; col++ {
			if isReddish(img.At(col, row)) {
				redCount++
			}
		}
		if redCount >= minRedPerSlice {
			lastRedRow = row
		}
	}
	box.Min.Y = lastRedRow + 1

	// Trim bottom
	lastRedRow = box.Max.Y - 1
	for row := box.Max.Y - 1; row > box.Max.Y-1-maxTrim && row > box.Min.Y; row-- {
		redCount := 0
		for col := box.Min.X; col < box.Max.X; col++ {
			if isReddish(img.At(col, row)) {
				redCount++
			}
		}
		if redCount >= minRedPerSlice {
			lastRedRow = row
		}
	}
	box.Max.Y = lastRedRow

	return box
}

func eraseRegion(img *image.NRGBA, box image.Rectangle) {
	white := color.NRGBA{255, 255, 255, 255}
	bounds := img.Bounds()
	for y := box.Min.Y; y < box.Max.Y && y < bounds.Max.Y; y++ {
		for x := box.Min.X; x < box.Max.X && x < bounds.Max.X; x++ {
			if x >= bounds.Min.X && y >= bounds.Min.Y {
				img.Set(x, y, white)
			}
		}
	}
}

func savePNG(img *image.NRGBA, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}

func isGreenPixel(c color.Color) bool {
	r, g, b, _ := c.RGBA()
	r8, g8, b8 := int(r>>8), int(g>>8), int(b>>8)
	return g8 > 120 && g8 > r8+20 && g8 > b8+20
}

func hasGreenCircle(img *image.NRGBA, region image.Rectangle) bool {
	count := 0
	for y := region.Min.Y; y < region.Max.Y; y++ {
		for x := region.Min.X; x < region.Max.X; x++ {
			if isGreenPixel(img.At(x, y)) {
				count++
			}
		}
	}
	return count > 50
}

func eraseGreenPixels(img *image.NRGBA, region image.Rectangle) {
	white := color.NRGBA{255, 255, 255, 255}
	for y := region.Min.Y; y < region.Max.Y; y++ {
		for x := region.Min.X; x < region.Max.X; x++ {
			if isGreenPixel(img.At(x, y)) {
				img.Set(x, y, white)
			}
		}
	}
}

func rotateCCW(img *image.NRGBA) *image.NRGBA {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	rotated := image.NewNRGBA(image.Rect(0, 0, h, w))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			rotated.Set(y, w-1-x, img.At(bounds.Min.X+x, bounds.Min.Y+y))
		}
	}
	return rotated
}
