package criticmarkup

import (
	"fmt"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type MarkupNode struct {
	ast.BaseInline
	kind ast.NodeKind
}

var KindAddition = ast.NewNodeKind("cmAdd")
var KindDelete = ast.NewNodeKind("cmDelete")

func (n *MarkupNode) Kind() ast.NodeKind {
	return n.kind
}

func NewMarkupNode(kind ast.NodeKind, textSegment text.Segment) *MarkupNode {
	node := &MarkupNode{kind: kind}
	node.AppendChild(node, ast.NewTextSegment(textSegment))
	return node
}

func (n *MarkupNode) Dump(source []byte, level int) {
	fmt.Printf("Node(%s): %s %d\n", n.Kind(), source, level)
}
