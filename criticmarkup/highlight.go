package criticmarkup

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var defaultHighlightParser = NewMarkupParser(higOSeq, higESeq, KindHighlight, nil)

var defaultHighlightRenderer = NewHTMLRenderer(KindHighlight, "mark")

type highlightExtension struct{}

// HighlightExtension supports the highlight syntax in Critic Markup
var HighlightExtension = &highlightExtension{}

func (e *highlightExtension) Extend(markdown goldmark.Markdown) {
	markdown.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(defaultHighlightParser, DefaultPriority)))

	markdown.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(defaultHighlightRenderer, DefaultPriority)))

}
