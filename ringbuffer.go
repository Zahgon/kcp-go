// The MIT License (MIT)
//
// Copyright (c) 2025 xtaci
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

const (
	RINGBUFFER_MIN = 8    // minimum ring buffer capacity
	RINGBUFFER_EXP = 1024 // growth threshold: below this, double; above this, grow by 25%
)

// RingBuffer is a generic ring (circular) buffer that supports dynamic resizing.
// It provides efficient FIFO queue behavior with amortized constant time operations.
type RingBuffer[T any] struct {
	head     int // Index of the next element to be popped
	tail     int // Index of the next empty slot to push into
	elements []T // Underlying slice storing elements in circular fashion
}

// NewRingBuffer creates a new Ring with a specified initial capacity.
// If the provided size is <= 8, it defaults to 8.
func NewRingBuffer[T any](size int) *RingBuffer[T] { _ = "STUB: not implemented"; return nil }

// Ensure a minimum size

// Len returns the number of elements currently in the ring.
func (r *RingBuffer[T]) Len() int { _ = "STUB: not implemented"; return 0 }

// Wrapped case: elements from head to end + elements from start to tail

// Push adds an element to the tail of the ring.
// If the ring is full, it will grow automatically.
func (r *RingBuffer[T]) Push(v T) { _ = "STUB: not implemented"; return }

// Pop removes and returns the element from the head of the ring.
// It returns the zero value and false if the ring is empty.
func (r *RingBuffer[T]) Pop() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// Optional: clear the slot to avoid retaining references

// Peek returns the element at the head of the ring without removing it.
// It returns the zero value and false if the ring is empty.
func (r *RingBuffer[T]) Peek() (*T, bool) { _ = "STUB: not implemented"; return nil, false }

// Discard discards the first N elements from the ring buffer.
// Returns the number of elements that are actually discarded (<= n).
func (r *RingBuffer[T]) Discard(n int) int { _ = "STUB: not implemented"; return 0 }

// no wrap: clear contiguous range

// wraps around

// ForEach iterates over each element in the ring buffer,
// applying the provided function. If the function returns false,
// iteration stops early.
func (r *RingBuffer[T]) ForEach(fn func(*T) bool) { _ = "STUB: not implemented"; return }

// Contiguous data: [head ... tail)

// Wrapped data: [head ... end) + [0 ... tail)

// ForEachReverse iterates over each element in the ring buffer in reverse order,
// applying the provided function. If the function returns false,
// iteration stops early.
func (r *RingBuffer[T]) ForEachReverse(fn func(*T) bool) { _ = "STUB: not implemented"; return }

// Contiguous data: [head ... tail)

// Clear resets the ring to an empty state and reinitializes the buffer.
func (r *RingBuffer[T]) Clear() {
	_ = "STUB: not implemented"

	// Only clear elements that contain data to avoid retaining references
	return
}

// IsEmpty returns true if the ring has no elements.
func (r *RingBuffer[T]) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// MaxLen returns the maximum capacity of the ring buffer.
func (r *RingBuffer[T]) MaxLen() int { _ = "STUB: not implemented"; return 0 }

// IsFull returns true if the ring buffer is full (tail + 1 == head).
func (r *RingBuffer[T]) IsFull() bool { _ = "STUB: not implemented"; return false }

// grow increases the ring buffer's capacity when full.
// Growth policy:
//   - If current size < RINGBUFFER_MIN : grow to RINGBUFFER_MIN
//   - If size < RINGBUFFER_EXP: double the size
//   - If size > RINGBUFFER_EXP: increase by 10% (rounded up)
func (r *RingBuffer[T]) grow() { _ = "STUB: not implemented"; return }

// +10%, rounded up

// Copy elements to new buffer preserving logical order

// Contiguous data: [head ... tail)

// Wrapped data: [head ... end) + [0 ... tail)
