package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
	"testing"
)

func TestReTagLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantFile string
		wantFull string
		noMatch  bool
	}{
		{
			name:     "IMG tag with caption",
			input:    "[[IMG:photo.png|A nice caption]]",
			wantFile: "photo.png",
			wantFull: "[[IMG:photo.png|A nice caption]]",
		},
		{
			name:     "IMG tag without caption",
			input:    "[[IMG:photo.png]]",
			wantFile: "photo.png",
			wantFull: "[[IMG:photo.png]]",
		},
		{
			name:     "R tag",
			input:    "[[R:diagram.png]]",
			wantFile: "diagram.png",
			wantFull: "[[R:diagram.png]]",
		},
		{
			name:     "IMG tag with pipe in caption",
			input:    "[[IMG:fig1.png|cap|extra]]",
			wantFile: "fig1.png",
			wantFull: "[[IMG:fig1.png|cap|extra]]",
		},
		{
			name:    "no match plain text",
			input:   "just some text",
			noMatch: true,
		},
		{
			name:    "no match partial tag",
			input:   "[[OTHER:file.png]]",
			noMatch: true,
		},
		{
			name:     "tag embedded in text",
			input:    "before [[R:test.png]] after",
			wantFile: "test.png",
			wantFull: "[[R:test.png]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := reTagLine.FindStringSubmatch(tt.input)
			if tt.noMatch {
				if m != nil {
					t.Errorf("expected no match, got %v", m)
				}
				return
			}
			if m == nil {
				t.Fatal("expected match, got nil")
			}
			if m[0] != tt.wantFull {
				t.Errorf("full match = %q, want %q", m[0], tt.wantFull)
			}
			if m[1] != tt.wantFile {
				t.Errorf("capture group = %q, want %q", m[1], tt.wantFile)
			}
		})
	}
}

