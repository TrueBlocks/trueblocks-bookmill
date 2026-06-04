# Pipeline Spec: Idea to .docx

> How the math essay pipeline transforms a two-sentence idea into a fully illustrated, publication-ready Word document.

---

## Overview

The pipeline is a Go-based automated system that takes a collection of mathematical essay **ideas** (each a slug, title, hook, and hidden-math sentence) and advances them through eight stages — research, outline, draft, fact-check, illustrate, second draft, and export — producing a formatted `.docx` file with embedded images for each essay.

The pipeline runs as a long-lived process with a web dashboard, cycling every ~15 seconds, processing multiple essays concurrently via the Anthropic (Claude) API. It handles three **item types** — essays, section dividers, and book introductions — across a three-book series (~132 items total). Section dividers and introductions use the same stage infrastructure but auto-skip stages where no real work is needed.

For details on sections, see [section-dividers.md](section-dividers.md).  
For details on introductions, see [book-introductions.md](book-introductions.md).

---

## Entry Points

| Binary | Purpose |
|--------|---------|
| `pipeline` | Main orchestrator — runs cycles, calls Claude, manages state, serves dashboard |
| `scaffold` | One-shot generator — creates idea `.md` + `.meta.yaml` files for essays, sections, and introductions |
| `imagerender` | Renders Mermaid diagrams, R charts, and AI-generated images to PNG |
| `imageswap` | Embeds PNG images into `.docx` files, replacing placeholder tags |

All binaries are built from `cmd/` and installed via the `Makefile` to `~/source/`.

---

## The Eight Stages

Each essay progresses sequentially through these directories inside `projects/math-books/`:

```
ideas/ → research/ → outline/ → draft/ → factcheck/ → illustrate/ → draft2/ → export/
```

At every stage, two files exist per item:
- `{slug}.md` — the content (markdown for stages 1–7; .docx in export)
- `{slug}.meta.yaml` — metadata tracking status, type, model, tokens, cost, timestamps

### Stage 1: Ideas

**Source:** `ideas/{slug}.md` + `ideas/{slug}.meta.yaml`

Created by the `scaffold` tool. Each idea contains:
- Essay title
- **The Everyday Experience** — a hook the reader can relate to
- **The Hidden Math** — the mathematical punchline
- **Placement** — book, part, and chapter order

The `.meta.yaml` records: slug, title, type (essay/section/introduction), book (I/II/III), part, part_title, order, status, model, created date.

The `type` field determines how the item flows through the pipeline:

| Type | Description | Skipped stages |
|------|-------------|----------------|
| `essay` | Normal essay. Full pipeline. | None |
| `section` | Part title page. Short paragraph + placeholder image. | research, outline, draft, factcheck, illustrate |
| `introduction` | Book-level intro chapter. ~400–500 words, no images. | research, factcheck, illustrate |

At skipped stages, the pipeline copies the previous stage's `.md` forward verbatim
and marks the stage complete with zero tokens and zero cost. See the
[section-dividers](section-dividers.md) and [book-introductions](book-introductions.md)
specs for full details.

### Stage 2: Research

**Prompt focus:** Produce a thorough research brief (2,000–3,000 words) covering:
- Mathematical foundation (core concepts, equations, common misconceptions)
- History and discovery (who, when, surprising details)
- Everyday connections (exact measurements, concrete examples)
- Worked examples (2–3 step-by-step numerical examples)
- Surprising extensions (connections, modern applications)
- Sources and citations (original papers, contested claims)

**Input:** The idea file (hook + hidden math).

### Stage 3: Outline

**Prompt focus:** Structure the essay using a **narrative arc** (one of 15 arcs, editorially assigned — see [narrative-arcs-and-endings.md](narrative-arcs-and-endings.md)). For each beat in the arc:
- Section title
- 2–3 sentence summary
- Key facts from research
- The "turn" (surprise moment)

The outline also specifies target reading time (drawn from a Gaussian around `read_mean` minutes), which sets the word-count target (1 minute ≈ 265 words).

**Input:** The idea + the research brief.

### Stage 4: Draft

