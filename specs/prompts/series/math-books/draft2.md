You are revising a popular mathematics essay for the Hidden Mathematics series, based on a fact-check report and illustration plan.

ESSAY: "{{.Title}}"

TARGET LENGTH: Approximately {{.TargetWords}} words (~{{.ReadMinutes}} minute read). If the original draft is longer than this target, trim it down — cut less essential examples, tighten prose, remove tangents. Every paragraph must earn its place.
{{.ArcDirective}}
ORIGINAL DRAFT:
{{.DraftContent}}

FACT-CHECK REPORT:
{{.FactcheckContent}}

ILLUSTRATED VERSION (with image tags):
{{.IllustrateContent}}

Your task is to produce a REVISED DRAFT. Start from the illustrated version as your base text (it already has image tags placed), then apply corrections from the fact-check report. When corrections conflict with word count targets, accuracy wins. Specifically:

{{.RevisionRules}}

{{.VoiceAntiPatterns}}

OUTPUT: The complete revised essay in markdown with image tags preserved. Do NOT include a changelog or notes about what you changed — just the clean, corrected essay ready for human review.

FORMAT: Markdown with section headers. Use **bold** for emphasis, *italics* for book/paper titles and foreign words. Use $...$ for inline math and $$...$$ for display math. Target approximately {{.TargetWords}} words (~{{.ReadMinutes}} minute read).{{.AttributeExamples}}