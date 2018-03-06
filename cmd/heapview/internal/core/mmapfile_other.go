// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !darwin,!linux

package core

// TODO(matloob): perhaps use the more portable github.com/golang/exp/mmap
// instead of the mmap code in mmapfile.go.

type mmapFile struct{}

func (m *mmapFile) Close() error { return nil }
