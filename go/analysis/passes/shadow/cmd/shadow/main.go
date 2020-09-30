// The shadow command runs the shadow analyzer.
package main

import (
	"github.com/bjulian5/tools/go/analysis/passes/shadow"
	"github.com/bjulian5/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(shadow.Analyzer) }
