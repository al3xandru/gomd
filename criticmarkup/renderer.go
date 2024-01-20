package criticmarkup

import (
	"fmt"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type HTMLRenderer struct {
	html.Config
	Kind ast.NodeKind
	Tag  string
}

// NewHTMLRenderer returns a new AdditionHTMLRenderer.
func NewHTMLRenderer(kind ast.NodeKind, tag string, opts ...html.Option) renderer.NodeRenderer {
	r := &HTMLRenderer{
		Config: html.NewConfig(),
		Kind:   kind,
		Tag:    tag,
	}
	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}
	return r
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *HTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(r.Kind, r.renderHTML)
}

func (r *HTMLRenderer) renderHTML(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		if n.Attributes() != nil {
			_, _ = w.WriteString(fmt.Sprintf("<%s", r.Tag))
			html.RenderAttributes(w, n, html.GlobalAttributeFilter)
			_ = w.WriteByte('>')
		} else {
			_, _ = w.WriteString(fmt.Sprintf("<%s>", r.Tag))
		}
	} else {
		_, _ = w.WriteString(fmt.Sprintf("</%s>", r.Tag))
	}
	return ast.WalkContinue, nil
}
