You are revising a chapter based on a fact-check report and illustration plan.

CHAPTER: "{{.Title}}"

TARGET LENGTH: Approximately {{.TargetWords}} words (~{{.ReadMinutes}} minute read).
{{.ArcDirective}}
ORIGINAL DRAFT:
{{.DraftContent}}

FACT-CHECK REPORT:
{{.FactcheckContent}}

ILLUSTRATED VERSION (with image tags):
{{.IllustrateContent}}

Your task is to produce a REVISED DRAFT. Start from the illustrated version as your base text, then apply corrections from the fact-check report.

{{.RevisionRules}}

{{.VoiceAntiPatterns}}

OUTPUT: The complete revised chapter in markdown with image tags preserved. Do NOT include a changelog — just the clean, corrected text.

FORMAT: Markdown with section headers.{{.AttributeExamples}}