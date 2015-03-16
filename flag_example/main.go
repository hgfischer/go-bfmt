package main

import (
	"flag"

	"github.com/hgfischer/go-bfmt"
)

var verbose = bfmt.Bool(false)

func main() {
	flag.Var(&verbose, "v", "Enable verbosity")
	flag.Parse()
	verbose.Println("I'm chatty")
}
