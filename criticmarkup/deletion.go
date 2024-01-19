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

type deleteParser struct{}

var defaultDeleteParser = &deleteParser{}

func (p *deleteParser) Trigger() []byte {
	return []byte{'{'}
}

func (p *deleteParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, seg := block.PeekLine()
	if len(line) == 0 {
		return nil
	}
	if !bytes.HasPrefix(line, delStartSeq) {
		return nil
	}
	if endIndx := bytes.Index(line, delEndSeq); endIndx > -1 {
		block.Advance(endIndx + len(delEndSeq))

		seg = text.NewSegment(seg.Start+len(delStartSeq), seg.Start+endIndx)
		node := NewMarkupNode(KindDelete, seg)

		return node
	} else {
		// might be split across multiple lines
		multiLine := bytes.Clone(line)
		contseg := seg
		for endIndx < 0 && bytes.HasSuffix(line, []byte{'\n'}) {
			block.Advance(len(line))
			line, contseg = block.PeekLine()
			endIndx = bytes.Index(line, delEndSeq)
			multiLine = append(multiLine, line...)
		}
		if endIndx >= 0 {
			block.Advance(endIndx + len(delEndSeq))

			seg = text.NewSegment(seg.Start+len(delStartSeq), contseg.Start+endIndx)
			node := NewMarkupNode(KindDelete, seg)
			return node
		} else {
			block.ResetPosition()
			return nil
		}
	}
}

func NewDeleteParser() parser.InlineParser {
	return defaultDeleteParser
}

// DeleteHTMLRenderer is a renderer.NodeRenderer implementation that
// renders CriticMarkup addition nodes.
type DeleteHTMLRenderer struct {
	html.Config
}

// NewDeleteHTMLRenderer returns a new AdditionHTMLRenderer.
func NewDeleteHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &DeleteHTMLRenderer{
		Config: html.NewConfig(),
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *DeleteHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindDelete, r.renderDeletion)
}

// DelAdditionAttributeFilter defines attribute names which del elements can have.
var DelAdditionAttributeFilter = html.GlobalAttributeFilter

func (r *DeleteHTMLRenderer) renderDeletion(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		if n.Attributes() != nil {
			_, _ = w.WriteString("<del")
			html.RenderAttributes(w, n, InsAdditionAttributeFilter)
			_ = w.WriteByte('>')
		} else {
			_, _ = w.WriteString("<del>")
		}
	} else {
		_, _ = w.WriteString("</del>")
	}
	return ast.WalkContinue, nil
}

type deletionExtension struct{}

// DeletionExtension supports CriticMarkup addition syntax
var DeletionExtension = &deletionExtension{}

func (e *deletionExtension) Extend(markdown goldmark.Markdown) {
	markdown.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(NewDeleteParser(), DefaultPriority)))

	markdown.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewDeleteHTMLRenderer(), DefaultPriority)))
}
