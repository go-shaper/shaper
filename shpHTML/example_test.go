package shpHTML_test

import (
	"fmt"

	"github.com/go-shaper/shaper/shpHTML"
)

////////////////////////////////////////////////////////////////////////////

// for standalone test, change package to `main` and the next func def to,
// func main() {
func Example() {
	var hu *shpHTML.ShpHTML

	hu = shpHTML.NewShaper()
	hu.ApplyToUpper()
	fmt.Printf("%s\n", hu.Process("2 &gt;= 1"))
	hu.ApplyHTMLUnescape()
	fmt.Printf("%s\n", hu.Process("2 &gt;= 1"))
	fmt.Println("=============================")

	hu = shpHTML.NewShaper()
	hu.ApplyTagRemove()
	fmt.Printf("%s\n", hu.Process(htmlStr))
	fmt.Println("=============================")

	hu.ApplyTrim().ApplyRegSpaces()
	fmt.Printf("%s\n", hu.Process(htmlStr))

	// Output:
	// 2 &GT;= 1
	// 2 >= 1
	// =============================
	//
	// H1
	//
	// H2
	//
	// H3
	// Emphasis, aka italics, with asterisks or underscores.
	// Strong emphasis, aka bold, with asterisks or underscores.
	// Combined emphasis with asterisks and underscores.
	// Strikethrough uses two tildes. Scratch this.
	//
	//   I'm an inline-style link
	//
	// I'm an  inline-style  link with title
	// I'm a reference-style link
	// I'm a relative reference to a repository file
	// You can use numbers for reference-style link definitions
	// Or leave it empty and use the link text itself.
	//
	// =============================
	// H1 H2 H3
	// Emphasis, aka italics, with asterisks or underscores.
	// Strong emphasis, aka bold, with asterisks or underscores.
	// Combined emphasis with asterisks and underscores.
	// Strikethrough uses two tildes. Scratch this. I'm an inline-style link I'm an inline-style link with title
	// I'm a reference-style link
	// I'm a relative reference to a repository file
	// You can use numbers for reference-style link definitions
	// Or leave it empty and use the link text itself.
}

// To show the full code in GoDoc
var htmlStr string = `<h1>
<a id="user-content-h1" class="anchor" href="#h1" aria-hidden="true"><svg aria-hidden="true" class="octicon octicon-link" height="16" version="1.1" viewbox="0 0 16 16" width="16"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>H1</h1>
<h2>
<a id="user-content-h2" class="anchor" href="#h2" aria-hidden="true"><svg aria-hidden="true" class="octicon octicon-link" height="16" version="1.1" viewbox="0 0 16 16" width="16"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>H2</h2>
<h3>
<a id="user-content-h3" class="anchor" href="#h3" aria-hidden="true"><svg aria-hidden="true" class="octicon octicon-link" height="16" version="1.1" viewbox="0 0 16 16" width="16"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>H3</h3>
<p>Emphasis, aka italics, with <em>asterisks</em> or <em>underscores</em>.</p>
<p>Strong emphasis, aka bold, with <strong>asterisks</strong> or <strong>underscores</strong>.</p>
<p>Combined emphasis with <strong>asterisks and <em>underscores</em></strong>.</p>
<p>Strikethrough uses two tildes. <del>Scratch this.</del></p>
<p><a href="https://www.google.com">
  I'm an inline-style link
</a></p>
<p><a href="https://www.google.com" title="Google's Homepage">I'm an  inline-style  link with title</a></p>
<p><a href="https://www.mozilla.org">I'm a reference-style link</a></p>
<p><a href="../blob/master/LICENSE">I'm a relative reference to a repository file</a></p>
<p><a href="http://slashdot.org">You can use numbers for reference-style link definitions</a></p>
<p>Or leave it empty and use the <a href="http://www.reddit.com">link text itself</a>.</p>
`
