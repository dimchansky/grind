// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/dimchansky/grind/grinder"
	"github.com/dimchansky/grind/grindtest"
)

var builtins = []grinder.Func{
	DeleteUnusedLabels,
}

func TestGrind(t *testing.T) {
	grindtest.TestGlob(t, "testdata/grind-*.go", builtins)
}
