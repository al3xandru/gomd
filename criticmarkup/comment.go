package criticmarkup

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

var defaultCommentParser = NewMarkupParser(comOSeq, comESeq, KindComment, func(node *MarkupNode) *MarkupNode {
	node.SetAttributeString("class", []byte("critic comment"))
	return node
})

func NewCommentParser() parser.InlineParser {
	return defaultCommentParser
}

// CommentHTMLRenderer is a renderer.NodeRenderer implementation that
// renders CriticMarkup addition nodes.
type CommentHTMLRenderer struct {
	html.Config
}

// NewCommentHTMLRenderer returns a new AdditionHTMLRenderer.
func NewCommentHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &CommentHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *CommentHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindComment, r.renderAddition)
}

func (r *CommentHTMLRenderer) renderAddition(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		if n.Attributes() != nil {
			_, _ = w.WriteString("<span")
			html.RenderAttributes(w, n, InsAdditionAttributeFilter)
			_ = w.WriteByte('>')
		} else {
			_, _ = w.WriteString("<span>")
		}
	} else {
		_, _ = w.WriteString("</span>")
	}
	return ast.WalkContinue, nil
}

type commentExtension struct{}

// CommentExtension supports CriticMarkup addition syntax
var CommentExtension = &commentExtension{}

func (e *commentExtension) Extend(markdown goldmark.Markdown) {
	markdown.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(NewCommentParser(), DefaultPriority)))

	markdown.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewCommentHTMLRenderer(), DefaultPriority)))
}
