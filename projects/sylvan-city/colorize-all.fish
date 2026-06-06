#!/usr/bin/env fish

set -l PROJECT /Users/jrush/Development/trueblocks-art/bookmill/projects/sylvan-city
set -l SUPPORT ~/Documents/Classics/"100 Collections"/Supporting/"A Sylvan City"
set -l SLUG "cEssay - 1883 - Fords - A Sylvan City or Quaint Corners in Philadelphia"

echo "=== Step 1: Colorize all images ==="
colorize \
    --input-dir "$PROJECT/extract-images" \
    --output-dir "$PROJECT/colorize" \
    --copy-to "$SUPPORT" \
    --tool openai
or begin
    echo "ERROR: colorize failed"
    exit 1
end

echo ""
echo "=== Step 2: Copy colorized images into export/images for imageswap ==="
set -l IMGDIR "$PROJECT/export/images/$SLUG"
mkdir -p "$IMGDIR"
/bin/cp "$PROJECT/colorize"/p*.png "$IMGDIR/"
or begin
    echo "ERROR: failed to copy images to export/images"
    exit 1
end

echo ""
echo "=== Step 3: Run imageswap to insert images into .docx ==="
set -l DOCX "$PROJECT/export/$SLUG.docx"
imageswap --images "$PROJECT/export/images" "$DOCX"
or begin
    echo "ERROR: imageswap failed"
    exit 1
end

echo ""
echo "=== Done ==="
echo "Colorized images in: $PROJECT/colorize/"
echo "Copied to: $SUPPORT"
echo "Updated docx: $DOCX"