**Prompt focus:** Write a complete first draft from the outline and research. Requirements:
- Match target word count and reading time
- Voice: warm, Asimov-meets-Mary Roach
- Every equation preceded by intuition, followed by interpretation
- Concrete examples with real numbers
- Avoid: jargon without explanation, false enthusiasm, hedging, "actually"

**Input:** The idea + the research + the outline.

### Stage 5: Fact-check

**Prompt focus:** Verify the draft against the research and general knowledge:
- Find factual errors, misleading simplifications, missing attribution, tone issues
- For each issue: quote the passage, explain the error, provide correction
- Overall assessment: PASS, REVISE, or FAIL

**Input:** The draft + the research.

### Stage 6: Illustrate

**Prompt focus:** Plan and generate image source code. Rules:
- 1–2 images maximum (1 hero image mandatory)
- Width: 4.5" max for 6×9 book format
- Three methods: Mermaid diagram, R plot, or AI prompt (DALL-E)
- No purely decorative images; every image must teach
- Output: `[[IMG:filename.png|caption]]` tags inserted in the markdown, plus image source code between delimited markers

**Input:** The draft + the fact-check results.

**Post-processing:** The pipeline parses image source blocks from the response, writes `.mermaid`, `.R`, or `.ai-prompt.txt` files into `images/{slug}/`, then calls `imagerender` to produce PNGs.

### Stage 7: Draft2 (Final Revision)

**Prompt focus:** Incorporate fact-check corrections and image references into a polished final draft:
- Fix all issues identified in fact-check
- Integrate image tags naturally into the text flow
- Maintain narrative arc coherence
- Match target word count

**Input:** The draft + the fact-check + the illustrate output.

### Stage 8: Export

**Not an AI stage.** This is a mechanical conversion:

1. Read the `draft2` markdown (or `illustrate` if draft2 is unavailable)
2. Sanitize markdown for Word compatibility (smart quotes, special characters)
3. Run `md2docx` — a Pandoc-based template conversion to `.docx`
4. Run `imagerender` — ensure all images are rendered as PNGs
5. Run `imageswap` — embed PNGs into the `.docx`, replacing `[[IMG:...]]` and `[[R:...]]` tags
6. Run `upgradeDocx` — AppleScript to set page layout, scale images, copy styles from a Word template

**Output:** `export/cChapter - 2026 - {Book}.{Part}.{Order} {Title}.docx`

---

## Narrative Arcs and Endings

Each essay (type `essay` only — not sections or introductions) is assigned one
of **15 narrative arcs** and one of **6 ending types**. Arcs shape the essay's
structure; endings shape how it lands emotionally. Assignment is editorial,
based on the essay's topic and position in the book, not random.

Arc and ending are recorded in each essay's `.meta.yaml`:

```yaml
arc: slow-build
ending: resolution
```

See [narrative-arcs-and-endings.md](narrative-arcs-and-endings.md) for the full
arc catalog, ending types, interaction matrix, and position-sensitive defaults.

---

## Configuration

Stored at `~/.local/share/trueblocks/math/config.yaml`:

```yaml
api:
  anthropic_key: "sk-..."
  openai_key: "sk-..."

models:
  research: claude-sonnet-4-20250514
  outline: claude-sonnet-4-20250514
  draft: claude-sonnet-4-20250514
  factcheck: claude-sonnet-4-20250514
  draft2: claude-sonnet-4-20250514
  illustrate: claude-sonnet-4-20250514

pipeline:
  max_per_cycle: 6        # Total essays to process per cycle
  new_per_cycle: 4        # Max new essays (not continuations) per cycle
  dry_run: false           # Use placeholder responses instead of API
  cycle_interval: 15       # Seconds between cycles
  api_timeout: 300         # Timeout per API call (seconds)
  concurrency: 3           # Concurrent essay jobs
  read_mean: 5.0           # Target minutes per essay (× 265 = word count)
  read_spread: 1.0         # Gaussian variance for target length

dashboard:
  port: 8787
```

Config is re-read from disk on every cycle, allowing live tuning.

---

## Image Pipeline

### Three Rendering Methods

