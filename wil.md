# 2024-01-14 Sun

I don't want the `linkify` extension: it's used to recognized any
URLs in text (e.g. www.example.org).

# 2024-01-18 Thu

## How are extensions enabled?

`Extend(goldmark.Markdwon)`

```go
func (e *strikethrough) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewStrikethroughParser(), 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewStrikethroughHTMLRenderer(), 500),
	))
}
```