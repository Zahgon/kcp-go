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
	"time"
)

// KCP Protocol Constants
const (
	// Retransmission Timeout (RTO) bounds, in milliseconds
	IKCP_RTO_NDL = 30    // no-delay mode: minimum RTO (ms)
	IKCP_RTO_MIN = 100   // normal mode: minimum RTO (ms)
	IKCP_RTO_DEF = 200   // default RTO (ms)
	IKCP_RTO_MAX = 60000 // maximum RTO (ms), 60 seconds

	// Command types for the KCP segment header (cmd field)
	IKCP_CMD_PUSH = 81 // cmd: push data
	IKCP_CMD_ACK  = 82 // cmd: acknowledge
	IKCP_CMD_WASK = 83 // cmd: window probe request (ask)
	IKCP_CMD_WINS = 84 // cmd: window size response (tell)

	// Probe flags (bitfield), set in kcp.probe to schedule probe commands
	IKCP_ASK_SEND = 1 // schedule sending IKCP_CMD_WASK
	IKCP_ASK_TELL = 2 // schedule sending IKCP_CMD_WINS

	// Default window and MTU sizes
	IKCP_WND_SND = 32   // default send window size (packets)
	IKCP_WND_RCV = 32   // default receive window size (packets)
	IKCP_MTU_DEF = 1400 // default MTU (bytes, not including UDP/IP header)

	// Protocol parameters
	IKCP_ACK_FAST    = 3      // fast retransmit trigger threshold (duplicate ACK count)
	IKCP_INTERVAL    = 100    // default flush interval (ms)
	IKCP_OVERHEAD    = 24     // per-segment header size: conv(4) + cmd(1) + frg(1) + wnd(2) + ts(4) + sn(4) + una(4) + len(4)
	IKCP_DEADLINK    = 20     // max retransmissions before declaring dead link
	IKCP_THRESH_INIT = 2      // initial slow-start threshold (packets)
	IKCP_THRESH_MIN  = 2      // minimum slow-start threshold (packets)
	IKCP_PROBE_INIT  = 500    // initial window probe timeout (ms)
	IKCP_PROBE_LIMIT = 120000 // maximum window probe timeout (ms), 120 seconds
	IKCP_SN_OFFSET   = 12     // byte offset of sequence number (sn) within the segment header
)

type PacketType int8

const (
	IKCP_PACKET_REGULAR PacketType = iota
	IKCP_PACKET_FEC
)

type FlushType int8

const (
	IKCP_FLUSH_ACKONLY FlushType = 1 << iota
	IKCP_FLUSH_FULL
)

type KCPLogType int32

const (
	IKCP_LOG_OUTPUT KCPLogType = 1 << iota
	IKCP_LOG_INPUT
	IKCP_LOG_SEND
	IKCP_LOG_RECV
	IKCP_LOG_OUT_ACK
	IKCP_LOG_OUT_PUSH
	IKCP_LOG_OUT_WASK
	IKCP_LOG_OUT_WINS
	IKCP_LOG_IN_ACK
	IKCP_LOG_IN_PUSH
	IKCP_LOG_IN_WASK
	IKCP_LOG_IN_WINS
)

const (
	IKCP_LOG_OUTPUT_ALL = IKCP_LOG_OUTPUT | IKCP_LOG_OUT_ACK | IKCP_LOG_OUT_PUSH | IKCP_LOG_OUT_WASK | IKCP_LOG_OUT_WINS
	IKCP_LOG_INPUT_ALL  = IKCP_LOG_INPUT | IKCP_LOG_IN_ACK | IKCP_LOG_IN_PUSH | IKCP_LOG_IN_WASK | IKCP_LOG_IN_WINS
	IKCP_LOG_ALL        = IKCP_LOG_OUTPUT_ALL | IKCP_LOG_INPUT_ALL | IKCP_LOG_SEND | IKCP_LOG_RECV
)

// monotonic reference time point
var refTime time.Time = time.Now()

// currentMs returns current elapsed monotonic milliseconds since program startup
func currentMs() uint32 { _ = "STUB: not implemented"; return 0 }

// output_callback is a prototype which ought capture conn and call conn.Write
type output_callback func(buf []byte, size int)

// logoutput_callback is a prototype which logging kcp trace output
type logoutput_callback func(msg string, args ...any)

func _itimediff(later, earlier uint32) int32 { _ = "STUB: not implemented"; return 0 }

