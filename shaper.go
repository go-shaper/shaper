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

// PassThrough will return the string as-is, used primarily for NewFilter()
func PassThrough(s string) string {
	return s
}

// NewFilter makes a new Shaper filter
func NewFilter() *Shaper {
	return &Shaper{ShaperStack: PassThrough}
}

// Copy returns a copy of the original object, instead of editing in-place,
// so make sure you've already got a reference to the original
// This should NEVER be hung off of a NewFilter string, or the original NewFilter will be lost
func (shaper *Shaper) Copy() *Shaper {
	return &Shaper{
		ShaperStack: shaper.ShaperStack,
	}
}

// Process will actually process a string using the built-up filter chain
func (shaper *Shaper) Process(s string) string {
	return shaper.ShaperStack(s)
}

// AddFilter is used to apply arbitrary filters
func (shaper *Shaper) AddFilter(f func(string) string) *Shaper {
	shaper.ShaperStack = func(a func(string) string, b func(string) string) func(string) string {
		return func(s string) string {
			return a(b(s))
		}
	}(f, shaper.ShaperStack)
	return shaper
}

// ApplyTrim will apply/add the strings.TrimSpace filter to the Shaper
// to trim all leading and trailing white spaces
func (shaper *Shaper) ApplyTrim() *Shaper {
	shaper.AddFilter(strings.TrimSpace)
	return shaper
}

// ApplyToLower will apply/add the strings.ToLower filter to the Shaper
func (shaper *Shaper) ApplyToLower() *Shaper {
	shaper.AddFilter(strings.ToLower)
	return shaper
}

// ApplyToUpper will apply/add the strings.ToUpper filter to the Shaper
func (shaper *Shaper) ApplyToUpper() *Shaper {
	shaper.AddFilter(strings.ToUpper)
	return shaper
}

// ApplyReplace will apply/add the strings.Replace filter to the Shaper
func (shaper *Shaper) ApplyReplace(old, new string, times int) *Shaper {
	shaper.AddFilter(func(s string) string {
		return strings.Replace(s, old, new, times)
	})
	return shaper
}

// ApplyRegexpReplaceAll will apply/add the regexp.ReplaceAllString filter to the Shaper
func (shaper *Shaper) ApplyRegexpReplaceAll(rexp, repl string) *Shaper {
	shaper.AddFilter(func(s string) string {
		return regexp.MustCompile(rexp).ReplaceAllString(s, repl)
	})
	return shaper
}
