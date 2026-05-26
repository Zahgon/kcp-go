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

//go:build linux

package kcp

const (
	batchSize = 256 // max packets per recvmmsg/sendmmsg call
)

// readLoop is the optimized read loop for Linux, utilizing the recvmmsg syscall
// to batch-receive multiple UDP packets in a single system call.
func (s *UDPSession) readLoop() {
	_ = "STUB: not implemented"
	// default version
	return
}

// x/net version

// make sure the packet is from the same source
// set source address if nil

// source and size has validated

// monitor is the optimized version of monitor for linux utilizing recvmmsg syscall
func (l *Listener) monitor() { _ = "STUB: not implemented"; return }

// default version

// x/net version
