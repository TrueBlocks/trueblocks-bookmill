You are revising a novel chapter based on a continuity report.

CHAPTER: "{{.Title}}"

TARGET LENGTH: Approximately {{.TargetWords}} words.

ORIGINAL DRAFT:
{{.DraftContent}}

CONTINUITY REPORT:
{{.ContinuityNotes}}

Your task is to produce a REVISED DRAFT. Start from the original draft, then apply all corrections from the continuity report.

REVISION RULES:
- Fix every issue flagged as REVISE or BROKEN severity
- Preserve the author's voice and style throughout
- Do not add new scenes or subplots — only fix what's broken
- Maintain or improve pacing; tighten prose where possible
- Ensure the chapter ending creates forward momentum
- Keep dialogue natural and character-consistent

{{.VoiceProfile}}

OUTPUT: The complete revised chapter in markdown. Do NOT include a changelog or notes — just the clean, corrected text.

FORMAT: Markdown with scene breaks indicated by `---`.