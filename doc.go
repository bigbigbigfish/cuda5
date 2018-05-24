/*
	Go bindings for nVIDIA CUDA 5.
	This package compiles with both gc and gccgo.
*/
package gocuda

// Dummy imports so that
// 	go get github.com/bigbigbigfish/gocuda
// will install everything.
import (
	_ "github.com/bigbigbigfish/gocuda/cu"
	_ "github.com/bigbigbigfish/gocuda/cufft"
	_ "github.com/bigbigbigfish/gocuda/safe"
)
