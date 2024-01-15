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
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"log"
	"os"

	"github.com/yuin/goldmark"
	"go.abhg.dev/goldmark/anchor"
	"go.abhg.dev/goldmark/frontmatter"
	"go.abhg.dev/goldmark/wikilink"
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

func main() {
	// Configure logging for a command-line program.
	log.SetFlags(0)
	log.SetPrefix("hello: ")

	flagVersion := flag.Bool("version", false, "displays version and exits")

	// Parse flags.
	flag.Usage = usage
	flag.Parse()

	if *flagVersion {
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

	markdown := goldmark.New(
		goldmark.WithParserOptions(parser.WithAutoHeadingID()), // required by anchor
		goldmark.WithExtensions(extension.Footnote,
			extension.Strikethrough,
			extension.Typographer,
			extension.DefinitionList,
			extension.Table,
			extension.TaskList,
			&frontmatter.Extender{},
			&anchor.Extender{},
			&wikilink.Extender{}),
	)

	ctx := parser.NewContext()
	if err := markdown.Convert(source, os.Stdout, parser.WithContext(ctx)); err != nil {
		panic(err)
	}
}
