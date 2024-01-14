---
front-matter section
---

# Example Markdown File

## Header H2

A short paragraph.

Underlines header H2
--------------------

This section has a code block:

```go
md := goldmark.New(
    goldmark.WithExtensions(extension.GFM),
        goldmark.WithParserOptions(
        parser.WithAutoHeadingID(),
    ),
    goldmark.WithRendererOptions(
        html.WithHardWraps(),
        html.WithXHTML(),
    ),
)
```

### Header H3

This is the paragraph with span elements. 
It starts with **bold**, followed by _italic_,
then another *italic*.

The next paragraph contains a ~~strikethrough~~
that prefixes a `monospace`.

#### Section H4

This is the typography section which tests the following elements:

- 'single quotes'
- "double quotes"
- "elipsis..."
- --EnDash
- ---EmDash

### H3 two

In this section we have a couple of line separators:

---

***

- - -

* * *

It is expected to see 4 horizontal lines above.

### Critical Markdown Section

### Footnotes

This paragraph uses footnotes[^1] 
which are an extension of Markdown.

[^1]: Still, even John Gruber,
the creator of Markdown,
uses them.
