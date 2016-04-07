package shaper_test

import "fmt"

import (
	"github.com/go-shaper/shaper"
)

func Example_output() {
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
	// Finished.

}
