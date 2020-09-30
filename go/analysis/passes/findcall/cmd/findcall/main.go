// The findcall command runs the findcall analyzer.
package main

import (
	"github.com/bjulian5/tools/go/analysis/passes/findcall"
	"github.com/bjulian5/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(findcall.Analyzer) }
