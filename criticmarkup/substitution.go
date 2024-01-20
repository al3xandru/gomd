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

func (p *subParser) Parse(parent ast.Node, block text.Reader, _ parser.Context) ast.Node {
	line, seg := block.PeekLine()
	if len(line) == 0 {
		return nil
	}
	if !bytes.HasPrefix(line, substitutionStart) {
		return nil
	}
	endIdx := bytes.Index(line, substitutionClose)
	endSeg := seg

	if endIdx < 0 {
		// might be split across multiple lines
		for endIdx < 0 && bytes.HasSuffix(line, []byte{'\n'}) {
			block.Advance(len(line))
			line, endSeg = block.PeekLine()
			endIdx = bytes.Index(line, substitutionClose)
		}
	}
	if endIdx >= 0 {
		// look for separator; if none found the markup is broken
		splitAt := bytes.Index(line[0:endIdx], []byte("~>"))
		if splitAt < 0 {
			block.ResetPosition()
			return nil
		}

		block.Advance(endIdx + len(substitutionClose))

		delSeg := text.NewSegment(seg.Start+len(substitutionStart), endSeg.Start+splitAt)
		delNode := NewMarkupNode(KindDelete, delSeg)
		parent.AppendChild(parent, delNode)

		insSeg := text.NewSegment(endSeg.Start+splitAt+2, endSeg.Start+endIdx)
		insNode := NewMarkupNode(KindAddition, insSeg)

		return insNode
	} else {
		block.ResetPosition()
		return nil
	}
}

type subExtension struct{}

// SubExtension supports substitution markup from CriticMarkup
var SubExtension = &subExtension{}

func (e *subExtension) Extend(markdown goldmark.Markdown) {
	markdown.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(defaultSubParser, DefaultPriority)))
}
