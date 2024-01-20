package criticmarkup

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var defaultCommentParser = NewMarkupParser(commentStart, commentClose, KindComment, func(node *MarkupNode) *MarkupNode {
	node.SetAttributeString("class", []byte("critic comment"))
	return node
})

var defaultCommentRenderer = NewHTMLRenderer(KindComment, "span")

type commentExtension struct{}

// CommentExtension supports CriticMarkup addition syntax
var CommentExtension = &commentExtension{}

func (e *commentExtension) Extend(markdown goldmark.Markdown) {
	markdown.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(defaultCommentParser, DefaultPriority)))

	markdown.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(defaultCommentRenderer, DefaultPriority)))
}
