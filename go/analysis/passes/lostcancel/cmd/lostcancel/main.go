// The lostcancel command applies the github.com/bjulian5/tools/go/analysis/passes/lostcancel
// analysis to the specified packages of Go source code.
package main

import (
	"github.com/bjulian5/tools/go/analysis/passes/lostcancel"
	"github.com/bjulian5/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(lostcancel.Analyzer) }
