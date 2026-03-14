You are revising a fiction chapter for "The Extra Dwarfs," based on a consistency report and illustration plan.

CHAPTER: "{{.Title}}"

TARGET LENGTH: Approximately {{.TargetWords}} words (~{{.ReadMinutes}} minute read).
{{.ArcDirective}}
ORIGINAL DRAFT:
{{.DraftContent}}

CONSISTENCY REPORT:
{{.FactcheckContent}}

ILLUSTRATED VERSION (with image tags):
{{.IllustrateContent}}

Your task is to produce a REVISED DRAFT. Start from the illustrated version as your base text, then apply corrections from the consistency report. Specifically:

{{.RevisionRules}}

{{.VoiceAntiPatterns}}

FICTION-SPECIFIC REVISION RULES:
- Tighten dialogue. Cut any line that exists only for exposition.
- Strengthen subtext. If a character says exactly what they mean, rewrite.
- Ground scenes in physical detail. If a scene could happen anywhere, add specificity.
- Check pacing. Scenes should breathe — but dead space should be cut.
- Ensure the emotional arc of the chapter lands. The last paragraph should resonate.

OUTPUT: The complete revised chapter in markdown with image tags preserved. Do NOT include a changelog — just the clean, corrected chapter.

FORMAT: Markdown with section headers. Use **bold** for emphasis, *italics* for internal thought and titles.{{.AttributeExamples}}