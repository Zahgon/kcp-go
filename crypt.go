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
	"sync"
)

var (
	// a defined initial vector
	// https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Initialization_vector_.28IV.29
	// https://en.wikipedia.org/wiki/Initialization_vector
	// actually initial vector is not used in this package, we prepend a random nonce to each outgoing packets.
	// though IV is fixed, the first 8 bytes of the encrypted data is always random.
	initialVector = []byte{167, 115, 79, 156, 18, 172, 27, 1, 164, 21, 242, 193, 252, 120, 230, 107}
	saltxor       = `sH3CIVoF#rWLtJo6`
)

// BlockCrypt defines encryption/decryption methods for a given byte slice.
// Notes on implementing: the data to be encrypted contains a builtin
// nonce at the first 16 bytes
type BlockCrypt interface {
	// Encrypt encrypts the whole block in src into dst.
	// Dst and src may point at the same memory.
	Encrypt(dst, src []byte)

	// Decrypt decrypts the whole block in src into dst.
	// Dst and src may point at the same memory.
	Decrypt(dst, src []byte)
}

var _ BlockCrypt = &aeadCrypt{}

// aeadCrypt implements BlockCrypt interface using cipher.AEAD
type aeadCrypt struct {
	aead cipher.AEAD
}

func (aeadCrypt) Encrypt(_, _ []byte) { _ = "STUB: not implemented"; return }

func (aeadCrypt) Decrypt(_, _ []byte) { _ = "STUB: not implemented"; return }

func (a *aeadCrypt) Seal(dst, nonce, plaintext, additionalData []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (a *aeadCrypt) Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *aeadCrypt) NonceSize() int { _ = "STUB: not implemented"; return 0 }

func (a *aeadCrypt) Overhead() int { _ = "STUB: not implemented"; return 0 }

// NewAEADCrypt creates an AEAD BlockCrypt instance from an existing cipher.AEAD
func NewAEADCrypt(aead cipher.AEAD) BlockCrypt { _ = "STUB: not implemented"; return *new(BlockCrypt) }

// NewAESGCMCrypt creates an AEAD BlockCrypt instance using AES-GCM
// key must be either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256.
func NewAESGCMCrypt(key []byte) (BlockCrypt, error) {
	_ = "STUB: not implemented"
	return *new(BlockCrypt), nil
}

var _ BlockCrypt = &blockCrypt{}

// blockCrypt implements BlockCrypt interface using a cipher.Block
type blockCrypt struct {
	encMu     sync.Mutex
	decMu     sync.Mutex
	encbuf    []byte // encryption working buffer
	decbuf    []byte // decryption working buffer
	block     cipher.Block
	blockSize int // cached block size
}

//go:nosplit
func (c *blockCrypt) Encrypt(dst, src []byte) { _ = "STUB: not implemented"; return }

//go:nosplit
func (c *blockCrypt) Decrypt(dst, src []byte) { _ = "STUB: not implemented"; return }

func newBlockCrypt(block cipher.Block) BlockCrypt {
	_ = "STUB: not implemented"
	return *new(BlockCrypt)
}

type salsa20BlockCrypt struct {
	key [32]byte
}

// NewSalsa20BlockCrypt https://en.wikipedia.org/wiki/Salsa20
func NewSalsa20BlockCrypt(key []byte) (BlockCrypt, error) {
	_ = "STUB: not implemented"
	return *new(BlockCrypt), nil
}

//go:nosplit
func (c *salsa20BlockCrypt) Encrypt(dst, src []byte) { _ = "STUB: not implemented"; return }

//go:nosplit
func (c *salsa20BlockCrypt) Decrypt(dst, src []byte) { _ = "STUB: not implemented"; return }

// NewSM4BlockCrypt https://github.com/tjfoc/gmsm/tree/master/sm4
func NewSM4BlockCrypt(key []byte) (BlockCrypt, error) {
	_ = "STUB: not implemented"
	return *new(BlockCrypt), nil
}

// NewTwofishBlockCrypt https://en.wikipedia.org/wiki/Twofish
func NewTwofishBlockCrypt(key []byte) (BlockCrypt, error) {
	_ = "STUB: not implemented"
	return *new(BlockCrypt), nil
}

// NewTripleDESBlockCrypt https://en.wikipedia.org/wiki/Triple_DES
func NewTripleDESBlockCrypt(key []byte) (BlockCrypt, error) {
	_ = "STUB: not implemented"
	return *new(BlockCrypt), nil
}

// NewCast5BlockCrypt https://en.wikipedia.org/wiki/CAST-128
func NewCast5BlockCrypt(key []byte) (BlockCrypt, error) {
	_ = "STUB: not implemented"
	return *new(BlockCrypt), nil
}

