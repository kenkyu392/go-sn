# go-sn - Serial Number Generator

[![Go](https://github.com/kenkyu392/go-sn/actions/workflows/go.yml/badge.svg)](https://github.com/kenkyu392/go-sn/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/kenkyu392/go-sn/branch/main/graph/badge.svg)](https://codecov.io/gh/kenkyu392/go-sn)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-00ADD8?logo=go)](https://pkg.go.dev/github.com/kenkyu392/go-sn)
[![Go Report Card](https://goreportcard.com/badge/github.com/kenkyu392/go-sn)](https://goreportcard.com/report/github.com/kenkyu392/go-sn)
[![license](https://img.shields.io/github/license/kenkyu392/go-sn)](LICENSE)

Go library for generate serial numbers according to rules.

This library is useful for generating serial numbers in human-readable format and for generating dummy data (e.g., email addresses, names of people, etc.).

## Installation

```
go get -u github.com/kenkyu392/go-sn
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/kenkyu392/go-sn"
	"github.com/kenkyu392/go-sn/rule"
)

func main() {
	// Create a generator with arbitrary rules.
	g := sn.NewGenerator(
		// You can create your own rules.
		func() []rune {
			return []rune("ANIMALS")
		},
		rule.Mixin(
			rule.RandomUppercaseAlphabet(2),
			rule.RandomNumeric(2),
		),
		rule.RandomList(
			[]string{"Dog", "Cat", "Fox", "Rat"},
		),
		rule.RandomAlphabetAndNumeric(4),
		// Use rules for numbering and padding.
		// You can also do customized numbering.
		// -> (&rule.Counter{Start: 2, Size: 2, Padding: 4}).NumberingWithPadding(),
		rule.NumberingWithPadding(4),
	)

	// You can change the delimiter.
	g.Delimiter = "_"

	for i := 0; i < 5; i++ {
		fmt.Println(g.String())
	}
	/*
		ANIMALS_OH42_Fox_x8m0_0000
		ANIMALS_GJ55_Rat_4ZsV_0001
		ANIMALS_QE35_Fox_bRYu_0002
		ANIMALS_VN38_Dog_Lk9p_0003
		ANIMALS_LQ52_Fox_XGmb_0004
	*/
}
```

To generate a dummy email address, you may use the following:

```go
g := sn.NewGenerator(
	rule.Mixin(
		rule.RandomAlphabet(4),
		rule.String("."),
		rule.RandomAlphabetAndNumeric(6),
	),
	rule.RandomList(
		[]string{
			"example.com",
			"example.net",
			"example.org",
		},
	),
)
g.Delimiter = "@"
for i := 0; i < 5; i++ {
	fmt.Println(g.String())
}
/*
	Lbtr.99VmO8@example.net
	Cnat.GAgx19@example.org
	KvdO.p87x51@example.net
	UKHv.gtAME8@example.org
	YTCp.632C96@example.com
*/
```

## License

[MIT](LICENSE)
