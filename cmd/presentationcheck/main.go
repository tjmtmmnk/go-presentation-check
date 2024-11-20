package main

import (
	presentationcheck "github.com/tjmtmmnk/go-presentation-check"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(presentationcheck.Analyzer)
}
