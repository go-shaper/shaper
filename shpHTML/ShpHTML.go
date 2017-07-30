////////////////////////////////////////////////////////////////////////////
// Package: shpHTML
// Purpose: HTML shaper
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

/*

Package shpHTML provides string shaping functionalities for HTML.

*/
package shpHTML

import (
	"html"
	"regexp"

	"github.com/go-shaper/shaper"
)

////////////////////////////////////////////////////////////////////////////
// Extending shaper.Shaper

// ShpHTML extends shaper.Shaper
type ShpHTML struct {
	*shaper.Shaper
}

// NewShaper makes a new ShpHTML filter
func NewShaper() *ShpHTML {
	return &ShpHTML{Shaper: shaper.NewShaper()}
}

// ApplyHTMLUnescape will apply/add to html.UnescapeString filter to the ShpHTML
func (shpr *ShpHTML) ApplyHTMLUnescape() *ShpHTML {
	shpr.AddShaper(html.UnescapeString)
	return shpr
}

// ApplyTagRemove will apply/add a simple html tag removal filter to the Shaper
func (shaper *ShpHTML) ApplyTagRemove() *ShpHTML {
	shaper.AddShaper(func(s string) string {
		return regexp.MustCompile(`<.+?>`).ReplaceAllString(s, "")
	})
	return shaper
}
