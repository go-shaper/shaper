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
	UpCase := shaper.NewFilter().ApplyToUpper()
	LCase := shaper.NewFilter().ApplyToLower()
	Replace := shaper.NewFilter().ApplyReplace("test", "biscuit", -1)
	RU := shaper.NewFilter().ApplyReplace("test", "biscuit", -1).ApplyToUpper()

	// Test pipelines
	fmt.Printf("%s\n", UpCase.Process("This is a test."))
	fmt.Printf("%s\n", LCase.Process("This is a test."))
	fmt.Printf("%s\n", Replace.Process("This is a test."))
	fmt.Printf("%s\n", RU.Process("This is a test."))

	// Note that we can reuse these stacks as many times as we like
	fmt.Printf("%s\n", Replace.Process("This is also a test. Testificate."))

	// We can also add stages later on - though we cannot remove stages using this style
	Replace.ApplyToUpper()
	fmt.Printf("%s\n", Replace.Process("This is also a test. Testificate."))
	LCase.ApplyReplace("test", "biscuit", -1)
	fmt.Printf("%s\n", LCase.Process("This is also a test. Testificate."))

	// Regexp.ReplaceAll
	RegReplace := shaper.NewFilter().ApplyRegexpReplaceAll("(?i)ht(ml)", "X$1")
	fmt.Printf("%s\n", RegReplace.Process("This is html Html HTML."))

	// Test trim
	spTrim := shaper.NewFilter().ApplyTrim()
	stFrom := " \t\n a   long \t lone\t gopher \n\t\r\n"
	stTo := spTrim.Process(stFrom)
	fmt.Printf("F: %q\nT: %q\n", stFrom, stTo)
	spTrim.ApplyRegSpaces()
	stTo = spTrim.Process(stFrom)
	fmt.Printf("R: %q\n", stTo)

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
	// Finished.

}

////////////////////////////////////////////////////////////////////////////
// Extending shaper.Shaper

// Shaper extends shaper.Shaper
type Shaper struct {
	*shaper.Shaper
}

// NewFilter makes a new Shaper filter
func NewFilter() *Shaper {
	return &Shaper{Shaper: shaper.NewFilter()}
}

// ApplyHTMLUnescape will apply/add to html.UnescapeString filter to the Shaper
func (shpr *Shaper) ApplyHTMLUnescape() *Shaper {
	shpr.AddFilter(html.UnescapeString)
	return shpr
}

func ExampleShaper() {
	// == Extending shaper.Shaper to add your own filters
	var hu *Shaper
	hu = NewFilter()
	hu.ApplyHTMLUnescape()
	fmt.Printf("%s\n", hu.Process("2 &gt;= 1"))
	// Output:
	// 2 >= 1
}