// NewBlowfishBlockCrypt https://en.wikipedia.org/wiki/Blowfish_(cipher)
func NewBlowfishBlockCrypt(key []byte) (BlockCrypt, error) {
	_ = "STUB: not implemented"
	return *new(BlockCrypt), nil
}

// NewAESBlockCrypt https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
func NewAESBlockCrypt(key []byte) (BlockCrypt, error) {
	_ = "STUB: not implemented"
	return *new(BlockCrypt), nil
}

// NewTEABlockCrypt https://en.wikipedia.org/wiki/Tiny_Encryption_Algorithm
func NewTEABlockCrypt(key []byte) (BlockCrypt, error) {
	_ = "STUB: not implemented"
	return *new(BlockCrypt), nil
}

// NewXTEABlockCrypt https://en.wikipedia.org/wiki/XTEA
func NewXTEABlockCrypt(key []byte) (BlockCrypt, error) {
	_ = "STUB: not implemented"
	return *new(BlockCrypt), nil
}

type simpleXORBlockCrypt struct {
	xortbl []byte
}

// NewSimpleXORBlockCrypt simple xor with key expanding
func NewSimpleXORBlockCrypt(key []byte) (BlockCrypt, error) {
	_ = "STUB: not implemented"
	return *new(BlockCrypt), nil
}

func (c *simpleXORBlockCrypt) Encrypt(dst, src []byte) { _ = "STUB: not implemented"; return }

func (c *simpleXORBlockCrypt) Decrypt(dst, src []byte) { _ = "STUB: not implemented"; return }

type noneBlockCrypt struct{}

// NewNoneBlockCrypt does nothing but copying
func NewNoneBlockCrypt(key []byte) (BlockCrypt, error) {
	_ = "STUB: not implemented"
	return *new(BlockCrypt), nil
}

//go:nosplit
func (c *noneBlockCrypt) Encrypt(dst, src []byte) { _ = "STUB: not implemented"; return }

//go:nosplit
func (c *noneBlockCrypt) Decrypt(dst, src []byte) { _ = "STUB: not implemented"; return }

// -----------------------------------------------------------------------
// CFB-mode encryption/decryption
// -----------------------------------------------------------------------
//
// For block ciphers (non-AEAD), packets are encrypted using a local variant
// of CFB (Cipher Feedback) mode. The encrypt/decrypt functions below are
// hand-optimized with 8x loop unrolling to reduce data dependencies
// between consecutive block cipher calls.
//
// Two block sizes are supported:
//   - 8-byte blocks  (e.g. Blowfish, CAST5, DES, TEA, XTEA)
//   - 16-byte blocks (e.g. AES, Twofish, SM4)

// encrypt dispatches to the appropriate block-size-specific CFB encryptor.
func encrypt(block cipher.Block, dst, src, buf []byte) { _ = "STUB: not implemented"; return }

// encrypt8 performs CFB encryption for 8-byte block ciphers.
// Uses 8x loop unrolling for throughput optimization.
func encrypt8(block cipher.Block, dst, src, buf []byte) { _ = "STUB: not implemented"; return }

// number of full 8-byte blocks

// number of 8-block groups (64 bytes each)
// remaining blocks after groups

// 1

// 2

// 3

// 4

// 5

// 6

// 7

// 8

// encrypt16 performs CFB encryption for 16-byte block ciphers.
// Uses 8x loop unrolling for throughput optimization.
func encrypt16(block cipher.Block, dst, src, buf []byte) { _ = "STUB: not implemented"; return }

// number of full 16-byte blocks

// number of 8-block groups (128 bytes each)
// remaining blocks after groups

// 1

// 2

// 3

// 4

// 5

// 6

// 7

// 8

// decrypt dispatches to the appropriate block-size-specific CFB decryptor.
func decrypt(block cipher.Block, dst, src, buf []byte) { _ = "STUB: not implemented"; return }

// decrypt8 performs CFB decryption for 8-byte block ciphers.
// Uses double-buffering (tbl/next) with 8x loop unrolling
// to break the data dependency chain between consecutive blocks.
func decrypt8(block cipher.Block, dst, src, buf []byte) { _ = "STUB: not implemented"; return }

// number of full 8-byte blocks

// number of 8-block groups (64 bytes each)
// remaining blocks after groups

// 8x loop unrolling: alternates tbl/next to relieve data dependency

// 1

// 2

// 3

// 4

// 5

// 6

// 7

// 8

// decrypt16 performs CFB decryption for 16-byte block ciphers.
// Uses double-buffering (tbl/next) with 8x loop unrolling.
func decrypt16(block cipher.Block, dst, src, buf []byte) { _ = "STUB: not implemented"; return }

// number of full 16-byte blocks

// number of 8-block groups (128 bytes each)
// remaining blocks after groups

// 8x loop unrolling: alternates tbl/next to relieve data dependency

// 1

// 2

// 3

// 4

// 5

// 6

// 7

// 8
