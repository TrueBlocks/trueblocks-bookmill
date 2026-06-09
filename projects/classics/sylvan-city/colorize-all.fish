#!/usr/bin/env fish

# Colorize images in batches, then imageswap them into the chapter .docx files.
# Output directory is read from manifest supporting_dir.
#
# Usage:
#   fish colorize-all.fish --start 1 --count 10    # colorize images 1-10, then imageswap
#   fish colorize-all.fish --start 11 --count 10   # colorize images 11-20, then imageswap
#   fish colorize-all.fish --start 21              # colorize images 21 through end, then imageswap
#   fish colorize-all.fish --no-colorize            # skip colorize, just imageswap existing images
#   fish colorize-all.fish --list                   # list all images with index numbers, then exit

set -l PROJECT /Users/jrush/Development/trueblocks-art/bookmill/projects/classics/sylvan-city
set -l MANIFEST $PROJECT/manifest.yaml
set -l CACHE $HOME/.local/share/trueblocks/bookmill/classics/projects/sylvan-city
set -l SUPPORT (grep 'supporting_dir:' "$MANIFEST" | sed 's/supporting_dir: *//' | sed "s|^~|$HOME|")
set -l IMPORTS (grep 'output_dir:' "$MANIFEST" | sed 's/output_dir: *//' | sed "s|^~|$HOME|")

# Read prompt from manifest (YAML folded scalar: lines between 'prompt: >' and next top-level key)
set -l PROMPT (python3 -c "
import yaml, sys
with open(sys.argv[1]) as f:
    m = yaml.safe_load(f)
print(m.get('prompt', ''))" "$MANIFEST")

set -l skip_colorize false
set -l list_only false
set -l start_idx 1
set -l count_val 0

set -l i 1
while test $i -le (count $argv)
    switch $argv[$i]
        case --no-colorize
            set skip_colorize true
        case --list
            set list_only true
        case --start
            set i (math $i + 1)
            set start_idx $argv[$i]
        case --count
            set i (math $i + 1)
            set count_val $argv[$i]
    end
    set i (math $i + 1)
end

# Build sorted image list from extract-images cache
set -l all_images (find "$CACHE/extract-images" -name 'p*.png' -maxdepth 1 | sort)
set -l total (count $all_images)

# ── --list: show all images with index numbers ─────────────────────
if test "$list_only" = true
    echo "Images in extract-images cache ($total total):"
    for idx in (seq 1 $total)
        set -l name (basename $all_images[$idx])
        echo "  $idx: $name"
    end
    exit 0
end

# Calculate range
if test $count_val -eq 0
    set count_val (math $total - $start_idx + 1)
end
set -l end_idx (math $start_idx + $count_val - 1)
if test $end_idx -gt $total
    set end_idx $total
end

echo "Images $start_idx-$end_idx of $total"
echo "  Output: $SUPPORT"

# ── Step 1: Colorize images ───────────────────────────────────────
if test "$skip_colorize" = true
    echo "=== Step 1: Skipping colorize (--no-colorize) ==="
else
    echo "=== Step 1: Colorize images $start_idx-$end_idx ==="
    set -l batch_dir "$CACHE/colorize-batch"
    rm -rf "$batch_dir"
    mkdir -p "$batch_dir"

    # Copy source images and build batch manifest with no_sky from main manifest
    printf 'book: batch\nimages:\n' > "$batch_dir/manifest.yaml"
    for idx in (seq $start_idx $end_idx)
        set -l name (basename $all_images[$idx])
        /bin/cp $all_images[$idx] "$batch_dir/"
        set -l no_sky (grep -A5 "file: $name" "$MANIFEST" | grep 'no_sky: true')
        if test -n "$no_sky"
            printf '  - file: %s\n    no_sky: true\n' $name >> "$batch_dir/manifest.yaml"
        else
            printf '  - file: %s\n' $name >> "$batch_dir/manifest.yaml"
        end
    end

    set -l colorize_args --input-dir "$batch_dir" --output-dir "$SUPPORT" --tool openai --workers 4
    if test -n "$PROMPT"
        set colorize_args $colorize_args --prompt "$PROMPT"
    end
    colorize $colorize_args
    or begin; echo "ERROR: colorize failed"; exit 1; end
end

# ── Step 2: Imageswap into each chapter .docx ──────────────────────
echo ""
echo "=== Step 2: Imageswap into chapter .docx files ==="

set -l swap_count 0

for docx in $IMPORTS/cEssay*.docx
    imageswap --images "$SUPPORT" "$docx"
    set swap_count (math $swap_count + 1)
    echo "  OK: "(basename "$docx" .docx)
end

# ── Done ───────────────────────────────────────────────────────────
echo ""
echo "=== Done ==="
echo "  Colorized: images $start_idx-$end_idx"
echo "  Source images: $SUPPORT"
echo "  Updated $swap_count chapter .docx files in: $IMPORTS"
