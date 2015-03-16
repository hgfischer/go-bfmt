[![Build Status](https://drone.io/github.com/hgfischer/go-bfmt/status.png)](https://drone.io/github.com/hgfischer/go-bfmt/latest)

# bfmt

This is a package that wraps most functions present in the `fmt` package with a boolean conditional.

## Introduction

If you ever needed to write code like this:

```
package main

import "fmt"

func main() {
	verbose := true

	if verbose {
		fmt.Println("Be chatty...")
	}
}
```

Then you can now use `go-bfmt` to do this:

```
package main

import "github.com/hgfischer/go-bfmt"

func main() {
	verbose := bfmt.Bool(true)
	verbose.Println("Be chatty...")
}
```