[![Build Status](https://drone.io/github.com/hgfischer/go-bfmt/status.png)](https://drone.io/github.com/hgfischer/go-bfmt/latest)

# bfmt

This is a package that wraps most functions present in the `fmt` package with a boolean conditional.

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
		fmt.Println("Be chatty...")
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

var verbose bfmt.Bool

func main() {
	flag.BoolVar((bool)(&verbose), "v", false, "Enable verbose output")
	flag.Parse()

	verbose.Println("Be chatty...")
}
```