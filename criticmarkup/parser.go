package criticmarkup

import (
	"bytes"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type applyFunc func(node *MarkupNode) *MarkupNode

type markupParser struct {
	openSequence  []byte
	closeSequence []byte
	kind          ast.NodeKind
	apply         applyFunc
}

func noOpApply(node *MarkupNode) *MarkupNode {
	return node
}

func NewMarkupParser(open, close []byte, kind ast.NodeKind, fn applyFunc) *markupParser {
	if fn == nil {
		fn = noOpApply
	}
	return &markupParser{
		openSequence:  open,
		closeSequence: close,
		kind:          kind,
		apply:         fn,
	}
}

func (p *markupParser) Trigger() []byte {
	return []byte{'{'}
}

//

var (
	additionStart     = []byte("{++")
	additionClose     = []byte("++}")
	deletionStart     = []byte("{--")
	deletionClose     = []byte("--}")
	substitutionStart = []byte("{~~")
	substitutionClose = []byte("~~}")
	commentStart      = []byte("{>>")
	commentClose      = []byte("<<}")
	highlightStart    = []byte("{==")
	highlightClose    = []byte("==}")
)

func (p *markupParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, seg := block.PeekLine()
	if len(line) == 0 {
		return nil
	}
	if !bytes.HasPrefix(line, p.openSequence) {
		return nil
	}
	endIndex := bytes.Index(line, p.closeSequence)
	endSeg := seg

	if endIndex < 0 {
		// might be split across multiple lines
		for endIndex < 0 && bytes.HasSuffix(line, []byte{'\n'}) {
			block.Advance(len(line))
			line, endSeg = block.PeekLine()
			endIndex = bytes.Index(line, p.closeSequence)
		}
	}
	if endIndex >= 0 {
		block.Advance(endIndex + len(p.closeSequence))

		seg = text.NewSegment(seg.Start+len(p.openSequence), endSeg.Start+endIndex)
		node := NewMarkupNode(p.kind, seg)
		node = p.apply(node)

		return node
	} else {
		block.ResetPosition()
		return nil
	}

}
