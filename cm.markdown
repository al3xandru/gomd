There are five additions from Critical Markdown:
first is {++an addition++},
followed by {--deletion of text --},
and {~~replacing text~>with substituting text~~}.

It also adds support for {>> comments that are visible <<}.
Last, but not least {==highlighting text==}{>>with comments<<}.

Some rules:

*   space inside the element must be part of the output
*   addition results in `<ins>`
*   deletion in `<del>`
*   substitutions in `<del>` followed by `<ins>`
*   comments in `<span class="critic comment">`
*   and highlights in `<mark>` (followed optionally by comment).

There are some rules about handling spaces, newlines, etc.
that I'll ignore initially.

How is an {++addition with
newlines behave
across three lines, but no new paragraphs++}. This is not clear.

Does ~~strike through
work across line~~?