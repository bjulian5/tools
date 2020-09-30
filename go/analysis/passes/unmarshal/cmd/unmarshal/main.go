// The unmarshal command runs the unmarshal analyzer.
package main

import (
	"github.com/bjulian5/tools/go/analysis/passes/unmarshal"
	"github.com/bjulian5/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(unmarshal.Analyzer) }
