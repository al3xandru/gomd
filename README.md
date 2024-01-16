# A Markdown CLI with All Batteries Included

This Markdown CLI is built using <https://github.com/yuin/goldmark>.
I've learned about [Goldmark](https://github.com/yuin/goldmark),
a Markdown parser written in Go,
from [Dan North](https://hachyderm.io/@tastapod@mastodon.social)
and decided to add it to my toolchain.
(This happened after me sharing 
that I've been using Perl 
to add support for footnotes in [Markdown.pl](https://daringfireball.net/projects/markdown/).)

The route I've taken with this tool was
to enable by default all extensions I normally use
and provide flags for disabling them.

The following extensions are enabled:

*   block level
    *   front matter (in YAML or TOML format)
    *   tasks `- [ ] a task`
    *   definition lists
    *   tables
    *   figures (replaces `img` with `figure`)
*   span level
    *   footnotes (footnote`[^1]`)
    *   strikethrough (`~~striked~~`)
    *   wikilinks (`[[example.markdown]]`)
    *   typography (transforms single, double quotes, elipsis into HTML entities)
    *   header IDs


## Build

On a macOS, run `make` and it will build for both architectures (ARM, Intel).

