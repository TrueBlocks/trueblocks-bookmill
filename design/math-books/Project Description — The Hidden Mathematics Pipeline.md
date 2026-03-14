# Project Description: The Hidden Mathematics Pipeline

## What This Is

An automated content pipeline that transforms two-sentence essay ideas into fully illustrated, fact-checked, print-ready Word documents — at scale. Built to produce *The Hidden Mathematics Series*, a three-book collection of popular mathematics essays.

The system is written in Go (~4,300 lines, single external dependency), runs on a laptop, and uses the Anthropic Claude API for all generative stages. Total production cost for 132 items across three books: under $90.

---

## The Books

*The Hidden Mathematics Series* finds the mathematics hiding inside everyday experiences — soap bubbles, piano tuning, cereal floating in a bowl, a shower curtain blowing inward — and makes it visible.

| Book | Title | Theme | Essays |
|------|-------|-------|--------|
| I | *Everything Is a Rate of Change* | The derivative chain as organizing metaphor; a journey from the intimate to the cosmic | ~34 |
| II | *The Edge of Knowing* | Frontiers where math breaks down, probability defies intuition, answers remain unknown | ~35 |
| III | *The Hidden Architecture* | Deep structural patterns in nature, art, music, and thought | ~36 |

Each book is organized into themed parts (8 per book), with section dividers and a book-level introduction. Total: 132 pipeline items (94 essays + 24 sections + 3 introductions + 11 new essays).

---

## The Eight-Stage Pipeline

Every essay passes sequentially through eight stages. At each stage, two files are produced: a `.md` (content) and a `.meta.yaml` (metadata tracking status, model, tokens, cost, timestamps).

```
ideas → research → outline → draft → factcheck → illustrate → draft2 → export
```

**1. Ideas** — Two-sentence seeds created by the `scaffold` tool: an everyday hook and a hidden mathematical punchline. Example: "The pond is half-empty one day before it's full" / "Humans linearize growth; exponentials ambush us every time."

**2. Research** — A 2,000-3,000 word brief covering mathematical foundations, historical context, worked examples, surprising connections, and citations.

**3. Outline** — The essay is structured according to an editorially assigned narrative arc (one of 15). The outline specifies section titles, key facts, "turns" (surprise moments), and a target word count sampled from a Gaussian distribution.

**4. Draft** — A complete first draft in an Asimov-meets-Mary Roach voice. Warm, conversational, precise. Every equation preceded by intuition and followed by interpretation. Concrete examples with real numbers.

**5. Fact-Check** — The draft is verified against the research. Factual errors, misleading simplifications, wrong attributions, and tone issues are quoted and corrected. Each essay receives a verdict: PASS, REVISE, or FAIL.

**6. Illustrate** — 1-2 images are planned per essay (one hero image mandatory). The system generates source code in three formats: Mermaid diagrams, R plots (ggplot2), or DALL-E prompts. Sources are rendered to PNG and embedded in the document.

**7. Draft2** — A final revision incorporating all fact-check corrections and naturally integrating image references into the prose. The narrative arc is maintained throughout.

**8. Export** — Mechanical conversion (no AI): markdown to Word via Pandoc, image embedding via custom tooling, style application via AppleScript automation with Microsoft Word. Output: a print-ready .docx formatted for a 6x9 trade paperback with mirror margins.

---

## The Narrative Framework

This is not a template system. Each essay is editorially assigned a **narrative arc** and an **ending type** based on its topic and position in the book.

### 15 Narrative Arcs

| Arc | Structure |
|-----|-----------|
| The Slow Build | Quiet scene → layers accumulate → late reveal → return transformed |
| The Cold Open Mystery | Baffling fact → investigate → eliminate suspects → crack the case |
| The Paradox | Impossible result → respect disbelief → build tools → intuition updates |
| The Demolition | Sell the myth → systematically dismantle → build the real answer |
| The Dual Timeline | Interleave historical discovery with modern understanding → converge |
| The Zoom | Journey through scales → same principle at every level → see the unity |
| The Argument | Competing explanations → give each its best case → jury reaches verdict |
| The Catalog of Wonders | Escalating variations on one theme → each more surprising than the last |
| The Confession | "I used to think X" → why X made sense → evidence against → new understanding |
| The Letter | Direct address to a curious person → patient explanation → shared discovery |
| The Walk | Physical journey through a place → each stop reveals mathematics |
| The Countdown | Start with the answer → peel back layers → drill to mathematical bedrock |
| The Recipe | Actionable instructions → mathematics emerges from doing |
| The Debate | Two historical figures argue → both compelling → evidence breaks the tie |
| The Inheritance | Idea passed through generations → each thinker transforms it |

### 6 Ending Types

Resolution, Provocation, Awe, Honesty, Communion, Handoff — each shapes how the essay lands emotionally.

### 7 Variation Attributes Per Essay

