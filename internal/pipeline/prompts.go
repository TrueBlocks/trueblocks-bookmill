package pipeline

import "fmt"

func (r *Runner) researchPrompt(series, title, hook, hiddenMath, setting string) string {
	settingDirective := ""
	if setting != "" {
		settingDirective = fmt.Sprintf("\nSETTING CONTEXT: %s\nWhen choosing examples, historical context, and cultural references, ground them in this setting where natural. Don't force it — but let the setting inform the kind of examples you reach for.\n", setting)
	}
	return r.executePrompt(series, "research", map[string]any{
		"Title":            title,
		"Hook":             hook,
		"HiddenMath":       hiddenMath,
		"SettingDirective": settingDirective,
	})
}

func (r *Runner) outlinePrompt(series, title, researchContent string, targetWords int, arc NarrativeArc, structure StructureHint, entry EntryHint, mathVis MathVisHint) string {
	extraDirectives := ""
	if structure.OutlineHint != "" {
		extraDirectives += "\n" + structure.OutlineHint + "\n"
	}
	if entry.DraftHint != "" {
		extraDirectives += "\n" + entry.DraftHint + "\n"
	}
	if mathVis.OutlineHint != "" {
		extraDirectives += "\n" + mathVis.OutlineHint + "\n"
	}
	examples := r.buildExamples(series,
		"arcs/"+arc.Name,
		"structures/"+structure.Name,
		"entries/"+entry.Name,
		"mathvis/"+mathVis.Name,
	)
	return r.executePrompt(series, "outline", map[string]any{
		"Title":             title,
		"TargetWords":       targetWords,
		"ReadMinutes":       targetWords / 265,
		"ResearchContent":   researchContent,
		"ArcOutlineHint":    arc.OutlineHint,
		"ExtraDirectives":   extraDirectives,
		"AttributeExamples": examples,
	})
}

func (r *Runner) draftPrompt(series, title, outlineContent, researchContent string, targetWords int, arc NarrativeArc, structure StructureHint, entry EntryHint, register RegisterHint, setting string, mathVis MathVisHint) string {
	specs := r.specsFor(series)
	arcDirective := ""
	if arc.DraftHint != "" {
		arcDirective = fmt.Sprintf("\nNARRATIVE ARC: %s\n%s\n", arc.Label, arc.DraftHint)
	}
	if structure.DraftHint != "" {
		arcDirective += "\n" + structure.DraftHint + "\n"
	}
	if entry.DraftHint != "" {
		arcDirective += "\n" + entry.DraftHint + "\n"
	}
	if register.DraftHint != "" {
		arcDirective += "\n" + register.DraftHint + "\n"
	}
	if setting != "" {
		arcDirective += fmt.Sprintf("\nSETTING: %s\nGround the essay in this setting — the place, the time, the sensory world. Let the setting shape the examples and the imagery. Don't force it, but let it inform every choice.\n", setting)
	}
	if mathVis.DraftHint != "" {
		arcDirective += "\n" + mathVis.DraftHint + "\n"
	}
	examples := r.buildExamples(series,
		"arcs/"+arc.Name,
		"structures/"+structure.Name,
		"entries/"+entry.Name,
		"registers/"+register.Name,
		"mathvis/"+mathVis.Name,
	)
	return r.executePrompt(series, "draft", map[string]any{
		"Title":             title,
		"TargetWords":       targetWords,
		"ReadMinutes":       targetWords / 265,
		"ArcDirective":      arcDirective,
		"OutlineContent":    outlineContent,
		"ResearchContent":   researchContent,
		"VoiceProfile":      specs.VoiceProfile,
		"DraftRules":        specs.DraftRules,
		"AttributeExamples": examples,
	})
}

func (r *Runner) factcheckPrompt(series, title, draftContent, researchContent string) string {
	return r.executePrompt(series, "factcheck", map[string]any{
		"Title":           title,
		"DraftContent":    draftContent,
		"ResearchContent": researchContent,
	})
}

func (r *Runner) illustratePrompt(series, title, draftContent, factcheckContent, slug, setting string, mathVis MathVisHint) string {
	contextDirectives := ""
	if setting != "" {
		contextDirectives += fmt.Sprintf("\nSETTING: %s — let this inform the visual style and imagery choices.\n", setting)
	}
	if mathVis.DraftHint != "" {
		contextDirectives += "\n" + mathVis.DraftHint + "\nChoose image types and density accordingly.\n"
	}
	return r.executePrompt(series, "illustrate", map[string]any{
		"Title":             title,
		"Slug":              slug,
		"ContextDirectives": contextDirectives,
		"DraftContent":      draftContent,
		"FactcheckContent":  factcheckContent,
	})
}

func (r *Runner) draft2Prompt(series, title, draftContent, factcheckContent, illustrateContent string, targetWords int, arc NarrativeArc, register RegisterHint) string {
	specs := r.specsFor(series)
	arcDirective := ""
	if arc.DraftHint != "" {
		arcDirective = fmt.Sprintf("\nNARRATIVE ARC: %s\n%s\n", arc.Label, arc.DraftHint)
	}
	if register.DraftHint != "" {
		arcDirective += "\n" + register.DraftHint + "\nEnforce this register throughout the revision. If the original draft drifts into a different tone, correct it.\n"
	}
	examples := r.buildExamples(series,
		"arcs/"+arc.Name,
		"registers/"+register.Name,
	)
	return r.executePrompt(series, "draft2", map[string]any{
		"Title":             title,
		"TargetWords":       targetWords,
		"ReadMinutes":       targetWords / 265,
		"ArcDirective":      arcDirective,
		"DraftContent":      draftContent,
		"FactcheckContent":  factcheckContent,
		"IllustrateContent": illustrateContent,
		"RevisionRules":     specs.RevisionRules,
		"VoiceAntiPatterns": specs.VoiceAntiPatterns,
		"AttributeExamples": examples,
	})
}

func (r *Runner) sectionDraft2Prompt(series, title, ideaContent, partTitle string) string {
	return r.executePrompt(series, "section-draft2", map[string]any{
		"Title":       title,
		"PartTitle":   partTitle,
		"IdeaContent": ideaContent,
	})
}

func (r *Runner) introOutlinePrompt(series, title, ideaContent string) string {
	return r.executePrompt(series, "intro-outline", map[string]any{
		"Title":       title,
		"IdeaContent": ideaContent,
	})
}

func (r *Runner) introDraftPrompt(series, title, outlineContent, ideaContent string) string {
	return r.executePrompt(series, "intro-draft", map[string]any{
		"Title":          title,
		"OutlineContent": outlineContent,
		"IdeaContent":    ideaContent,
	})
}

func (r *Runner) introDraft2Prompt(series, title, draftContent string) string {
	return r.executePrompt(series, "intro-draft2", map[string]any{
		"Title":        title,
		"DraftContent": draftContent,
	})
}
