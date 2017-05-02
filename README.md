
# shaper

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-shaper/shaper?status.svg)](http://godoc.org/github.com/go-shaper/shaper)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-shaper/shaper)](https://goreportcard.com/report/github.com/go-shaper/shaper)
[![travis Status](https://travis-ci.org/go-shaper/shaper.svg?branch=master)](https://travis-ci.org/go-shaper/shaper)
[![codeship Status](https://codeship.com/projects/04245480-e7ff-0133-dc96-46bb3aa6b241/status?branch=master)](https://codeship.com/projects/147070)

## TOC
- [shaper - shape strings into desired forms](#shaper---shape-strings-into-desired-forms)
- [API](#api)
  - [> shaper_test.go](#-shaper_testgo)

# shaper - shape strings into desired forms

The `shaper` project provides general purpose building blocks to shape strings into desired forms. The cumulative building blocks can easily be chained together. They are used like building pipes in shell. 

Architected by Howard C. Shaw III, it needs neither go-routines nor channels. Instead, it creates a composable set of filters you could keep around and apply to a single string at a time. Behind the scenes, it is still building a composition of functions, but by currying that composition into a new function and holding the state of the stack in a struct, it restores the left-right ordering of the filters. Note that the filter stages are basically just compile-time freezes of a call to the currying function; i.e., no matter how complicated the shaping filter chain is, all are done at the compile time, so no run-time overhead when using `shaper`. 

# API

Check out the [provided examples](https://godoc.org/github.com/go-shaper/shaper#example-package--Output)  to see how to use it with the ready-made filters, and the [![Go Doc](https://godoc.org/github.com/go-shaper/shaper?status.svg)](https://godoc.org/github.com/go-shaper/shaper) document for further details.

To extend and add your own filters take a look at `func ExampleShaper` in:

#### > shaper_test.go
```go
package shaper_test

import (
	"fmt"
	"html"
)

import (
	"github.com/go-shaper/shaper"
)

////////////////////////////////////////////////////////////////////////////
// How to use it with the ready-made filters

func Example_output() {
	// == Using ready-made filters

	// Construct pipelines
	UpCase := shaper.NewShaper().ApplyToUpper()
	LCase := shaper.NewShaper().ApplyToLower()
	Replace := shaper.NewShaper().ApplyReplace("test", "biscuit", -1)

	// Test pipelines
	fmt.Printf("%s\n", UpCase.Process("This is a test."))
	fmt.Printf("%s\n", LCase.Process("This is a test."))
	fmt.Printf("%s\n", Replace.Process("This is a test."))

	// Demonstrating copy
	// to retain the `Replace` functionality but have another filter building
	// on top of it. Note: without the `Copy()`, the next call to Replace will
	// be affected by the ToUpper
	RU := Replace.Copy().ApplyToUpper()
	fmt.Printf("%s\n", RU.Process("This is a test."))

	// Note that we can reuse these stacks as many times as we like
	fmt.Printf("%s\n", Replace.Process("This is also a test. Testificate."))

	// We can also add stages later on - though we cannot remove stages using this style
	Replace.ApplyToUpper()
	fmt.Printf("%s\n", Replace.Process("This is also a test. Testificate."))
	LCase.ApplyReplace("test", "biscuit", -1)
	fmt.Printf("%s\n", LCase.Process("This is also a test. Testificate."))

	// Regexp.ReplaceAll
	RegReplace := shaper.NewShaper().ApplyRegexpReplaceAll("(?i)ht(ml)", "X$1")
	fmt.Printf("%s\n", RegReplace.Process("This is html Html HTML."))

	// Test trim
	spTrim := shaper.NewShaper().ApplyTrim()
	stFrom := " \t\n a   long \t lone\t gopher \n\t\r\n"
	stTo := spTrim.Process(stFrom)
	fmt.Printf("F: %q\nT: %q\n", stFrom, stTo)
	spTrim.ApplyRegSpaces()
	stTo = spTrim.Process(stFrom)
	fmt.Printf("R: %q\n", stTo)

	// Test ProcessAny
	fmt.Printf("%s\n", LCase.ProcessAny("This is also a test. Testificate."))
	fmt.Printf("%s\n",
		LCase.ProcessAny([]string{"This is also a test.", " Testificate."}))

	// == All done.
	fmt.Printf("Finished.\n")

	// Output:
	// THIS IS A TEST.
	// this is a test.
	// This is a biscuit.
	// THIS IS A BISCUIT.
	// This is also a biscuit. Testificate.
	// THIS IS ALSO A BISCUIT. TESTIFICATE.
	// this is also a biscuit. biscuitificate.
	// This is Xml Xml XML.
	// F: " \t\n a   long \t lone\t gopher \n\t\r\n"
	// T: "a   long \t lone\t gopher"
	// R: "a long lone gopher"
	// this is also a biscuit. biscuitificate.
	// this is also a biscuit. biscuitificate.
	// Finished.

}

////////////////////////////////////////////////////////////////////////////
// Extending shaper.Shaper

// Shaper extends shaper.Shaper
type Shaper struct {
	*shaper.Shaper
}

// NewShaper makes a new Shaper filter
func NewShaper() *Shaper {
	return &Shaper{Shaper: shaper.NewShaper()}
}

// ApplyHTMLUnescape will apply/add to html.UnescapeString filter to the Shaper
func (shpr *Shaper) ApplyHTMLUnescape() *Shaper {
	shpr.AddShaper(html.UnescapeString)
	return shpr
}

func ExampleShaper() {
	// == Extending shaper.Shaper to add your own filters
	var hu *Shaper
	hu = NewShaper()
	hu.ApplyToUpper()
	fmt.Printf("%s\n", hu.Process("2 &gt;= 1"))
	hu.ApplyHTMLUnescape()
	fmt.Printf("%s\n", hu.Process("2 &gt;= 1"))
	// Output:
	// 2 &GT;= 1
	// 2 >= 1
}
```


All patches welcome. 
