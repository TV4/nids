# nids

[![Build Status](https://travis-ci.org/TV4/nids.svg?branch=master)](https://travis-ci.org/TV4/nids)
[![Go Report Card](https://goreportcard.com/badge/github.com/TV4/nids)](https://goreportcard.com/report/github.com/TV4/nids)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/TV4/nids)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/TV4/nids#license-mit)

The nids package is used to create slugs/tags.

## Installation

    go get -u github.com/TV4/nids

You probably want to vendor this package using your favorite vendoring tool.

# Configuration

Default             | WithÅÄÖ
--------------------|--------------------
`\A[0-9a-z-]{1,}\z` | `\A[0-9a-zåäö-]*\z`

The default configuration is meant to be useful for most use-cases, and the specialized WithÅÄÖ configuration
is meant to match the algorithm used in the [nid_utils](https://github.com/TV4/nid_utils) gem.

## Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/TV4/nids"
)

func main() {
	for i, s := range os.Args[1:] {
		if nids.Possible(s) {
			fmt.Printf("[%d] the string %q is already a nid.\n", i, s)
		} else {
			fmt.Printf("[%d] nid of %q is %q\n", i, s, nids.Case(s))
		}
	}
}
```

```bash
$ go run n.go 'Dürén Ibrahimović' 'Alvinnn!! & the Chipmunks' 'kale8^79_0-' foo-bar
[0] nid of "Dürén Ibrahimović" is "duren-ibrahimovic"
[1] nid of "Alvinnn!! & the Chipmunks" is "alvinnn-the-chipmunks"
[2] nid of "kale8^79_0-" is "kale879-0"
[3] the string "foo-bar" is already a nid.
```

## License (MIT)

Copyright (c) 2015-2018 TV4

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
