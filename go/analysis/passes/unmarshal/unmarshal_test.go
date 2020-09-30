// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unmarshal_test

import (
	"testing"

	"github.com/bjulian5/tools/go/analysis/analysistest"
	"github.com/bjulian5/tools/go/analysis/passes/unmarshal"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, unmarshal.Analyzer, "a")
}
