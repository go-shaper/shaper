
# shaper - shape strings into desired forms

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Doc](https://godoc.org/github.com/go-shaper/shaper?status.svg)](https://godoc.org/github.com/go-shaper/shaper)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-shaper/shaper)](https://goreportcard.com/report/github.com/go-shaper/shaper)
[ ![Codeship Status for go-shaper/shaper](https://codeship.com/projects/04245480-e7ff-0133-dc96-46bb3aa6b241/status?branch=master)](https://codeship.com/projects/147070)

The `shaper` project provides general purpose building blocks to shape strings into desired forms. The cumulative building blocks can easily be chained together. They are used like building pipes in shell. 

Architected by Howard C. Shaw III, it needs neither go-routines nor channels. Instead, it creates a composable set of filters you could keep around and apply to a single string at a time. Behind the scenes, it is still building a composition of functions, but by currying that composition into a new function and holding the state of the stack in a struct, it restores the left-right ordering of the filters. Note that the filter stages are basically just compile-time freezes of a call to the currying function; i.e., no matter how complicated the shaping filter chain is, all are done at the compile time, so no run-time overhead when using `shaper`. 

Check out the [provided examples](https://godoc.org/github.com/go-shaper/shaper#example-package--Output)  to see how to use it, and the [![Go Doc](https://godoc.org/github.com/go-shaper/shaper?status.svg)](https://godoc.org/github.com/go-shaper/shaper) document for further details.


All patches welcome. 
