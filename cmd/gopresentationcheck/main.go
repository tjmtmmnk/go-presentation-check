// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go-presentation-check calculates the cyclomatic complexities of functions and
// methods in Go source code.
//
// Usage:
//
//	go-presentation-check [<flag> ...] <Go file or directory> ...
//
// Flags:
//
//	-ignore REGEX         exclude files matching the given regular expression
//
// The output fields for each line are:
// <complexity> <package> <function> <file:line:column>
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	presentationcheck "github.com/tjmtmmnk/go-presentation-check"
)

const usageDoc = `Calculate cyclomatic complexities of Go functions.
Usage:
    go-presentation-check [flags] <Go file or directory> ...

Flags:
    -ignore REGEX         exclude files matching the given regular expression

The output fields for each line are:
<complexity> <package> <function> <file:line:column>
`

func main() {
	over := 1
	ignore := flag.String("ignore", "", "exclude files matching the given regular expression")

	log.SetFlags(0)
	log.SetPrefix("go-presentation-check: ")
	flag.Usage = usage
	flag.Parse()
	paths := flag.Args()
	if len(paths) == 0 {
		usage()
	}

	allStats := presentationcheck.Analyze(paths, regex(*ignore))
	shownStats := allStats.SortAndFilter(len(paths), over)

	printStats(shownStats)

	if len(shownStats) > 0 {
		os.Exit(1)
	}
}

func regex(expr string) *regexp.Regexp {
	if expr == "" {
		return nil
	}
	re, err := regexp.Compile(expr)
	if err != nil {
		log.Fatal(err)
	}
	return re
}

func printStats(s presentationcheck.Stats) {
	for _, stat := range s {
		fmt.Println(stat)
	}
}

func usage() {
	_, _ = fmt.Fprint(os.Stderr, usageDoc)
	os.Exit(2)
}
