////////////////////////////////////////////////////////////////////////////
// Porgram: shaper.go
// Purpose: mold strings into shape
// authors: Tong Sun (c) 2016, All rights reserved
// Credits: Howard C. Shaw III
//          https://groups.google.com/d/msg/golang-nuts/snoIyANd-8c/V_IC57y4AwAJ
//          https://groups.google.com/d/msg/golang-nuts/snoIyANd-8c/hhOnu-lFAgAJ
////////////////////////////////////////////////////////////////////////////

/*

Package shaper is an universal string shaping library that can easily build
up string transformations step by step.

*/
package shaper

import (
	"regexp"
	"strings"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// Shaper struct holds the filter chain
type Shaper struct {
	ShaperStack func(string) string
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Function definitions

// PassThrough will return the string as-is, used primarily for NewShaper()
func PassThrough(s string) string {
	return s
}

// NewShaper makes a new Shaper filter
func NewShaper() *Shaper {
	return &Shaper{ShaperStack: PassThrough}
}

// Copy returns a copy of the original object, instead of editing in-place,
// so make sure you've already got a reference to the original
// This should NEVER be hung off of a NewShaper string, or the original NewShaper will be lost
func (shaper *Shaper) Copy() *Shaper {
	return &Shaper{
		ShaperStack: shaper.ShaperStack,
	}
}

// Process will actually process a string using the built-up filter chain
func (shaper *Shaper) Process(s string) string {
	return shaper.ShaperStack(s)
}

// ProcessAny will do the actual processing using the built-up filter chain on the given `interface{}` type input
func (shaper *Shaper) ProcessAny(si interface{}) string {
	if s, found := si.(string); found {
		return shaper.ShaperStack(s)
	}
	if sa, found := si.([]string); found {
		r := ""
		for _, s := range sa {
			r += shaper.ShaperStack(s)
		}
		return r
	}
	return ""
}

// AddShaper is used to apply arbitrary filters
func (shaper *Shaper) AddShaper(f func(string) string) *Shaper {
	shaper.ShaperStack = func(a func(string) string, b func(string) string) func(string) string {
		return func(s string) string {
			return a(b(s))
		}
	}(f, shaper.ShaperStack)
	return shaper
}

// NoOp will do nothing.
// It is used to convert the extended class back to the base class.
func (shaper *Shaper) NoOp() *Shaper {
	return shaper
}

// ApplyTrim will apply/add the strings.TrimSpace filter to the Shaper
// to trim all leading and trailing white spaces
func (shaper *Shaper) ApplyTrim() *Shaper {
	shaper.AddShaper(strings.TrimSpace)
	return shaper
}

// ApplyToLower will apply/add the strings.ToLower filter to the Shaper
func (shaper *Shaper) ApplyToLower() *Shaper {
	shaper.AddShaper(strings.ToLower)
	return shaper
}

// ApplyToUpper will apply/add the strings.ToUpper filter to the Shaper
func (shaper *Shaper) ApplyToUpper() *Shaper {
	shaper.AddShaper(strings.ToUpper)
	return shaper
}

// ApplyTitle will apply/add the strings.Title filter to the Shaper
func (shaper *Shaper) ApplyTitle() *Shaper {
	shaper.AddShaper(strings.Title)
	return shaper
}

// ApplyReplace will apply/add the strings.Replace filter to the Shaper
func (shaper *Shaper) ApplyReplace(old, new string, times int) *Shaper {
	shaper.AddShaper(func(s string) string {
		return strings.Replace(s, old, new, times)
	})
	return shaper
}

// ApplyRegexpReplaceAll will apply/add the regexp.ReplaceAllString filter to the Shaper
func (shaper *Shaper) ApplyRegexpReplaceAll(rexp, repl string) *Shaper {
	shaper.AddShaper(func(s string) string {
		return regexp.MustCompile(rexp).ReplaceAllString(s, repl)
	})
	return shaper
}

// ApplyRegSpaces will apply/add the regular-spaces filter to the Shaper
// to consolidate the string to be single-space delimited
func (shaper *Shaper) ApplyRegSpaces() *Shaper {
	shaper.AddShaper(func(s string) string {
		return regexp.MustCompile(`\s\s+`).ReplaceAllString(s, " ")
	})
	return shaper
}
