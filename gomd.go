// Copyright 2024 Alex Popescu.
// All rights reserved.
//
// gomd is a Markdown command line tool based on
// yuin/goildmark Markdown library.
//
// Usage:
package main

import (
	"flag"
	"fmt"
	"github.com/al3xandru/gomarkdown/criticalmarkdown"
	"github.com/google/uuid"
	figure "github.com/mangoumbrella/goldmark-figure"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/anchor"
	"go.abhg.dev/goldmark/frontmatter"
	"go.abhg.dev/goldmark/wikilink"
	"log"
	"os"
)

var (
	buildVersion = "1.0.0"
	buildSha     = "!"
	buildDate    = "!"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gomd [options] file\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func version() {
	fmt.Fprintf(os.Stdout, "gomd %sb%s@%s\n", buildVersion, buildSha, buildDate)
	os.Exit(0)
}

const disableBatteries = false

var flagVersion,
	noFrontmatter, noDefLists, noTables, noTasks, noFigures,
	noCriticalMarkdown, noFootnotes, noStrikethrough, noTypography, noWikilinks, noIds bool

func main() {
	// Configure logging for a command-line program.
	log.SetFlags(0)
	log.SetPrefix("hello: ")

	flag.BoolVar(&flagVersion, "version", false, "displays version and exits")
	flag.BoolVar(&noFrontmatter, "no-frontmatter", disableBatteries, "disables support for front-matter")
	flag.BoolVar(&noDefLists, "no-definition-lists", disableBatteries, "disables support for definition lists")
	flag.BoolVar(&noTables, "no-tables", disableBatteries, "disables support for tables")
	flag.BoolVar(&noTasks, "no-tasks", disableBatteries, "disables support for tasks")
	flag.BoolVar(&noFigures, "no-figures", disableBatteries, "disables suport for using figure instead of img")
	flag.BoolVar(&noFootnotes, "no-footnotes", disableBatteries, "disables footnote[^fn] processing")
	flag.BoolVar(&noStrikethrough, "no-strikethrough", disableBatteries, "disables ~~strikethrough processing")
	flag.BoolVar(&noTypography, "no-typography", disableBatteries, "disables typography/smartypants characters")
	flag.BoolVar(&noWikilinks, "no-wikilinks", disableBatteries, "disables [[wikilinks]] processing")
	flag.BoolVar(&noIds, "no-ids", disableBatteries, "disables generating header IDs")
	flag.Usage = usage

	// Parse flags.
	flag.Parse()
	if flagVersion {
		version()
	}
	if len(flag.Args()) == 0 {
		flag.Usage()
	}
	source, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file %s: %v\n",
			flag.Arg(0),
			err)
		panic(err)
	}

	footnotePrefix := uuid.NewString()[0:6]

	exts := make([]goldmark.Extender, 0)
	// block level
	if !noFrontmatter {
		exts = append(exts, &frontmatter.Extender{})
	}
	if !noDefLists {
		exts = append(exts, extension.DefinitionList)
	}
	if !noTables {
		exts = append(exts, extension.Table)
	}
	if !noTasks {
		exts = append(exts, extension.TaskList)
	}
	if !noFigures {
		exts = append(exts, figure.Figure)
	}
	// span level
	if !noCriticalMarkdown {
		exts = append(exts, criticalmarkdown.Extension)
	}
	if !noStrikethrough {
		exts = append(exts, extension.Strikethrough)
	}
	if !noFootnotes {
		exts = append(exts, extension.NewFootnote(extension.WithFootnoteIDPrefix([]byte(footnotePrefix))))
	}
	if !noTypography {
		exts = append(exts, extension.Typographer)
	}
	if !noWikilinks {
		exts = append(exts, &wikilink.Extender{})
	}
	if !noIds {
		exts = append(exts, &anchor.Extender{})
	}

	markdown := goldmark.New(
		goldmark.WithParserOptions(parser.WithAutoHeadingID()), // required by anchor
		goldmark.WithExtensions(exts...),
	)

	ctx := parser.NewContext()
	if err := markdown.Convert(source, os.Stdout, parser.WithContext(ctx)); err != nil {
		panic(err)
	}
}
