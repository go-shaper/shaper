package shpHTML_test

import (
	"fmt"

	"github.com/go-shaper/shaper/shpHTML"
)

////////////////////////////////////////////////////////////////////////////
// How to use it with the ready-made filters

func Example() {
	// == Extending shaper.Shaper to add your own filters
	var hu *shpHTML.ShpHTML
	hu = shpHTML.NewShaper()
	hu.ApplyToUpper()
	fmt.Printf("%s\n", hu.Process("2 &gt;= 1"))
	hu.ApplyHTMLUnescape()
	fmt.Printf("%s\n", hu.Process("2 &gt;= 1"))
	// Output:
	// 2 &GT;= 1
	// 2 >= 1
}
