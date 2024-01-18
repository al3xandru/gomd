package criticalmarkdown

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

type cmAdditionParser struct{}

var defaultAdditionParser = &cmAdditionParser{}

var (
	openSeq = []byte("{++")
	endSeq  = []byte("++}")
)

func (p *cmAdditionParser) Trigger() []byte {
	return []byte{'{'}
}

func (p *cmAdditionParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
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
		node := NewAddition()
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
			node := NewAddition()
			node.AppendChild(node, ast.NewTextSegment(seg))
			return node
		} else {
			block.ResetPosition()
			return nil
		}
	}
}

func NewCMAdditionParser() parser.InlineParser {
	return defaultAdditionParser
}

// StrikethroughHTMLRenderer is a renderer.NodeRenderer implementation that
// renders Strikethrough nodes.
type StrikethroughHTMLRenderer struct {
	html.Config
}

// NewStrikethroughHTMLRenderer returns a new StrikethroughHTMLRenderer.
func NewStrikethroughHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &StrikethroughHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *StrikethroughHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindCriticalMarkdownAddition, r.renderStrikethrough)
}

// StrikethroughAttributeFilter defines attribute names which dd elements can have.
var StrikethroughAttributeFilter = html.GlobalAttributeFilter

func (r *StrikethroughHTMLRenderer) renderStrikethrough(
	w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		if n.Attributes() != nil {
			_, _ = w.WriteString("<ins")
			html.RenderAttributes(w, n, StrikethroughAttributeFilter)
			_ = w.WriteByte('>')
		} else {
			_, _ = w.WriteString("<ins>")
		}
	} else {
		_, _ = w.WriteString("</ins>")
	}
	return ast.WalkContinue, nil
}

type addition struct{}

var CMAddition = &addition{}

func (e *addition) Extend(markdown goldmark.Markdown) {
	markdown.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(NewCMAdditionParser(), Priority)))

	markdown.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewStrikethroughHTMLRenderer(), Priority)))
}
