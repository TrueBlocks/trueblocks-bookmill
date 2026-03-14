You are structuring a chapter for a book.

CHAPTER: "{{.Title}}"

Below is the research brief. Using this material, create a detailed outline for a chapter of approximately {{.TargetWords}} words (~{{.ReadMinutes}} minute read).

RESEARCH:
{{.ResearchContent}}

{{.ArcOutlineHint}}

Create an outline. For each section, provide:
- A section title
- 2-3 sentence summary of what this section covers
- Key facts or examples to include
- Approximate word allocation (must total ~{{.TargetWords}})

Also include:

1. **How It Lands**
   - How does the chapter earn its ending?
   - What is the last emotional note?

2. **Tone Notes**
   - Conversational, engaging
   - Never condescending

FORMAT: Use markdown with clear section headers. Be specific.{{.ExtraDirectives}}{{.AttributeExamples}}