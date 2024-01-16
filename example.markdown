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


### Footnotes

This paragraph uses footnotes[^1] 
which are an extension of Markdown.

[^1]: Still, even John Gruber,
the creator of Markdown,
uses them.

### Tasks section

- [ ] Add some Critical Markdown text
 - [ ] With a child task
- [x] Mark one of these as done
 
### Images

There are 2 types of images:

![no alt](https://64.media.tumblr.com/51065dfde5563682bea1e6e1cd21348c/fa295057f1c6a349-24/s540x810/2365757f68b5f2184e705bb263b64a20ab5c29cf.jpg)

And one with `figure`:

![with figure](https://64.media.tumblr.com/51065dfde5563682bea1e6e1cd21348c/fa295057f1c6a349-24/s540x810/2365757f68b5f2184e705bb263b64a20ab5c29cf.jpg)
The figcapture is here.

### Critical Markdown Section

There are five additions from Critical Markdown: 
first is {++ an addition ++},
followed by {-- deletion of text --},
and {~~ replacing text ~> with substituting text ~~}.

It also adds support for {>> comments that are visible <<}.
Last, but not least {== highlighting text ==}{>> with att

