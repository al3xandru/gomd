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

type additionParser struct{}

var defaultAdditionParser = &additionParser{}

var (
	openSeq = []byte("{++")
	endSeq  = []byte("++}")
)

func (p *additionParser) Trigger() []byte {
	return []byte{'{'}
}

func (p *additionParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, seg := block.PeekLine()
	if len(line) == 0 {
		return nil
	}
	if !bytes.HasPrefix(line, openSeq) {
		return nil
	}
	if endIndx := bytes.Index(line, endSeq); endIndx > -1 {
		block.Advance(endIndx + len(endSeq))

		seg = text.NewSegment(seg.Start+len(openSeq), seg.Start+endIndx)
		node := NewAdditionNode()
		node.AppendChild(node, ast.NewTextSegment(seg))
		return node
	} else {
		// might be split across multiple lines
		multiLine := bytes.Clone(line)
		contseg := seg
		for endIndx < 0 && bytes.HasSuffix(line, []byte{'\n'}) {
			block.Advance(len(line))
			line, contseg = block.PeekLine()
			endIndx = bytes.Index(line, endSeq)
			multiLine = append(multiLine, line...)
		}
		if endIndx >= 0 {
			block.Advance(endIndx + len(endSeq))

			seg = text.NewSegment(seg.Start+len(openSeq), contseg.Start+endIndx)
			node := NewAdditionNode()
			node.AppendChild(node, ast.NewTextSegment(seg))
			return node
		} else {
			block.ResetPosition()
			return nil
		}
	}
}

func NewAdditionParser() parser.InlineParser {
	return defaultAdditionParser
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
			util.Prioritized(NewAdditionParser(), Priority)))

	markdown.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewAdditionHTMLRenderer(), Priority)))
}
