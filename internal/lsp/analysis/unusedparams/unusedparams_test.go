// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unusedparams_test

import (
	"testing"

	"github.com/bjulian5/tools/go/analysis/analysistest"
	"github.com/bjulian5/tools/internal/lsp/analysis/unusedparams"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, unusedparams.Analyzer, "a")
}
