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
	"log"
	"os"

	"github.com/yuin/goldmark"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gomd [options] file\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	// Configure logging for a command-line program.
	log.SetFlags(0)
	log.SetPrefix("hello: ")

	// Parse flags.
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
	}
	source, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file %s: %v\n",
			flag.Arg(0),
			err)
		os.Exit(5)
	}
	if err := goldmark.Convert(source, os.Stdout); err != nil {
		panic(err)
	}
}
