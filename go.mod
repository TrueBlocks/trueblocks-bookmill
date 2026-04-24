module github.com/TrueBlocks/trueblocks-math

go 1.25.1

require (
	github.com/TrueBlocks/trueblocks-art/packages/ai v0.0.0
	github.com/TrueBlocks/trueblocks-art/packages/bookgen v0.0.0-00010101000000-000000000000
	github.com/TrueBlocks/trueblocks-art/packages/docxzip v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.1
)

require (
	golang.org/x/image v0.39.0 // indirect
	golang.org/x/text v0.36.0 // indirect
)

replace github.com/TrueBlocks/trueblocks-art/packages/ai => ../packages/ai

replace github.com/TrueBlocks/trueblocks-art/packages/bookgen => ../packages/bookgen

replace github.com/TrueBlocks/trueblocks-art/packages/docxzip => ../packages/docxzip