func TestFindMaxRID(t *testing.T) {
	tests := []struct {
		name    string
		relsXML string
		want    int
	}{
		{
			name:    "empty rels",
			relsXML: `<Relationships></Relationships>`,
			want:    0,
		},
		{
			name:    "single rId",
			relsXML: `<Relationship Id="rId1" Type="foo" Target="bar"/>`,
			want:    1,
		},
		{
			name: "multiple rIds",
			relsXML: `<Relationships>
				<Relationship Id="rId1" Type="a" Target="b"/>
				<Relationship Id="rId5" Type="c" Target="d"/>
				<Relationship Id="rId3" Type="e" Target="f"/>
			</Relationships>`,
			want: 5,
		},
		{
			name:    "no rId attributes",
			relsXML: `<Relationships><Something/></Relationships>`,
			want:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxRID(tt.relsXML)
			if got != tt.want {
				t.Errorf("findMaxRID() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestFindMaxImageNum(t *testing.T) {
	tests := []struct {
		name     string
		existing map[string]bool
		want     int
	}{
		{
			name:     "empty map",
			existing: map[string]bool{},
			want:     0,
		},
		{
			name: "single image",
			existing: map[string]bool{
				"word/media/image1.png": true,
			},
			want: 1,
		},
		{
			name: "multiple images",
			existing: map[string]bool{
				"word/media/image1.png": true,
				"word/media/image7.png": true,
				"word/media/image3.png": true,
			},
			want: 7,
		},
		{
			name: "non-matching paths ignored",
			existing: map[string]bool{
				"word/media/image2.png":  true,
				"word/media/image5.jpeg": true, // not png
				"word/media/photo1.png":  true, // not imageN
				"other/image99.png":      true, // wrong prefix
			},
			want: 2,
		},
		{
			name: "only non-matching paths",
			existing: map[string]bool{
				"word/media/chart1.emf": true,
				"word/media/photo.png":  true,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxImageNum(tt.existing)
			if got != tt.want {
				t.Errorf("findMaxImageNum() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestNormalizeTagParagraph(t *testing.T) {
	tests := []struct {
		name     string
		docXML   string
		fullTag  string
		filename string
		want     string
	}{
		{
			name:     "tag found and normalized",
			docXML:   `<w:p><w:r><w:t>[[IMG:photo.png|caption]]</w:t></w:r></w:p>`,
			fullTag:  "[[IMG:photo.png|caption]]",
			filename: "photo.png",
			want:     `<w:p><w:pPr><w:pStyle w:val="ImageTag"/></w:pPr><w:r><w:rPr><w:vanish/></w:rPr><w:t>[[R:photo.png]]</w:t></w:r></w:p>`,
		},
		{
			name:     "tag not found returns unchanged",
			docXML:   `<w:p><w:r><w:t>no tag here</w:t></w:r></w:p>`,
			fullTag:  "[[IMG:missing.png]]",
			filename: "missing.png",
			want:     `<w:p><w:r><w:t>no tag here</w:t></w:r></w:p>`,
		},
		{
			name:     "R tag normalized",
			docXML:   `<w:p><w:r><w:t>[[R:fig.png]]</w:t></w:r></w:p>`,
			fullTag:  "[[R:fig.png]]",
			filename: "fig.png",
			want:     `<w:p><w:pPr><w:pStyle w:val="ImageTag"/></w:pPr><w:r><w:rPr><w:vanish/></w:rPr><w:t>[[R:fig.png]]</w:t></w:r></w:p>`,
		},
		{
			name:     "paragraph with attributes",
			docXML:   `<w:p w:rsid="001"><w:r><w:t>[[IMG:test.png]]</w:t></w:r></w:p>`,
			fullTag:  "[[IMG:test.png]]",
			filename: "test.png",
			want:     `<w:p><w:pPr><w:pStyle w:val="ImageTag"/></w:pPr><w:r><w:rPr><w:vanish/></w:rPr><w:t>[[R:test.png]]</w:t></w:r></w:p>`,
		},
		{
			name:     "preserves surrounding paragraphs",
			docXML:   `<w:p><w:r><w:t>before</w:t></w:r></w:p><w:p><w:r><w:t>[[R:x.png]]</w:t></w:r></w:p><w:p><w:r><w:t>after</w:t></w:r></w:p>`,
			fullTag:  "[[R:x.png]]",
			filename: "x.png",
			want:     `<w:p><w:r><w:t>before</w:t></w:r></w:p><w:p><w:pPr><w:pStyle w:val="ImageTag"/></w:pPr><w:r><w:rPr><w:vanish/></w:rPr><w:t>[[R:x.png]]</w:t></w:r></w:p><w:p><w:r><w:t>after</w:t></w:r></w:p>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeTagParagraph(tt.docXML, tt.fullTag, tt.filename)
			if got != tt.want {
				t.Errorf("normalizeTagParagraph()\ngot:  %s\nwant: %s", got, tt.want)
			}
		})
	}
}

func TestFindExistingImageRID(t *testing.T) {
	tests := []struct {
		name    string
		docXML  string
		file    string
		wantRID string
	}{
		{
			name:    "no tag found",
			docXML:  `<w:p><w:r><w:t>nothing</w:t></w:r></w:p>`,
			file:    "missing.png",
			wantRID: "",
		},
		{
			name: "tag found but no drawing after",
			docXML: `<w:p><w:r><w:t>[[R:photo.png]]</w:t></w:r></w:p>` +
				`<w:p><w:r><w:t>plain text</w:t></w:r></w:p>`,
			file:    "photo.png",
			wantRID: "",
		},
		{
			name: "tag found with drawing after",
			docXML: `<w:p><w:r><w:t>[[R:photo.png]]</w:t></w:r></w:p>` +
				`<w:p><w:r><w:drawing><wp:inline><a:blip r:embed="rId7"/></wp:inline></w:drawing></w:r></w:p>`,
			file:    "photo.png",
			wantRID: "rId7",
		},
		{
			name: "drawing without r:embed",
			docXML: `<w:p><w:r><w:t>[[R:photo.png]]</w:t></w:r></w:p>` +
				`<w:p><w:r><w:drawing><wp:inline><a:blip/></wp:inline></w:drawing></w:r></w:p>`,
			file:    "photo.png",
			wantRID: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findExistingImageRID(tt.docXML, tt.file)
			if got != tt.wantRID {
				t.Errorf("findExistingImageRID() = %q, want %q", got, tt.wantRID)
			}
		})
	}
}

func TestFindRelTarget(t *testing.T) {
	rels := `<Relationships>
		<Relationship Id="rId1" Type="a" Target="styles.xml"/>
		<Relationship Id="rId5" Type="image" Target="media/image3.png"/>
	</Relationships>`

	tests := []struct {
		name string
		rID  string
		want string
	}{
		{name: "found rId1", rID: "rId1", want: "styles.xml"},
		{name: "found rId5", rID: "rId5", want: "media/image3.png"},
		{name: "not found", rID: "rId99", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRelTarget(rels, tt.rID)
			if got != tt.want {
				t.Errorf("findRelTarget() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestInsertParagraphAfterTag(t *testing.T) {
	tests := []struct {
		name         string
		docXML       string
		filename     string
		imgParagraph string
		want         string
	}{
		{
			name:         "tag found, paragraph inserted",
			docXML:       `<w:p><w:r><w:t>[[R:fig.png]]</w:t></w:r></w:p><w:p><w:r><w:t>next</w:t></w:r></w:p>`,
			filename:     "fig.png",
			imgParagraph: `<w:p>IMAGE</w:p>`,
			want:         `<w:p><w:r><w:t>[[R:fig.png]]</w:t></w:r></w:p><w:p>IMAGE</w:p><w:p><w:r><w:t>next</w:t></w:r></w:p>`,
		},
		{
			name:         "tag not found, unchanged",
			docXML:       `<w:p><w:r><w:t>no tag</w:t></w:r></w:p>`,
			filename:     "missing.png",
			imgParagraph: `<w:p>IMAGE</w:p>`,
			want:         `<w:p><w:r><w:t>no tag</w:t></w:r></w:p>`,
		},
		{
			name:         "tag at end of document",
			docXML:       `<w:p><w:r><w:t>[[R:end.png]]</w:t></w:r></w:p>`,
			filename:     "end.png",
			imgParagraph: `<w:p>IMG</w:p>`,
			want:         `<w:p><w:r><w:t>[[R:end.png]]</w:t></w:r></w:p><w:p>IMG</w:p>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertParagraphAfterTag(tt.docXML, tt.filename, tt.imgParagraph)
			if got != tt.want {
				t.Errorf("insertParagraphAfterTag()\ngot:  %s\nwant: %s", got, tt.want)
			}
		})
	}
}

func TestBuildImageParagraph(t *testing.T) {
	result := buildImageParagraph("rId10", "image5.png", "photo.png", 1000000, 750000, 5)

	// Verify key elements are present
	checks := []string{
		`cx="1000000"`,
		`cy="750000"`,
		`r:embed="rId10"`,
		`name="image5.png"`,
		`descr="photo.png"`,
		`id="5"`,
		`<w:jc w:val="center"/>`,
		`<wp:inline`,
		`<pic:pic`,
	}
	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("buildImageParagraph() missing %q", check)
		}
	}

	// Verify it starts and ends correctly
	if !strings.HasPrefix(result, "<w:p>") {
		t.Error("should start with <w:p>")
	}
	if !strings.HasSuffix(result, "</w:p>") {
		t.Error("should end with </w:p>")
	}
}

func TestPngDimensionsEMU(t *testing.T) {
	const emuPerInch int64 = 914400

	// Helper to create a test PNG of given pixel dimensions
	createTestPNG := func(t *testing.T, w, h int) string {
		t.Helper()
		f, err := os.CreateTemp(t.TempDir(), "test-*.png")
		if err != nil {
			t.Fatal(err)
		}
		img := image.NewRGBA(image.Rect(0, 0, w, h))
		// Fill with a color so it's a valid PNG
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				img.Set(x, y, color.RGBA{R: 100, G: 100, B: 100, A: 255})
			}
		}
		if err := png.Encode(f, img); err != nil {
			t.Fatal(err)
		}
		f.Close()
		return f.Name()
	}

	t.Run("small image under limits", func(t *testing.T) {
		// 300x300 pixels = 1" x 1" at 300dpi — well under 4.5" x 6"
		path := createTestPNG(t, 300, 300)
		cx, cy := pngDimensionsEMU(path)
		wantCX := emuPerInch // 1 inch
		wantCY := emuPerInch // 1 inch
		if cx != wantCX || cy != wantCY {
			t.Errorf("got (%d, %d), want (%d, %d)", cx, cy, wantCX, wantCY)
		}
	})

	t.Run("wide image needs width scaling", func(t *testing.T) {
		// 3000x300 pixels = 10" x 1" at 300dpi — needs width scaling to 4.5"
		path := createTestPNG(t, 3000, 300)
		cx, cy := pngDimensionsEMU(path)
		// After scaling: 4.5" wide, 0.45" tall
		wantCX := int64(float64(emuPerInch) * 4.5) // 4114800
		wantCY := int64(float64(emuPerInch) * 0.45) // 411480
		if cx != wantCX || cy != wantCY {
			t.Errorf("got (%d, %d), want (%d, %d)", cx, cy, wantCX, wantCY)
		}
	})

	t.Run("tall image needs height scaling", func(t *testing.T) {
		// 300x3000 pixels = 1" x 10" at 300dpi — needs height scaling to 6"
		path := createTestPNG(t, 300, 3000)
		cx, cy := pngDimensionsEMU(path)
		// After scaling: 0.6" wide, 6" tall
		wantCX := int64(float64(emuPerInch) * 0.6)
		wantCY := int64(float64(emuPerInch) * 6.0)
		if cx != wantCX || cy != wantCY {
			t.Errorf("got (%d, %d), want (%d, %d)", cx, cy, wantCX, wantCY)
		}
	})

	t.Run("both dimensions exceed limits", func(t *testing.T) {
		// 3000x2700 pixels = 10" x 9" — first width-scaled to 4.5"/4.05", then height-scaled
		path := createTestPNG(t, 3000, 2700)
		cx, cy := pngDimensionsEMU(path)
		// After width scaling: 4.5" x 4.05" — both under limits, no height scaling needed
		wantCX := int64(float64(emuPerInch) * 4.5)
		wantCY := int64(float64(emuPerInch) * 4.05)
		if cx != wantCX || cy != wantCY {
			t.Errorf("got (%d, %d), want (%d, %d)", cx, cy, wantCX, wantCY)
		}
	})

	t.Run("invalid path returns default", func(t *testing.T) {
		cx, cy := pngDimensionsEMU("/nonexistent/path.png")
		if cx != 4114800 || cy != 2743200 {
			t.Errorf("got (%d, %d), want (4114800, 2743200)", cx, cy)
		}
	})
}

func TestUpdateExistingImageDimensions(t *testing.T) {
	t.Run("rID found with dimensions", func(t *testing.T) {
		docXML := `<w:p><w:drawing><wp:inline><wp:extent cx="100" cy="200"/><a:blip r:embed="rId5"/><a:ext cx="100" cy="200"/></wp:inline></w:drawing></w:p>`
		got := updateExistingImageDimensions(docXML, "rId5", 999, 888)
		if !strings.Contains(got, `cx="999"`) || !strings.Contains(got, `cy="888"`) {
			t.Errorf("dimensions not updated:\n%s", got)
		}
		// Should not contain old values
		if strings.Contains(got, `cx="100"`) || strings.Contains(got, `cy="200"`) {
			t.Errorf("old dimensions still present:\n%s", got)
		}
	})

	t.Run("rID not found returns unchanged", func(t *testing.T) {
		docXML := `<w:p><w:drawing><wp:inline><wp:extent cx="100" cy="200"/><a:blip r:embed="rId5"/></wp:inline></w:drawing></w:p>`
		got := updateExistingImageDimensions(docXML, "rId99", 999, 888)
		if got != docXML {
			t.Errorf("expected unchanged document, got:\n%s", got)
		}
	})
}

