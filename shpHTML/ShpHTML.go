package shpHTML

import (
	"html"

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
