#!/usr/bin/env fish

# Colorize images 71-167 in batches of 20, with retry logic for rate limiting.
# Moderation-blocked images are recorded but not retried.
# Final imageswap runs once at the end.

set -l PROJECT /Users/jrush/Development/trueblocks-art/bookmill/projects/classics/sylvan-city
set -l CACHE $HOME/.local/share/trueblocks/bookmill/classics/projects/sylvan-city
set -l SUPPORT $HOME/Documents/Classics/"100 Collections"/Supporting/"A Sylvan City"
set -l IMPORTS $HOME/Development/trueblocks-art/works/imports/files
set -l PDF_NAME "1883 - Fords - A Sylvan City or Quaint Corners in Philadelphia.pdf"
set -l YEAR (echo "$PDF_NAME" | sed 's/ -.*//')
set -l BOOK_PREFIX "cEssay - $YEAR - ch"

set -l BATCH_SIZE 20
set -l START 71
set -l TOTAL 167
set -l WORKERS 4
set -l PAUSE 15
set -l all_failed
set -l moderation_blocked

echo "=== Colorizing remaining images $START-$TOTAL ==="
echo "  Batch size: $BATCH_SIZE, Workers: $WORKERS, Pause: {$PAUSE}s"
echo ""

set -l all_images (find "$SUPPORT" -name 'p*.png' -maxdepth 1 | sort)
set -l actual_total (count $all_images)
if test $TOTAL -gt $actual_total
    set TOTAL $actual_total
end

set -l batch_start $START
while test $batch_start -le $TOTAL
    set -l remaining (math $TOTAL - $batch_start + 1)
    set -l count $BATCH_SIZE
    if test $remaining -lt $BATCH_SIZE
        set count $remaining
    end
    set -l batch_end (math $batch_start + $count - 1)

    echo "── Batch: images $batch_start-$batch_end ($WORKERS workers) ──"

    # Build batch manifest
    set -l batch_dir "$CACHE/colorize-batch"
    rm -rf "$batch_dir"
    mkdir -p "$batch_dir"

    printf 'book: batch\nimages:\n' > "$batch_dir/manifest.yaml"
    for idx in (seq $batch_start $batch_end)
        set -l name (basename $all_images[$idx])
        /bin/cp "$all_images[$idx]" "$batch_dir/"
        printf '  - file: %s\n' $name >> "$batch_dir/manifest.yaml"
    end

    # Run colorize
    colorize \
        --input-dir "$batch_dir" \
        --output-dir "$SUPPORT" \
        --tool openai \
        --workers $WORKERS 2>&1 | tee /tmp/colorize-batch.log

    # Check for failures
    set -l batch_failures (grep "FAILED" /tmp/colorize-batch.log | grep -oE 'p[0-9]+-[0-9]+\.png')

    if test (count $batch_failures) -gt 0
        echo "  "(count $batch_failures)" failed in this batch"

        # Separate moderation blocks from retryable failures
        set -l retryable
        for img in $batch_failures
            set -l line (grep "$img" /tmp/colorize-batch.log | grep "FAILED")
            if echo "$line" | grep -q "moderation_blocked"
                echo "  $img — moderation blocked (will not retry)"
                set -a moderation_blocked $img
            else
                set -a retryable $img
            end
        end

        # Retry non-moderation failures
        if test (count $retryable) -gt 0
            # Check if rate limited
            if grep -qiE "rate_limit|429|Too Many Requests" /tmp/colorize-batch.log
                echo "  Rate limiting detected — reducing workers and pausing 60s"
                set WORKERS 2
                set PAUSE 60
                sleep 60
            else
                echo "  Retrying "(count $retryable)" images after 30s pause..."
                sleep 30
            end

            set -l retry_dir "$CACHE/colorize-retry"
            rm -rf "$retry_dir"
            mkdir -p "$retry_dir"
            printf 'book: retry\nimages:\n' > "$retry_dir/manifest.yaml"
            for img in $retryable
                if test -f "$SUPPORT/$img"
                    /bin/cp "$SUPPORT/$img" "$retry_dir/"
                    printf '  - file: %s\n' $img >> "$retry_dir/manifest.yaml"
                end
            end

            echo "  Retrying "(count $retryable)" images with 1 worker..."
            colorize \
                --input-dir "$retry_dir" \
                --output-dir "$SUPPORT" \
                --tool openai \
                --workers 1 2>&1 | tee /tmp/colorize-retry.log

            set -l retry_failures (grep "FAILED" /tmp/colorize-retry.log | grep -oE 'p[0-9]+-[0-9]+\.png')
            for f in $retry_failures
                if not echo "$f" | grep -q "moderation_blocked"
                    set -a all_failed $f
                end
            end
            rm -rf "$retry_dir"
        end
    end

    rm -rf "$batch_dir"
    set batch_start (math $batch_start + $count)

    if test $batch_start -le $TOTAL
        echo "  Pausing {$PAUSE}s..."
        sleep $PAUSE
    end
    echo ""
end

# Final imageswap on all chapter .docx files
echo "=== Final imageswap ==="
set -l tmpimg (mktemp -d)
for docx in $IMPORTS/$BOOK_PREFIX*.docx
    set -l base (basename "$docx" .docx)
    set -l slugdir "$tmpimg/$base/images"
    mkdir -p "$slugdir"
    for img in $SUPPORT/p*.png
        /bin/cp "$img" "$slugdir/"
    end
    imageswap --images "$tmpimg" "$docx"
    echo "  OK: $base"
end
rm -rf "$tmpimg"

# Summary
echo ""
echo "=========================================="
echo "ALL DONE — images $START through $TOTAL"
echo "=========================================="
if test (count $moderation_blocked) -gt 0
    echo "  Moderation blocked ("(count $moderation_blocked)"):"
    for f in $moderation_blocked
        echo "    $f"
    end
end
if test (count $all_failed) -gt 0
    echo "  Other failures ("(count $all_failed)"):"
    for f in $all_failed
        echo "    $f"
    end
end
if test (count $moderation_blocked) -eq 0; and test (count $all_failed) -eq 0
    echo "  All images colorized successfully!"
end
