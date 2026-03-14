You are writing a section divider page for a popular mathematics book.

SECTION: "{{.Title}}"
PART TITLE: {{.PartTitle}}

IDEA FILE:
{{.IdeaContent}}

Write a very short introductory paragraph for this section divider — 2 to 3 sentences maximum.
The paragraph sets up what this part of the book explores, in a warm, inviting voice. It does not
summarize the essays — it invites the reader in. Think of it as a breath between chapters, not an
announcement.

After the paragraph, include this image tag on its own line:

[[IMG:section-placeholder.png|Insert cartoon here]]

REQUIREMENTS:
- Maximum 3 sentences
- Warm, conversational tone — consistent with the book's character
- Do NOT list essays or topics explicitly
- Do NOT use phrases like "In this section..." or "The following chapters..."
- The reader should feel curiosity, not obligation

OUTPUT: A markdown document with:
1. A level-1 heading with the part title
2. The 2-3 sentence introductory paragraph
3. The image tag

Example format:
# What Your Kitchen Knows

Your kitchen is a physics laboratory with terrible safety protocols.
Every morning, your cereal, your coffee, and your shower conspire to
demonstrate graduate-level mathematics — without your permission.

[[IMG:section-placeholder.png|Insert cartoon here]]