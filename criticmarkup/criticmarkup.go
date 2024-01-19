// Package criticmarkup is a goldmark extension providing support
// for CriticMarkup syntax https://github.com/CriticMarkup/CriticMarkup-toolkit
// This extension doesn't work across paragraphs,
// i.e. using CriticMarkup to add/remove paragraphs will not work.
// CriticMarkup provides support for inline text editing
// rather than paragraph or structural editing.
package criticmarkup

import (
	"github.com/yuin/goldmark"
)

// Priority uses same value 500 as extension.Strikethrough
const Priority = 500

type criticMarkupExtension struct {
}

var Extension = &criticMarkupExtension{}

func (e *criticMarkupExtension) Extend(markdown goldmark.Markdown) {
	AdditionExtension.Extend(markdown)
}