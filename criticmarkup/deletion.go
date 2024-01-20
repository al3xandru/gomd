package criticmarkup

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var defaultDeleteParser = NewMarkupParser(deletionStart, deletionClose, KindDelete, nil)

var defaultDeleteRenderer = NewHTMLRenderer(KindDelete, "del")

type deletionExtension struct{}

// DeletionExtension supports CriticMarkup addition syntax
var DeletionExtension = &deletionExtension{}

func (e *deletionExtension) Extend(markdown goldmark.Markdown) {
	markdown.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(defaultDeleteParser, DefaultPriority)))

	markdown.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(defaultDeleteRenderer, DefaultPriority)))
}