Arc, ending, structure hint, entry point, register (tone), setting, math visibility — all stored in YAML and injected into prompts.

---

## The Four Binaries

| Binary | Purpose |
|--------|---------|
| `pipeline` | Main orchestrator. Long-running process: cycles every ~15 seconds, calls Claude API concurrently, manages state, serves a web dashboard at port 8787 |
| `scaffold` | One-shot generator. Parses "Plan for Book" markdown tables and creates idea files + metadata for all essays, sections, and introductions |
| `imagerender` | Renders Mermaid diagrams, R plots, and DALL-E images to PNG. Includes AI-assisted repair (up to 3 attempts on syntax failures) |
| `imageswap` | Embeds rendered PNGs into .docx files by parsing Open XML and inserting proper `<w:drawing>` elements |

All binaries share the `internal/pipeline/` package. Built via `make` with a single dependency (`gopkg.in/yaml.v3`).

---

## Operational Features

**Concurrent processing** — A configurable semaphore limits simultaneous API calls (default: 3). Continuations (in-progress essays) are prioritized over new starts.

**Live configuration** — The YAML config file is re-read every cycle. Model assignments, concurrency limits, word-count targets, and cycle intervals can be tuned without restarting.

**Cost tracking** — Per-essay metadata records input tokens, output tokens, and cost at each stage. The dashboard aggregates session totals.

**Retry logic** — Transient API failures (rate limits, overloaded, credit balance) are retried up to 30 times over a 1-hour deadline with 2-minute intervals. Dashboard shows retry status in real time.

**Smart skip** — Section dividers and book introductions flow through the same pipeline but auto-skip irrelevant stages (e.g., sections skip research, outline, draft, factcheck, illustrate). Content is copied forward with zero cost.

**Web dashboard** — Real-time HTML UI showing per-project progress, cycle count, cost, essay status, log viewer, and manual step trigger.

**Image pipeline** — Three rendering methods with a consistent 7-color palette. Mermaid and R failures trigger Claude-assisted code repair. All images are embedded as proper Open XML drawings with correct relationships and content types.

**Word template upgrade** — AppleScript automation opens each .docx in Microsoft Word, copies styles from a .dotm template, sets 6x9 page layout with mirror margins and gutter, and scales all images to page width.

---

## Project Structure

```
math/
├── cmd/
│   ├── pipeline/main.go        # Main orchestrator
│   ├── scaffold/main.go        # Essay/section/intro scaffolding
│   ├── imagerender/main.go     # Mermaid/R/AI image rendering
│   └── imageswap/main.go       # PNG embedding in .docx
├── internal/pipeline/
│   ├── api.go                  # Claude API client with retry
│   ├── arcs.go                 # 15 narrative arcs + 6 endings
│   ├── config.go               # YAML config with live reload
│   ├── dashboard.go            # Web dashboard HTTP server
│   ├── prompts.go              # Stage-specific prompt templates
│   ├── runner.go               # Cycle execution engine
│   ├── state.go                # Essay state machine + disk I/O
│   └── variation.go            # Editorial attribute loading
├── data/
│   ├── common.R                # Shared R theme + color palette
│   ├── mermaid-theme.json      # Mermaid diagram styling
│   └── STYLE_GUIDE.md          # Visual design spec
├── projects/math-books/
│   ├── ideas/                  # Essay seeds
│   ├── research/               # Research briefs
│   ├── outline/                # Structured outlines
│   ├── draft/                  # First drafts
│   ├── factcheck/              # Fact-check reports
│   ├── illustrate/             # Illustration plans
│   ├── draft2/                 # Final revisions
│   ├── export/                 # Print-ready .docx files
│   └── images/                 # Per-essay image sources + PNGs
├── design/                     # Book plans, arcs, editorial docs
├── specs/                      # Pipeline and feature specifications
└── Makefile                    # Build four binaries
```

---

## By the Numbers

| Metric | Value |
|--------|-------|
| Total essays | 132 |
| Final word count | 141,062 |
| Tokens processed | ~12.4 million |
| Total API cost | $86.48 |
| Cost per essay | $0.66 |
| Go source lines | ~4,300 |
| External dependencies | 1 (yaml.v3) |
| Narrative arcs | 15 |
| Ending types | 6 |
| Image rendering methods | 3 |
| Concurrent API calls | 3 (configurable) |
| Cycle interval | 15 seconds (configurable) |

---

## Technology Stack

- **Language**: Go 1.25
- **AI**: Anthropic Claude API (essay generation), OpenAI DALL-E 3 (image generation)
- **Document conversion**: Pandoc (markdown to .docx)
- **Image rendering**: mmdc (Mermaid), Rscript + ggplot2 (R plots), DALL-E 3 (AI images)
- **Word automation**: AppleScript + Microsoft Word (style application, page layout, image scaling)
- **Configuration**: YAML
- **Web dashboard**: Go net/http (no frameworks)
