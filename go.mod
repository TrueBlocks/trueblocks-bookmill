module github.com/TrueBlocks/trueblocks-bookmill

go 1.25.1

require (
	github.com/TrueBlocks/trueblocks-art/packages/ai v0.0.0
	github.com/TrueBlocks/trueblocks-art/packages/appd v0.0.0
	github.com/TrueBlocks/trueblocks-art/packages/bookgen v0.0.0-00010101000000-000000000000
	github.com/TrueBlocks/trueblocks-art/packages/creds v0.0.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/clipperhouse/uax29/v2 v2.7.0 // indirect
	github.com/hhrutter/lzw v1.0.0 // indirect
	github.com/hhrutter/pkcs7 v0.2.2 // indirect
	github.com/hhrutter/tiff v1.0.3 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/mattn/go-runewidth v0.0.23 // indirect
	github.com/pdfcpu/pdfcpu v0.12.1
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	golang.org/x/crypto v0.51.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/TrueBlocks/trueblocks-art/packages/aiflags v0.0.0
	github.com/TrueBlocks/trueblocks-art/packages/cli v0.0.0
	golang.org/x/image v0.43.0 // indirect
	golang.org/x/text v0.38.0 // indirect
)

replace github.com/TrueBlocks/trueblocks-art/packages/appd => ../packages/appd

replace github.com/TrueBlocks/trueblocks-art/packages/ai => ../packages/ai

replace github.com/TrueBlocks/trueblocks-art/packages/bookgen => ../packages/bookgen

replace github.com/TrueBlocks/trueblocks-art/packages/docxzip => ../packages/docxzip

replace github.com/TrueBlocks/trueblocks-art/packages/cli => ../packages/cli

replace github.com/TrueBlocks/trueblocks-art/packages/creds => ../packages/creds

replace github.com/TrueBlocks/trueblocks-art/packages/aiflags => ../packages/aiflags
