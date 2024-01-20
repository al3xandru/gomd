package criticmarkup

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var defaultAdditionParser = NewMarkupParser(additionStart, additionClose, KindAddition, nil)

var defaultAdditionRenderer = NewHTMLRenderer(KindAddition, "ins")

type additionExtension struct{}

// AdditionExtension supports CriticMarkup addition syntax
var AdditionExtension = &additionExtension{}

func (e *additionExtension) Extend(markdown goldmark.Markdown) {
	markdown.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(defaultAdditionParser, DefaultPriority)))

	markdown.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(defaultAdditionRenderer, DefaultPriority)))
}