// segment defines a KCP segment
type segment struct {
	conv     uint32
	cmd      uint8
	frg      uint8
	wnd      uint16
	ts       uint32
	sn       uint32
	una      uint32
	rto      uint32
	xmit     uint32
	resendts uint32
	fastack  uint32
	acked    uint32 // mark if the seg has acked
	data     []byte
}

// encode a segment header into buffer
func (seg *segment) encode(ptr []byte) []byte { _ = "STUB: not implemented"; return nil }

// BCE hint

// segmentHeap is a min-heap of segments, used for receiving segments in order
type segmentHeap struct {
	segments []segment
	marks    map[uint32]struct{} // to avoid duplicates
}

func newSegmentHeap() *segmentHeap { _ = "STUB: not implemented"; return nil }

func (h *segmentHeap) Len() int { _ = "STUB: not implemented"; return 0 }

func (h *segmentHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (h *segmentHeap) Swap(i, j int) { _ = "STUB: not implemented"; return }
func (h *segmentHeap) Push(x any)    { _ = "STUB: not implemented"; return }

func (h *segmentHeap) Pop() any { _ = "STUB: not implemented"; return *new(any) }

// clear reference to avoid memory leak

func (h *segmentHeap) Has(sn uint32) bool { _ = "STUB: not implemented"; return false }

// KCP defines a single KCP connection's protocol state machine.
// It is a pure ARQ (Automatic Repeat reQuest) implementation with no I/O.
type KCP struct {
	// Connection identity and framing
	conv  uint32 // conversation id, must be equal on both sides
	mtu   uint32 // maximum transmission unit (bytes)
	mss   uint32 // maximum segment size = mtu - IKCP_OVERHEAD
	state uint32 // connection state, 0 = active, 0xFFFFFFFF = dead link

	// Sequence numbers and acknowledgment tracking
	snd_una uint32 // oldest unacknowledged sequence number
	snd_nxt uint32 // next sequence number to send
	rcv_nxt uint32 // next expected sequence number to receive

	// Congestion control (RFC 5681 / RFC 6937)
	ssthresh           uint32 // slow-start threshold (packets)
	rx_rttvar, rx_srtt int32  // RTT variance and smoothed RTT (ms), per RFC 6298
	rx_rto, rx_minrto  uint32 // retransmission timeout and its lower bound (ms)
	snd_wnd            uint32 // local send window size (packets)
	rcv_wnd            uint32 // local receive window size (packets)
	rmt_wnd            uint32 // remote advertised window size (packets)
	cwnd               uint32 // congestion window (packets)
	incr               uint32 // bytes accumulated for cwnd increment

	// Window probing
	probe      uint32 // probe flags (IKCP_ASK_SEND / IKCP_ASK_TELL)
	ts_probe   uint32 // timestamp for next window probe (ms)
	probe_wait uint32 // current probe timeout (ms), doubles on each retry

	// Timers and scheduling
	interval uint32 // flush interval (ms)
	ts_flush uint32 // next flush timestamp (ms)
	nodelay  uint32 // 0: normal, 1: no-delay mode (reduces RTO aggressively)
	updated  uint32 // whether Update() has been called at least once

	// Reliability
	dead_link  uint32 // max retransmit count before link is considered dead
	fastresend int32  // fast retransmit trigger count, 0 = disabled
	nocwnd     int32  // 1 = disable congestion control
	stream     int32  // 1 = stream mode (no message boundaries), 0 = message mode

	// Logging
	logmask KCPLogType

	// Data queues and buffers
	snd_queue *RingBuffer[segment] // send queue: segments waiting to enter the send window
	rcv_queue *RingBuffer[segment] // receive queue: ordered segments ready for user read
	snd_buf   *RingBuffer[segment] // send buffer: segments in-flight (sent but unacknowledged)
	rcv_buf   *segmentHeap         // receive buffer: out-of-order segments awaiting reordering

	acklist []ackItem // pending ACKs to be flushed

	buffer []byte          // pre-allocated encoding buffer for flush()
	output output_callback // callback to write data to the underlying transport

	log logoutput_callback // trace log callback
}

type ackItem struct {
	sn uint32
	ts uint32
}

// NewKCP create a new kcp state machine
//
// 'conv' must be equal in the connection peers, or else data will be silently rejected.
//
// 'output' function will be called whenever these is data to be sent on wire.
func NewKCP(conv uint32, output output_callback) *KCP { _ = "STUB: not implemented"; return nil }

// newSegment creates a KCP segment
func (kcp *KCP) newSegment(size int) (seg segment) { _ = "STUB: not implemented"; return *new(segment) }

// recycleSegment recycles a KCP segment
func (kcp *KCP) recycleSegment(seg *segment) { _ = "STUB: not implemented"; return }

// PeekSize checks the size of next message in the recv queue
func (kcp *KCP) PeekSize() (length int) { _ = "STUB: not implemented"; return 0 }

// Receive data from kcp state machine
//
// Return number of bytes read.
//
// Return -1 when there is no readable data.
//
// Return -2 if len(buffer) is smaller than kcp.PeekSize().
func (kcp *KCP) Recv(buffer []byte) (n int) { _ = "STUB: not implemented"; return 0 }

// merge fragment

// move available data from rcv_buf -> rcv_queue

// push back segment

// fast recover

// ready to send back IKCP_CMD_WINS in ikcp_flush
// tell remote my window size

// Send is user/upper level send, returns below zero for error
func (kcp *KCP) Send(buffer []byte) int { _ = "STUB: not implemented"; return 0 }

// append to previous segment in streaming mode (if possible)

// grow slice, the underlying cap is guaranteed to
// be larger than kcp.mss

// message mode

// stream mode

// update_ack updates the smoothed RTT and RTO based on a new RTT sample.
// Algorithm follows RFC 6298: Computing TCP's Retransmission Timer.
func (kcp *KCP) update_ack(rtt int32) { _ = "STUB: not implemented"; return }

// if the new RTT sample is below the bottom of the range of
// what an RTT measurement is expected to be.
// give an 8x reduced weight versus its normal weighting

// shrink_buf advances snd_una to the oldest unacknowledged segment in snd_buf.
func (kcp *KCP) shrink_buf() { _ = "STUB: not implemented"; return }

// parse_ack marks a segment as acknowledged in snd_buf by sequence number.
// The segment is not removed immediately; it stays until snd_una advances past it,
// avoiding expensive shifts in the ring buffer.
func (kcp *KCP) parse_ack(sn uint32) { _ = "STUB: not implemented"; return }

// mark and free space, but leave the segment here,
// and wait until `una` to delete this, then we don't
// have to shift the segments behind forward,
// which is an expensive operation for large window

// parse_fastack increments the fast-ack counter for segments with sn < the given sn.
// Returns 1 if any segment's fastack counter has reached the fast retransmit threshold.
func (kcp *KCP) parse_fastack(sn, ts uint32) int { _ = "STUB: not implemented"; return 0 }

// parse_una removes all segments from snd_buf that have been cumulatively acknowledged
// (i.e., segments with sn < una). Returns the number of segments removed.
func (kcp *KCP) parse_una(una uint32) int { _ = "STUB: not implemented"; return 0 }

// ack append
func (kcp *KCP) ack_push(sn, ts uint32) { _ = "STUB: not implemented"; return }

// returns true if data has repeated
func (kcp *KCP) parse_data(newseg segment) bool { _ = "STUB: not implemented"; return false }

// replicate the content if it's new

// insert the new segment into rcv_buf

// move available data from rcv_buf -> rcv_queue

// push back segment

// Input a packet into kcp state machine.
//
// 'regular' indicates it's a real data packet from remote, and it means it's not generated from ReedSolomon
// codecs.
//
// 'ackNoDelay' will trigger immediate ACK, but surely it will not be efficient in bandwidth
func (kcp *KCP) Input(data []byte, pktType PacketType, ackNoDelay bool) int {
	_ = "STUB: not implemented"
	return 0
}

// the latest ack packet

// signal to flush segments

// BCE hint

// only trust window updates from regular packets. i.e: latest update

// delayed data copying

// ready to send back IKCP_CMD_WINS in Ikcp_flush
// tell remote my window size

// update rtt with the latest ts
// ignore the FEC packet

// Congestion window (cwnd) update on ACK arrival.
// Uses Reno-style algorithm: slow-start below ssthresh, then AIMD.

// Determine if we need to flush data segments or acks

// If window has slided or, a fastack should be triggered,
// Flush immediately. In previous implementations, we only
// send out fastacks when interval timeouts, so the resending packets
// have to wait until then. Now, we try to flush as soon as we can.

// clocking
// This serves as the clock for low-latency network.(i.e. the latency is less than the interval.)
// If the other end is waiting for confirmations, it has to want until the interval timeouts then
// the flush() is triggered to send out the una & acks. In low-latency network, the interval time is too long to wait,
// so acks have to be sent out immediately when there are too many.

// testing(xtaci): ack immediately if acNoDelay is set

func (kcp *KCP) wnd_unused() uint16 { _ = "STUB: not implemented"; return 0 }

// flush sends pending data through the KCP output callback.
// This is the core scheduling function, organized in 6 phases:
//
//	Phase 1: Flush pending ACKs
//	Phase 2: Window probing (when remote window is zero)
//	Phase 3: Send window probe commands (WASK/WINS)
//	Phase 4: Move segments from snd_queue to snd_buf (sliding window)
//	Phase 5: Retransmit segments (initial, fast, early, RTO)
//	Phase 6: Update SNMP counters and congestion window
//
// Returns the suggested interval (ms) until the next flush call.
func (kcp *KCP) flush(flushType FlushType) (nextUpdate uint32) { _ = "STUB: not implemented"; return 0 }

// makeSpace makes room for writing

// flush bytes in buffer if there is any

// --- Phase 1: Flush pending ACKs ---

// filter jitters caused by bufferbloat

// --- Phase 2: Window probing (when remote window is zero) ---

// --- Phase 3: Flush window probing commands ---

// flush window probing commands

// --- Phase 4: Move segments from snd_queue to snd_buf (sliding window) ---
// Effective window = min(snd_wnd, rmt_wnd, cwnd)

// calculate resent

// --- Phase 5: Retransmit segments from snd_buf ---
// Determines which segments need (re)transmission:
// - Initial transmit (xmit == 0)
// - Fast retransmit (fastack >= fastresend threshold)
// - Early retransmit (fastack > 0, no new segments queued)
// - RTO-based retransmit (current >= resendts)

// initial transmit

// fast retransmit

// must wait until RTO to reset

// early retransmit

// RTO

// mark connection as dead

// get the nearest rto

// --- Phase 6: Update SNMP counters and congestion window ---

// cwnd update

// Update ssthresh after fast retransmit.
// Rate halving per RFC 6937: ssthresh = inflight / 2

// Congestion control after RTO: reset cwnd per RFC 5681

// (deprecated)
//
// Update updates state (call it repeatedly, every 10ms-100ms), or you can ask
// ikcp_check when to call it again (without ikcp_input/_send calling).
// 'current' - current timestamp in millisec.
func (kcp *KCP) Update() { _ = "STUB: not implemented"; return }

// (deprecated)
//
// Check determines when should you invoke ikcp_update:
// returns when you should invoke ikcp_update in millisec, if there
// is no ikcp_input/_send calling. you can call ikcp_update in that
// time, instead of call update repeatly.
// Important to reduce unnacessary ikcp_update invoking. use it to
// schedule ikcp_update (eg. implementing an epoll-like mechanism,
// or optimize ikcp_update when handling massive kcp connections)
func (kcp *KCP) Check() uint32 { _ = "STUB: not implemented"; return 0 }

// SetMtu changes MTU size, default is 1400
func (kcp *KCP) SetMtu(mtu int) int { _ = "STUB: not implemented"; return 0 }

// NoDelay options
// fastest: ikcp_nodelay(kcp, 1, 20, 2, 1)
// nodelay: 0:disable(default), 1:enable
// interval: internal update timer interval in millisec, default is 100ms
// resend: 0:disable fast resend(default), 1:enable fast resend
// nc: 0:normal congestion control(default), 1:disable congestion control
func (kcp *KCP) NoDelay(nodelay, interval, resend, nc int) int { _ = "STUB: not implemented"; return 0 }

// WndSize sets maximum window size: sndwnd=32, rcvwnd=32 by default
func (kcp *KCP) WndSize(sndwnd, rcvwnd int) int { _ = "STUB: not implemented"; return 0 }

// WaitSnd gets how many packet is waiting to be sent
func (kcp *KCP) WaitSnd() int { _ = "STUB: not implemented"; return 0 }

// SetLogger configures the trace logger
func (kcp *KCP) SetLogger(mask KCPLogType, logger logoutput_callback) {
	_ = "STUB: not implemented"
	return
}
