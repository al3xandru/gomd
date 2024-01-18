package criticalmarkdown

import (
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
)

// Priority uses same value 500 as extension.Strikethrough
const Priority = 500

type criticalMarkdown struct {
}

var Extension = &criticalMarkdown{}

func (e *criticalMarkdown) Extend(markdown goldmark.Markdown) {
	CMAddition.Extend(markdown)
}

type CMAdditionAST struct {
	ast.BaseInline
}

func (n *CMAdditionAST) Dump(source []byte, level int) {
	//TODO implement me
	fmt.Printf("CMAdditionAST.Dump: %s %d\n", source, level)
}

var KindCriticalMarkdownAddition = ast.NewNodeKind("CriticalMarkdownAddition")

func (n *CMAdditionAST) Kind() ast.NodeKind {
	return KindCriticalMarkdownAddition
}

func NewAddition() *CMAdditionAST {
	return &CMAdditionAST{}
}
