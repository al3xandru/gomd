package criticmarkup

import (
	"fmt"
	"github.com/yuin/goldmark/ast"
)

type AdditionNode struct {
	ast.BaseInline
}

// TODO implement Dump
func (n *AdditionNode) Dump(source []byte, level int) {
	fmt.Printf("AdditionNode.Dump: %s %d\n", source, level)
}

var KindAddition = ast.NewNodeKind("CriticMarkupAddition")

func (n *AdditionNode) Kind() ast.NodeKind {
	return KindAddition
}

func NewAdditionNode() *AdditionNode {
	return &AdditionNode{}
}
