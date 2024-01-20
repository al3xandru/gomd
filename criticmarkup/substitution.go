package criticmarkup

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type subParser struct{}

var defaultSubParser = &subParser{}

func (p *subParser) Trigger() []byte {
	return []byte{'{'}
}

func (p *subParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	line, seg := block.PeekLine()
	if len(line) == 0 {
		return nil
	}
	if !bytes.HasPrefix(line, subOSeq) {
		return nil
	}
	if endIndx := bytes.Index(line, subESeq); endIndx > -1 {
		sepIdx := bytes.Index(line[0:endIndx], []byte("~>"))
		if sepIdx < 0 {
			return nil
		}

		block.Advance(endIndx + len(subESeq))

		delSeg := text.NewSegment(seg.Start+len(subOSeq), seg.Start+sepIdx)
		node := NewMarkupNode(KindDelete, delSeg)
		parent.AppendChild(parent, node)
		insSeg := text.NewSegment(seg.Start+sepIdx+2, seg.Start+endIndx)
		node = NewMarkupNode(KindAddition, insSeg)

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

func NewSubParser() parser.InlineParser {
	return defaultSubParser
}

type subExtension struct{}

// DeletionExtension supports CriticMarkup addition syntax
var SubExtension = &subExtension{}

func (e *subExtension) Extend(markdown goldmark.Markdown) {
	markdown.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(NewSubParser(), DefaultPriority)))
}