**Mermaid** — Diagrams and flowcharts
- Source: `.mermaid` files in `images/{slug}/`
- Rendered via `mmdc` CLI with custom theme (`data/mermaid-theme.json`)
- On failure: up to 3 AI-assisted repair attempts via Claude

**R** — Data visualizations
- Source: `.R` files in `images/{slug}/`
- Rendered via `Rscript`, sourcing `data/common.R` for consistent theme
- Color palette: 7 named colors (math_blue, math_purple, math_emerald, etc.)
- On failure: up to 3 AI-assisted repair attempts

**AI** — Complex visual concepts
- Source: `.ai-prompt.txt` files in `images/{slug}/`
- Generated via OpenAI DALL-E-3 API
- No repair cycle (prompt-based, not code-based)

### Image Embedding (imageswap)

After `.docx` creation, `imageswap`:
1. Opens the `.docx` as a ZIP archive
2. Parses `word/document.xml` for `[[IMG:filename|caption]]` and `[[R:filename]]` tags
3. For each tag:
   - Reads the PNG and calculates dimensions in EMU (English Metric Units)
   - Adds the PNG to `word/media/` inside the ZIP
   - Inserts proper Open XML `<w:drawing>` elements
   - Updates `word/_rels/document.xml.rels` and `[Content_Types].xml`
4. Normalizes all tags to `[[R:filename]]` with ImageTag style
5. Re-zips everything

### Word Template Upgrade (upgradeDocx)

AppleScript automation:
- Opens the exported `.docx` in Microsoft Word
- Copies styles from a `.dotm` template
- Sets page layout (6×9 book format, mirror margins, gutter)
- Scales all images to fit the page width
- Sets the document title metadata

---

## Dashboard

HTTP server at `http://127.0.0.1:8787/` providing:

| Endpoint | Method | Purpose |
|----------|--------|---------|
| `/` | GET | HTML dashboard with real-time status |
| `/api/status` | GET | JSON: cycle count, costs, timing, project summaries |
| `/api/essays` | GET | List all essays with current status |
| `/api/step` | POST | Manually trigger one processing cycle |
| `/api/open` | GET | Open essay folder in Finder |
| `/api/open-docx` | GET | Open exported .docx in Word |
| `/api/logs` | GET | Fetch log buffer (last 1000 entries) |
| `/api/settings` | GET | Current configuration |

The dashboard shows per-project progress bars, cycle timer, cost tracking, and a log viewer.

---

## Cycle Selection Logic

Each cycle, the pipeline selects essays to process:

1. **Continuations first** — essays already in progress (past the ideas stage) are prioritized
2. **New essays** — essays still at the ideas stage, up to `new_per_cycle` limit
3. **Total cap** — no more than `max_per_cycle` essays per cycle total
4. **Concurrency** — a semaphore limits concurrent API calls to `concurrency` goroutines

This means a cycle might advance 2 in-progress essays to their next stage *and* start 4 new essays at the research stage, processing up to 3 simultaneously.

---

## Cost Model

API costs are tracked per essay and per session:

| Model | Input | Output |
|-------|-------|--------|
| Opus models | $15 / 1M tokens | $75 / 1M tokens |
| All others | $3 / 1M tokens | $15 / 1M tokens |

Costs are recorded in each essay's `.meta.yaml` and aggregated in the dashboard.

---

## Retry Logic

The Anthropic API client handles transient failures:
- Retryable conditions: credit balance, rate limit, overloaded, deadline exceeded
- Up to 30 attempts over a 1-hour deadline
- 2-minute wait between retries
- Dashboard callback on each retry (shows status)

---

## File Naming Conventions

