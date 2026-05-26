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

// Snmp defines network statistics indicator
type Snmp struct {
	BytesSent           uint64 // bytes sent from upper level
	BytesReceived       uint64 // bytes received to upper level
	MaxConn             uint64 // max number of connections ever reached
	ActiveOpens         uint64 // accumulated active open connections
	PassiveOpens        uint64 // accumulated passive open connections
	CurrEstab           uint64 // current number of established connections
	InErrs              uint64 // UDP read errors reported from net.PacketConn
	InCsumErrors        uint64 // checksum errors from CRC32
	KCPInErrors         uint64 // packet input errors reported from KCP
	InPkts              uint64 // incoming packets count
	OutPkts             uint64 // outgoing packets count
	InSegs              uint64 // incoming KCP segments
	OutSegs             uint64 // outgoing KCP segments
	InBytes             uint64 // UDP bytes received
	OutBytes            uint64 // UDP bytes sent
	RetransSegs         uint64 // accumulated retransmitted segments
	FastRetransSegs     uint64 // accumulated fast retransmitted segments
	EarlyRetransSegs    uint64 // accumulated early retransmitted segments
	LostSegs            uint64 // number of segs inferred as lost
	RepeatSegs          uint64 // number of segs duplicated
	FECFullShardSet     uint64 // number of FEC segments that are full
	FECRecovered        uint64 // correct packets recovered from FEC
	FECErrs             uint64 // incorrect packets recovered from FEC
	FECParityShards     uint64 // FEC segments received
	FECShardSet         uint64 // number of shard sets that are not yet complete
	FECShardMin         uint64 // minimum shard ID among active FEC shard sets
	RingBufferSndQueue  uint64 // Len of segments in send queue ring buffer
	RingBufferRcvQueue  uint64 // Len of segments in receive queue ring buffer
	RingBufferSndBuffer uint64 // Len of segments in send buffer ring buffer
	OOBPackets          uint64 // number of OOB packets received
}

func newSnmp() *Snmp {
	_ = "STUB: not implemented"

	// Header returns all field names
	return nil
}

func (s *Snmp) Header() []string { _ = "STUB: not implemented"; return nil }

// ToSlice returns current snmp info as slice
func (s *Snmp) ToSlice() []string { _ = "STUB: not implemented"; return nil }

// Copy make a copy of current snmp snapshot
func (s *Snmp) Copy() *Snmp { _ = "STUB: not implemented"; return nil }

// Reset values to zero
func (s *Snmp) Reset() { _ = "STUB: not implemented"; return }

// DefaultSnmp is the global KCP connection statistics collector
var DefaultSnmp *Snmp

func init() {
	DefaultSnmp = newSnmp()
}
