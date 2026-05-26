// The MIT License (MIT)
//
// Copyright (c) 2015 xtaci
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package kcp

import (
	"crypto/cipher"
	"io"
	"math/rand/v2"
	"runtime"
	"sync"

	"golang.org/x/sys/cpu"
)

const reseedInterval = 1 << 24 // reseed after ~16M reads to limit key exposure

var (
	hasAESAsmAMD64 = cpu.X86.HasAES && cpu.X86.HasSSE41 && cpu.X86.HasSSSE3
	hasAESAsmARM64 = cpu.ARM64.HasAES
	hasAESAsmS390X = cpu.S390X.HasAES
	hasAESAsmPPC64 = runtime.GOARCH == "ppc64" || runtime.GOARCH == "ppc64le"

	hasAESHardwareSupport = hasAESAsmAMD64 || hasAESAsmARM64 || hasAESAsmS390X || hasAESAsmPPC64

	entropy io.Reader = NewEntropy()
)

// NewEntropy creates a new entropy source.
func NewEntropy() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// SetEntropy sets the global entropy source used by fillRand.
func SetEntropy(r io.Reader) {
	_ = "STUB: not implemented"

	// fillRand fills p with random data from the global entropy source.
	return
}

func fillRand(p []byte) { _ = "STUB: not implemented"; return }

// rngAES is an AES-based random number generator.
type rngAES struct {
	mutex sync.Mutex
	block cipher.Block
	seed  [16]byte
	count uint64
}

// NewEntropyAES creates a new AES-based entropy source.
func NewEntropyAES() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// updateSeed updates the AES seed after a certain number of reads.
func (r *rngAES) updateSeed() { _ = "STUB: not implemented"; return }

// Read fills p with random data using AES encryption.
func (r *rngAES) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// rngChacha8 is a ChaCha8-based random number generator.
type rngChacha8 struct {
	mutex sync.Mutex
	rand  *rand.ChaCha8
	count uint64
}

// NewEntropyChacha8 creates a new ChaCha8-based entropy source.
func NewEntropyChacha8() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// updateSeed updates the ChaCha8 seed after a certain number of reads.
func (r *rngChacha8) updateSeed() { _ = "STUB: not implemented"; return }

// Read fills p with random data using ChaCha8.
func (r *rngChacha8) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