| Type | Pattern | Example |
|------|---------|---------|
| Idea markdown | `ideas/{slug}.md` | `ideas/cheerios-effect.md` |
| Idea metadata | `ideas/{slug}.meta.yaml` | `ideas/cheerios-effect.meta.yaml` |
| Stage content | `{stage}/{slug}.md` | `draft/cheerios-effect.md` |
| Stage metadata | `{stage}/{slug}.meta.yaml` | `draft/cheerios-effect.meta.yaml` |
| Image source | `images/{slug}/{name}.{mermaid\|R\|ai-prompt.txt}` | `images/cheerios-effect/meniscus-diagram.mermaid` |
| Image output | `images/{slug}/{name}.png` | `images/cheerios-effect/meniscus-diagram.png` |
| Image meta | `images/{slug}/{name}.meta.yaml` | `images/cheerios-effect/meniscus-diagram.meta.yaml` |
| Exported essay | `export/cChapter - 2026 - {B}.{PP}.{OO} {Title}.docx` | `export/cChapter - 2026 - I.03.01 The Cheerios Effect.docx` |
| Exported section | `export/cSection - 2026 - {B}.{PP}.00 {Part Name}.docx` | `export/cSection - 2026 - I.03.00 What Your Kitchen Knows.docx` |
| Exported intro | `export/cChapter - 2026 - {B}.00.00 Introduction.docx` | `export/cChapter - 2026 - I.00.00 Introduction.docx` |
| Shared image | `images/shared/section-placeholder.png` | Placeholder "Insert cartoon here" for all sections |

---

## Project Structure

```
math/
├── Makefile                      # Builds pipeline, imagerender, imageswap
├── go.mod                        # Module: github.com/TrueBlocks/trueblocks-bookmill
├── cmd/
│   ├── pipeline/main.go          # Main orchestrator binary
│   ├── scaffold/main.go          # Scaffolding for essays, sections, and introductions
│   ├── imagerender/main.go       # Mermaid/R/AI image rendering
│   └── imageswap/main.go         # Embed PNGs into .docx files
├── internal/pipeline/
│   ├── api.go                    # Anthropic Claude API client
│   ├── arcs.go                   # 15 narrative arc definitions + 6 ending types
│   ├── config.go                 # YAML config loading
│   ├── dashboard.go              # Web dashboard HTTP server
│   ├── prompts.go                # AI prompt templates per stage
│   ├── runner.go                 # Cycle execution engine
│   └── state.go                  # Essay state management, disk I/O
├── data/
│   ├── common.R                  # Shared R theme and color palette
│   └── mermaid-theme.json        # Mermaid diagram styling
├── projects/math-books/
│   └── (generated stage folders)  # ideas/, research/, outline/, etc.
│   ├── ideas/                    # Essay, section, and introduction starter files
│   ├── research/                 # AI-generated research briefs
│   ├── outline/                  # AI-generated outlines
│   ├── draft/                    # AI-generated first drafts
│   ├── factcheck/                # AI-generated fact-checks
│   ├── illustrate/               # AI-generated illustration plans
│   ├── draft2/                   # AI-generated final revisions
│   ├── export/                   # Finished .docx files (cChapter + cSection)
│   └── images/                   # Per-essay image sources and PNGs + shared/
├── design/                       # Book planning documents (see below)
├── specs/                        # This spec and future specs
└── _backup/                      # Earlier versions of written essays
```

---

## Current State

As of March 2026:
- **94 essays** have completed the full pipeline — ideas through export
- **All 94 .docx files exist** in `export/`
- The pipeline, scaffold, imagerender, and imageswap tools are all functional
- A three-book reorganization is underway: Book I (~34 essays), Book II (~35 essays), Book III (~36 essays, of which ~11 are new)
- Section dividers (24) and book introductions (3) bring the planned total to ~132 items
- The **Plan for Book** documents (in `design/`) are the authoritative source for essay assignments, part names, and book arcs

---

## Related Specs

| Spec | Covers |
|------|--------|
| [narrative-arcs-and-endings.md](narrative-arcs-and-endings.md) | 15 arcs, 6 ending types, assignment rules, biographical enrichment |
| [section-dividers.md](section-dividers.md) | Part title pages, auto-skip mechanism, placeholder images |
| [book-introductions.md](book-introductions.md) | Book-level intro chapters, skip list, prompt guidance |

## What This Spec Does NOT Cover

This spec describes the **downstream pipeline**: given an idea, produce a .docx. It does not cover:

- **Idea generation and curation** — how essay topics are discovered, evaluated, and organized into books
- **Book-level structure** — how parts are themed, ordered, and balanced across volumes (see the Plan for Book documents in `design/` and the straw man in `design/Three-Book Straw Man.md`)
- **Series architecture** — the relationship between Books I, II, and III
