// +build ignore

// This file provides an example command for static checkers
// conforming to the github.com/bjulian5/tools/go/analysis API.
// It serves as a model for the behavior of the cmd/vet tool in $GOROOT.
// Being based on the unitchecker driver, it must be run by go vet:
//
//   $ go build -o unitchecker main.go
//   $ go vet -vettool=unitchecker my/project/...
//
// For a checker also capable of running standalone, use multichecker.
package main

import (
	"github.com/bjulian5/tools/go/analysis/unitchecker"

	"github.com/bjulian5/tools/go/analysis/passes/asmdecl"
	"github.com/bjulian5/tools/go/analysis/passes/assign"
	"github.com/bjulian5/tools/go/analysis/passes/atomic"
	"github.com/bjulian5/tools/go/analysis/passes/bools"
	"github.com/bjulian5/tools/go/analysis/passes/buildtag"
	"github.com/bjulian5/tools/go/analysis/passes/cgocall"
	"github.com/bjulian5/tools/go/analysis/passes/composite"
	"github.com/bjulian5/tools/go/analysis/passes/copylock"
	"github.com/bjulian5/tools/go/analysis/passes/errorsas"
	"github.com/bjulian5/tools/go/analysis/passes/httpresponse"
	"github.com/bjulian5/tools/go/analysis/passes/loopclosure"
	"github.com/bjulian5/tools/go/analysis/passes/lostcancel"
	"github.com/bjulian5/tools/go/analysis/passes/nilfunc"
	"github.com/bjulian5/tools/go/analysis/passes/printf"
	"github.com/bjulian5/tools/go/analysis/passes/shift"
	"github.com/bjulian5/tools/go/analysis/passes/stdmethods"
	"github.com/bjulian5/tools/go/analysis/passes/structtag"
	"github.com/bjulian5/tools/go/analysis/passes/tests"
	"github.com/bjulian5/tools/go/analysis/passes/unmarshal"
	"github.com/bjulian5/tools/go/analysis/passes/unreachable"
	"github.com/bjulian5/tools/go/analysis/passes/unsafeptr"
	"github.com/bjulian5/tools/go/analysis/passes/unusedresult"
)

func main() {
	unitchecker.Main(
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		bools.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		errorsas.Analyzer,
		httpresponse.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		nilfunc.Analyzer,
		printf.Analyzer,
		shift.Analyzer,
		stdmethods.Analyzer,
		structtag.Analyzer,
		tests.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		unsafeptr.Analyzer,
		unusedresult.Analyzer,
	)
}
