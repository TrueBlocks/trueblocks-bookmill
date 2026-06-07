#!/usr/bin/env fish

# Full pipeline: PDF → update-manifests → extract-images → extract-text → compose --split → colorize → export
#
# Usage:
#   fish pipeline.fish                  # full run with colorize
#   fish pipeline.fish --no-colorize    # B&W images only (skip OpenAI)
#   fish pipeline.fish --skip-extract   # skip steps 1-2, reuse cached images/text

set -l BOOKMILL /Users/jrush/Development/trueblocks-art/bookmill
set -l PROJECT $BOOKMILL/projects/classics/sylvan-city
set -l MANIFEST $PROJECT/manifest.yaml
set -l CACHE $HOME/.local/share/trueblocks/bookmill/classics/projects/sylvan-city
set -l PDF "$BOOKMILL/../gutenberg/Historical Books/1883 - Fords - A Sylvan City or Quaint Corners in Philadelphia.pdf"
set -l TEMPLATE $HOME/.local/share/trueblocks/works/works/templates/book-template.dotm

set -l skip_colorize false
set -l skip_extract false
for arg in $argv
    if test "$arg" = "--no-colorize"
        set skip_colorize true
    else if test "$arg" = "--skip-extract"
        set skip_extract true
    end
end

# ── Step 0: Update manifest from PDF annotations ──────────────────
echo "=== Step 0: Update manifest from PDF annotations ==="
extract-images --pdf "$PDF" --manifest "$MANIFEST" --update-manifests
or begin; echo "ERROR: update-manifests failed"; exit 1; end

if test "$skip_extract" = true
    echo "  Skipping extract steps (--skip-extract)"
else
    # ── Step 1: Extract images from annotated PDF ──────────────────
    echo ""
    echo "=== Step 1: Extract images from PDF ==="
    mkdir -p "$CACHE/extract-images"
    extract-images --pdf "$PDF" --output-dir "$CACHE/extract-images"
    or begin; echo "ERROR: extract-images failed"; exit 1; end
    set -l img_count (ls "$CACHE/extract-images"/p*.png 2>/dev/null | wc -l | string trim)
    echo "  Extracted $img_count images"

    # ── Step 2: Extract text from PDF ──────────────────────────────
    echo ""
    echo "=== Step 2: Extract text from PDF ==="
    mkdir -p "$CACHE/extract-text"
    extract-text --pdf "$PDF" --output "$CACHE/extract-text/sylvan-city.md"
    or begin; echo "ERROR: extract-text failed"; exit 1; end
    echo "  Text extracted to extract-text/sylvan-city.md"
end

# ── Step 3: Compose with chapter splitting ─────────────────────────
echo ""
echo "=== Step 3: Compose chapters (--split) ==="
rm -rf "$CACHE/compose/chapters"
compose \
    --text "$CACHE/extract-text/sylvan-city.md" \
    --manifest "$MANIFEST" \
    --output "$CACHE/compose/chapters" \
    --source-images "$CACHE/extract-images" \
    --split
or begin; echo "ERROR: compose failed"; exit 1; end

# ── Step 4: Colorize or copy B&W images ────────────────────────────
echo ""
if test "$skip_colorize" = true
    echo "=== Step 4: Skipping colorize (--no-colorize) ==="
else
    echo "=== Step 4: Colorize all images ==="
    colorize \
        --input-dir "$CACHE/extract-images" \
        --tool openai \
        --workers 4
    or begin; echo "ERROR: colorize failed"; exit 1; end
end

# ── Step 5: Export each chapter to .docx ───────────────────────────
echo ""
echo "=== Step 5: Export chapters to .docx ==="
for f in $CACHE/compose/chapters/ch*.md
    set -l base (basename $f .md)
    command export \
        --input "$f" \
        --manifest "$MANIFEST" \
        --template "$TEMPLATE" \
        --book-file "$base" \
        --skip-imageswap
    or begin; echo "ERROR: export failed for $base"; exit 1; end
end
echo "  Done"
