package criticmarkup

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type addParser struct{}

var defaultaddParser = &addParser{}

func (p *addParser) Trigger() []byte {
	return []byte{'{'}
}

func (p *addParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, seg := block.PeekLine()
	if len(line) == 0 {
		return nil
	}
	if !bytes.HasPrefix(line, addStartSeq) {
		return nil
	}

	endIndex := bytes.Index(line, addEndSeq)
	endSeg := seg

	if endIndex < 0 {
		// might be split across multiple lines
		for endIndex < 0 && bytes.HasSuffix(line, []byte{'\n'}) {
			block.Advance(len(line))
			line, endSeg = block.PeekLine()
			endIndex = bytes.Index(line, addEndSeq)
		}
	}
	if endIndex >= 0 {
		seg = text.NewSegment(seg.Start+len(addStartSeq), endSeg.Start+endIndex)
		node := NewMarkupNode(KindAddition, seg)

		if node != nil {
			block.Advance(endIndex + len(addEndSeq))
		}
		return node
	} else {
		// the end markup was not found; need to tell the parser to get back to the original position
		block.ResetPosition()
		return nil
	}

}

func NewAddParser() parser.InlineParser {
	return defaultaddParser
}

// AdditionHTMLRenderer is a renderer.NodeRenderer implementation that
// renders CriticMarkup addition nodes.
type AdditionHTMLRenderer struct {
	html.Config
}

// NewAdditionHTMLRenderer returns a new AdditionHTMLRenderer.
func NewAdditionHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &AdditionHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *AdditionHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindAddition, r.renderAddition)
}

// InsAdditionAttributeFilter defines attribute names which ins elements can have.
var InsAdditionAttributeFilter = html.GlobalAttributeFilter

func (r *AdditionHTMLRenderer) renderAddition(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		if n.Attributes() != nil {
			_, _ = w.WriteString("<ins")
			html.RenderAttributes(w, n, InsAdditionAttributeFilter)
			_ = w.WriteByte('>')
		} else {
			_, _ = w.WriteString("<ins>")
		}
	} else {
		_, _ = w.WriteString("</ins>")
	}
	return ast.WalkContinue, nil
}

type additionExtension struct{}

// AdditionExtension supports CriticMarkup addition syntax
var AdditionExtension = &additionExtension{}

func (e *additionExtension) Extend(markdown goldmark.Markdown) {
	markdown.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(NewAddParser(), DefaultPriority)))

	markdown.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewAdditionHTMLRenderer(), DefaultPriority)))
}
