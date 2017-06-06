# bfmt

[![Build Status](https://travis-ci.org/hgfischer/go-bfmt.svg?branch=master)](https://travis-ci.org/hgfischer/go-bfmt)

This is a package that wraps most functions present in the `fmt` package with a boolean conditional.

Docs can be found at http://godoc.org/github.com/hgfischer/go-bfmt

## Introduction

If you ever needed to write code like this:

```
package main

import (
	"fmt"
	"flag"
)

var verbose bool

func main() {
	flag.BoolVar(&verbose, "v", false, "Enable verbose output")
	flag.Parse()

	if verbose {
		fmt.Println("I'm chatty...")
	}
}
```

Then you can now use `go-bfmt` to do this:

```
package main

import (
	"flag"
	"github.com/hgfischer/go-bfmt"
)

var verbose = bfmt.Bool(false)

func main() {
	flag.Var(&verbose, "v", "Enable verbose output")
	flag.Parse()

	verbose.Println("I'm chatty...")
}
```