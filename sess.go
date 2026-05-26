// The MIT License (MIT)
//
// # Copyright (c) 2015 xtaci
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

// [THE GENERALIZED DATA PIPELINE FOR KCP-GO]
//
// Outgoing Data Pipeline:                        Incoming Data Pipeline:
// Stream          (Input Data)                   Packet Network  (Network Interface Card)
//   |                                               |
//   v                                               v
// KCP Output      (Reliable Transport Layer)     Reader/Listener (Reception Queue)
//   |                                               |
//   v                                               v
// FEC Encoding    (Forward Error Correction)     Decryption      (Data Security)
//   |                                               |
//   v                                               v
// CRC32 Checksum  (Error Detection)              CRC32 Checksum  (Error Detection)
//   |                                               |
//   v                                               v
// Encryption      (Data Security)                FEC Decoding    (Forward Error Correction)
//   |                                               |
//   v                                               v
// TxQueue         (Transmission Queue)           KCP Input       (Reliable Transport Layer)
//   |                                               |
//   v                                               v
// Packet Network  (Network Transmission)         Stream          (Input Data)

package kcp

import (
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
)

// Session-layer constants
const (
	// 16-bytes nonce for each packet
	nonceSize = 16

	// 4-bytes CRC32 checksum per packet
	crcSize = 4

	// overall crypto header size: nonce + CRC32
	cryptHeaderSize = nonceSize + crcSize

	// maximum packet size (Ethernet MTU)
	mtuLimit = 1500

	// conversation ID field size (bytes)
	convSize = 4

	// accept backlog: max pending connections for Listener
	acceptBacklog = 128

	// devBacklog: channel buffer size for post-processing pipeline
	devBacklog = 2048

	// max latency for consecutive FEC encoding (ms).
	// If the interval between two data packets exceeds this,
	// parity generation is skipped.
	maxFECEncodeLatency = 500

	// max number of packets batched in a single sendmmsg/writev call
	maxBatchSize = 64
)

var (
	errInvalidOperation = errors.New("invalid operation")
	errTimeout          = timeoutError{}
	errNotOwner         = errors.New("not the owner of this connection")
)

// timeoutError implements net.Error
type timeoutError struct{}

func (timeoutError) Error() string { _ = "STUB: not implemented"; return "" }
func (timeoutError) Timeout() bool { _ = "STUB: not implemented"; return false }
func (timeoutError) Temporary() bool {
	_ = "STUB: not implemented"

	// sendRequest defines a write request before encoding and transmission
	return false
}

type sendRequest struct {
	buffer []byte
	oob    bool
}

// OOB callback function
type OOBCallBackType func([]byte)

type (
	// UDPSession defines a KCP session implemented by UDP
	UDPSession struct {
		conn    net.PacketConn // the underlying packet connection
		ownConn bool           // true if we created conn internally, false if provided by caller
		kcp     *KCP           // KCP ARQ protocol
		l       *Listener      // pointing to the Listener object if it's been accepted by a Listener
		block   BlockCrypt     // block encryption object

		// kcp receiving is based on packets
		// recvbuf turns packets into stream
		recvbuf []byte
		bufptr  []byte

		// FEC codec
		fecDecoder *fecDecoder
		fecEncoder *fecEncoder

		// settings
		remote     net.Addr     // remote peer address
		rd         atomic.Value // read deadline
		wd         atomic.Value // write deadline
		headerSize int          // the header size additional to a KCP frame
		ackNoDelay bool         // send ack immediately for each incoming packet(testing purpose)
		writeDelay bool         // delay kcp.flush() for Write() for bulk transfer
		dup        int          // duplicate udp packets(testing purpose)

		// notifications
		die          chan struct{} // notify current session has Closed
		dieOnce      sync.Once
		chReadEvent  chan struct{} // notify Read() can be called without blocking
		chWriteEvent chan struct{} // notify Write() can be called without blocking

		// socket error handling
		socketReadError      atomic.Value
		socketWriteError     atomic.Value
		chSocketReadError    chan struct{}
		chSocketWriteError   chan struct{}
		socketReadErrorOnce  sync.Once
		socketWriteErrorOnce sync.Once

		// packets waiting to be sent on wire
		chPostProcessing chan sendRequest

		// platform-dependent optimizations
		platform platform

		// rate limiter (bytes per second)
		rateLimiter atomic.Value

		mu sync.Mutex

		// callbackForOOB is an optional callback for handling received out-of-band (OOB) data.
		//
		// OOB data bypasses the KCP reliable data path and is delivered unreliably.
		// The callback is invoked synchronously from the KCP input processing path.
		callbackForOOB atomic.Value
	}

	setReadBuffer interface {
		SetReadBuffer(bytes int) error
	}

	setWriteBuffer interface {
		SetWriteBuffer(bytes int) error
	}

	setDSCP interface {
		SetDSCP(int) error
	}
)

// newUDPSession create a new udp session for client or server
func newUDPSession(conv uint32, dataShards, parityShards int, l *Listener, conn net.PacketConn, ownConn bool, remote net.Addr, block BlockCrypt) *UDPSession {
	_ = "STUB: not implemented"
	return nil
}

// calculate additional header size introduced by encryption

// FEC codec initialization

// calculate additional header size introduced by FEC

// A basic check for the minimum packet size

// make a copy

// copy the data to a new buffer, and reserve header space

// delivery to post processing (non-blocking to avoid deadlock under lock)

// drop and recycle to avoid blocking; KCP will retransmit if needed

// Set Default MTU

// create post-processing goroutine

// it's a client connection

// start per-session updater

// Read implements net.Conn
func (s *UDPSession) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// deadline for current reading operation

// Pre-Go 1.23: Reset does not drain the channel;
// callers must drain at the goto-site before arriving here.

// disable timeout select case

// bufptr points to the current position of recvbuf,
// if previous 'b' is insufficient to accommodate the data, the
// remaining data will be stored in bufptr for next read.

// peek data size from kcp
// if 'b' is large enough to accommodate the data, read directly
// from kcp.recv() to 'b', like 'DMA'.

// otherwise, read to recvbuf first, then copy to 'b'.
// dynamically adjust the buffer size to the maximum of 'packet size' when necessary.

// usually recvbuf has a size of maximum packet size

// resize the length of recvbuf to match the data size

// read data to recvbuf first
// then copy bytes to 'b' as many as possible
// pointer update

// if it runs here, that means we have to block the call, and wait until the
// next data packet arrives.

// Write implements net.Conn
func (s *UDPSession) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// WriteBuffers write a vector of byte slices to the underlying connection
func (s *UDPSession) WriteBuffers(v [][]byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Pre-Go 1.23: Reset does not drain the channel;
// callers must drain at the goto-site before arriving here.

// disable timeout select case

// check for connection close and socket error

// make sure write do not overflow the max sliding window on both side

// transmit all data sequentially, make sure every packet size is within 'mss'

// handle each slice for packet splitting

// put the packets on wire immediately if the inflight window is full
// or if we've specified write no delay(NO merging of outgoing bytes)
// we don't have to wait until the periodical update() procedure uncorks.

// if it runs here, that means we have to block the call, and wait until the
// transmit buffer to become available again.

func (s *UDPSession) isClosed() bool { _ = "STUB: not implemented"; return false }

// Close closes the connection.
func (s *UDPSession) Close() error { _ = "STUB: not implemented"; return nil }

// try best to send all queued messages especially the data in txqueue

// belongs to listener

// client socket close

// LocalAddr returns the local network address. The Addr returned is shared by all invocations of LocalAddr, so do not modify it.
func (s *UDPSession) LocalAddr() net.Addr {
	_ = "STUB: not implemented"
	return *

	// RemoteAddr returns the remote network address. The Addr returned is shared by all invocations of RemoteAddr, so do not modify it.
	new(net.Addr)
}

func (s *UDPSession) RemoteAddr() net.Addr {
	_ = "STUB: not implemented"

	// SetDeadline sets the deadline associated with the listener. A zero time value disables the deadline.
	return *new(net.Addr)
}

func (s *UDPSession) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetReadDeadline implements the Conn SetReadDeadline method.
func (s *UDPSession) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (s *UDPSession) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetWriteDelay delays write for bulk transfer until the next update interval
func (s *UDPSession) SetWriteDelay(delay bool) { _ = "STUB: not implemented"; return }

// SetWindowSize set maximum window size
func (s *UDPSession) SetWindowSize(sndwnd, rcvwnd int) { _ = "STUB: not implemented"; return }

// SetMtu sets the maximum transmission unit(not including UDP header)
func (s *UDPSession) SetMtu(mtu int) bool { _ = "STUB: not implemented"; return false }

// kcp mtu is not including udp header

// Deprecated: toggles the stream mode on/off
func (s *UDPSession) SetStreamMode(enable bool) { _ = "STUB: not implemented"; return }

// SetACKNoDelay changes ack flush option, set true to flush ack immediately,
func (s *UDPSession) SetACKNoDelay(nodelay bool) { _ = "STUB: not implemented"; return }

// (deprecated)
//
// SetDUP duplicates udp packets for kcp output.
func (s *UDPSession) SetDUP(dup int) { _ = "STUB: not implemented"; return }

// SetNoDelay calls nodelay() of kcp
// https://github.com/skywind3000/kcp/blob/master/README.en.md#protocol-configuration
func (s *UDPSession) SetNoDelay(nodelay, interval, resend, nc int) {
	_ = "STUB: not implemented"
	return
}

// SetDSCP sets the 6bit DSCP field in IPv4 header, or 8bit Traffic Class in IPv6 header.
//
// if the underlying connection has implemented `func SetDSCP(int) error`, SetDSCP() will invoke
// this function instead.
//
// It has no effect if it's accepted from Listener.
func (s *UDPSession) SetDSCP(dscp int) error { _ = "STUB: not implemented"; return nil }

// interface enabled

// SetReadBuffer sets the socket read buffer, no effect if it's accepted from Listener
func (s *UDPSession) SetReadBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

// SetWriteBuffer sets the socket write buffer, no effect if it's accepted from Listener
func (s *UDPSession) SetWriteBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

// SetRateLimit sets the rate limit for this session in bytes per second,
// by setting to 0 will disable rate limiting.
func (s *UDPSession) SetRateLimit(bytesPerSecond uint32) { _ = "STUB: not implemented"; return }

// SetLogger configures the kcp trace logger
func (s *UDPSession) SetLogger(mask KCPLogType, logger logoutput_callback) {
	_ = "STUB: not implemented"
	return
}

// Control applys a procedure to the underly socket fd.
// CAUTION: BE VERY CAREFUL TO USE THIS FUNCTION, YOU MAY BREAK THE PROTOCOL.
func (s *UDPSession) Control(f func(conn net.PacketConn) error) error {
	_ = "STUB: not implemented"
	return nil
}

// postProcess is the goroutine that handles the outgoing packet pipeline.
// It runs the following stages sequentially for each packet:
//  1. FEC encoding   — generate parity shards (Reed-Solomon)
//  2. Encryption     — AEAD (e.g. AES-GCM) or CFB mode with CRC32
//  3. TX batching    — accumulate packets and flush via sendmmsg/writev
//
// Pipeline: KCP output -> chPostProcessing -> [FEC] -> [Encrypt] -> TxQueue -> Network
func (s *UDPSession) postProcess() { _ = "STUB: not implemented"; return }

// dequeue from post processing

// --- Stage 1: FEC encoding ---

// --- Stage 2: Encryption ---
// Two modes supported:
//   - AEAD (e.g. AES-GCM): nonce + authenticated ciphertext, no separate CRC
//   - CFB (legacy block ciphers): random nonce + CRC32 checksum + CFB encryption

// AEAD mode

// Cipher Feedback (CFB) mode

// --- Stage 3: TX batching ---

// original copy, move buf to txqueue directly

// dup copies for testing if set

// parity

// transmit when chPostProcessing is empty or we've reached max batch size

// WaitN only returns error if the limiter is misconfigured
// or context is cancelled. In either case, we continue sending.

// recycle

// re-enable die channel

// remaining packets in txqueue should be sent out

// block chDie temporarily

// sess update to trigger protocol
func (s *UDPSession) update() { _ = "STUB: not implemented"; return }

// self-synchronized timed scheduling

// GetConv gets conversation id of a session
func (s *UDPSession) GetConv() uint32 {
	_ = "STUB: not implemented"

	// GetRTO gets current rto of the session
	return 0
}

func (s *UDPSession) GetRTO() uint32 { _ = "STUB: not implemented"; return 0 }

// GetSRTT gets current srtt of the session
func (s *UDPSession) GetSRTT() int32 { _ = "STUB: not implemented"; return 0 }

// GetRTTVar gets current rtt variance of the session
func (s *UDPSession) GetSRTTVar() int32 { _ = "STUB: not implemented"; return 0 }

// SetOOBHandler registers a callback for receiving out-of-band (OOB) data.
//
// OOB data is delivered unreliably and bypasses the KCP reliable data path.
// The callback is invoked synchronously from the KCP input processing path.
//
// The callback MUST return quickly and MUST NOT perform any blocking operations.
// Blocking inside the callback will stall processing of all other KCP packets.
//
// Passing a nil callback unregisters the current OOB callback.
//
// OOB support requires FEC to be enabled, as the OOB packet format
// reuses the FEC header layout for demultiplexing.
func (s *UDPSession) SetOOBHandler(callback OOBCallBackType) error {
	_ = "STUB: not implemented"
	return nil
}

// GetOOBMaxSize returns the maximum payload size for an OOB packet.
//
// The returned value is the maximum number of bytes that can be carried as
// OOB data in a single packet, based on the current MTU and protocol layout.
//
// If FEC is not enabled, OOB is unsupported and this function returns 0.
func (s *UDPSession) GetOOBMaxSize() int { _ = "STUB: not implemented"; return 0 }

// Packet layout: | conv (4B) | OOB payload |

// SendOOB sends an out-of-band (OOB) data packet.
//
// OOB packets:
//   - Are unreliable: they are NOT retransmitted if lost.
//   - Are unordered: delivery order is not guaranteed.
//   - Are unacknowledged: no ACKs are generated.
//   - Bypass the KCP reliable data path.
//   - Reuse the FEC header layout for demultiplexing, but are NOT protected by FEC.
//
// The OOB payload MUST fit into a single packet.
// If the payload is too large, an error is returned.
//
// If the internal send queue is full, the OOB packet is dropped silently.
func (s *UDPSession) SendOOB(data []byte) error { _ = "STUB: not implemented"; return nil }

// lock the session during OOB packet construction

// Packet layout: | conv (4B) | OOB payload |

// Allocate buffer with reserved header space.
// s.headerSize includes the space needed by the FEC encoder.

// Encode conversation ID.

// Copy OOB payload immediately after the conversation ID.

// Enqueue the packet for post-processing.
// Performs OOB framing, encryption, and transmission, bypassing FEC and KCP.

// Session is closing.

// Drop silently to avoid blocking the sender.
// OOB delivery is best-effort by design.

func (s *UDPSession) notifyReadEvent() { _ = "STUB: not implemented"; return }

func (s *UDPSession) notifyWriteEvent() { _ = "STUB: not implemented"; return }

func (s *UDPSession) notifyReadError(err error) { _ = "STUB: not implemented"; return }

func (s *UDPSession) notifyWriteError(err error) { _ = "STUB: not implemented"; return }

// -----------------------------------------------------------------------
// Packet input pipeline (decryption -> integrity check -> FEC -> KCP)
// -----------------------------------------------------------------------

// packetInput is the entry point for incoming packets.
// It handles decryption and CRC32 verification before passing data to kcpInput.
//
// Pipeline: Network -> [Decrypt] -> [CRC32] -> kcpInput
func (s *UDPSession) packetInput(data []byte) { _ = "STUB: not implemented"; return }

// decryption and crc32 check

// basic check for minimum packet size
// NOTE: OOB allows sending small packets and even empty packets.

// kcpInput routes a decrypted packet into the KCP state machine,
// handling FEC decoding and OOB delivery.
//
// Packet demultiplexing uses the 16-bit field at offset 4:
//   - 0xf1 (typeData) / 0xf2 (typeParity): FEC-encoded packet
//   - 0xf3 (typeOOB): out-of-band packet (unreliable, bypasses KCP)
//   - other values: raw KCP packet (no FEC)
//
// Note: KCP cmd values [81-84] with frg [0-255] do not collide with
// FEC type markers 0x00f1/0x00f2/0x00f3 in little-endian.
func (s *UDPSession) kcpInput(data []byte) { _ = "STUB: not implemented"; return }

// 16bit kcp cmd [81-84] and frg [0-255] will not overlap with FEC type 0x00f1 0x00f2

// packet with FEC

// lock

// if fecDecoder is not initialized, create one with default parameter
// lazy initialization

// KCP input for data packets
// only data packets are fed into kcp directly
// parity packets are only used for recovery

// FEC decoding
// If there're some packets recovered from FEC, feed them into kcp

// must be larger than 2bytes

// recycle the buffer

// to notify the readers to receive the data if there's any

// to notify the writers if the window size allows to send more packets
// and the remote window size is not full.

// Count received OOB packet

// If an OOB callback is registered, invoke it synchronously.
// The callback is responsible for ensuring non-blocking behavior.

// Data layout: | FEC header (fecHeaderSizePlus2) | conv (4B) | OOB payload |

// packet without FEC

// -----------------------------------------------------------------------
// Listener: server-side session multiplexer
// -----------------------------------------------------------------------

type (
	// Listener defines a server which will be waiting to accept incoming connections
	Listener struct {
		block        BlockCrypt     // block encryption
		dataShards   int            // FEC data shard
		parityShards int            // FEC parity shard
		conn         net.PacketConn // the underlying packet connection
		ownConn      bool           // true if we created conn internally, false if provided by caller

		sessions    map[string]*UDPSession // all sessions accepted by this Listener
		sessionLock sync.RWMutex
		chAccepts   chan *UDPSession // Listen() backlog

		die     chan struct{} // notify the listener has closed
		dieOnce sync.Once

		// socket error handling
		socketReadError     atomic.Value
		chSocketReadError   chan struct{}
		socketReadErrorOnce sync.Once

		rd atomic.Value // read deadline for Accept()
	}
)

// packetInput is the Listener's packet input handler.
// It decrypts the packet, demultiplexes by remote address,
// and dispatches to existing sessions or creates new ones.
func (l *Listener) packetInput(data []byte, addr net.Addr) { _ = "STUB: not implemented"; return }

// decryption and crc32 check

// basic check for minimum packet size
// NOTE: OOB allows sending small packets and even empty packets.

// look for existing session

// try to get conversation id from the packet
// 16bit kcp cmd [81-84] and frg [0-255] will not overlap with FEC type 0x00f1 0x00f2

// data packet of FEC, conversation id inside

// parity packet of FEC, conversation id inside

// OOB packets always carry the conversation ID immediately after the FEC header.

// Data layout: | FEC header (fecHeaderSizePlus2) | conv (4B) | OOB payload |

// packet without FEC
// basic check for minimum kcp packet size

// on an existing connection

// If we have a valid conversation id or we cannot get conversation id from the packet,
// just feed the data into the existing session.

// conversation id mismatched, only accept reset packet with sn == 0

// Close will remove the session from listener's session map,
// So we can create a new session with the same addr below.

// The connection does not exist, try to create a new one.
// But if we don't have a valid conversation id, nothing we can do here except dropping the packet.

// Now we have a valid conversation id here without a session object, create a new session.
// do not let the new sessions overwhelm accept queue

// new session

func (l *Listener) notifyReadError(err error) { _ = "STUB: not implemented"; return }

// propagate read error to all sessions

// SetReadBuffer sets the socket read buffer for the Listener
func (l *Listener) SetReadBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

// SetWriteBuffer sets the socket write buffer for the Listener
func (l *Listener) SetWriteBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

// SetDSCP sets the 6bit DSCP field in IPv4 header, or 8bit Traffic Class in IPv6 header.
//
// if the underlying connection has implemented `func SetDSCP(int) error`, SetDSCP() will invoke
// this function instead.
func (l *Listener) SetDSCP(dscp int) error {
	_ = "STUB: not implemented"
	// interface enabled
	return nil
}

// Accept implements the Accept method in the Listener interface; it waits for the next call and returns a generic Conn.
func (l *Listener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *

	// AcceptKCP accepts a KCP connection
	new(net.Conn), nil
}

func (l *Listener) AcceptKCP() (*UDPSession, error) { _ = "STUB: not implemented"; return nil, nil }

// SetDeadline sets the deadline associated with the listener. A zero time value disables the deadline.
func (l *Listener) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetReadDeadline implements the Conn SetReadDeadline method.
func (l *Listener) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (l *Listener) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// Close stops listening on the UDP address, and closes the socket
func (l *Listener) Close() error { _ = "STUB: not implemented"; return nil }

// Control applys a procedure to the underly socket fd.
// CAUTION: BE VERY CAREFUL TO USE THIS FUNCTION, YOU MAY BREAK THE PROTOCOL.
func (l *Listener) Control(f func(conn net.PacketConn) error) error {
	_ = "STUB: not implemented"
	return nil
}

// closeSession notify the listener that a session has closed
func (l *Listener) closeSession(remote net.Addr) (ret bool) {
	_ = "STUB: not implemented"
	return false
}

// Addr returns the listener's network address, The Addr returned is shared by all invocations of Addr, so do not modify it.
func (l *Listener) Addr() net.Addr {
	_ = "STUB: not implemented"
	return *

	// -----------------------------------------------------------------------
	// Public API: Dial, Listen, and connection factory functions
	// -----------------------------------------------------------------------
	new(net.Addr)
}

// Listen listens for incoming KCP packets addressed to the local address laddr on the network "udp",
func Listen(laddr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

// ListenWithOptions listens for incoming KCP packets addressed to the local address laddr on the network "udp" with packet encryption.
//
// 'block' is the block encryption algorithm to encrypt packets.
//
// 'dataShards', 'parityShards' specify how many parity packets will be generated following the data packets.
//
// Check https://github.com/klauspost/reedsolomon for details
func ListenWithOptions(laddr string, block BlockCrypt, dataShards, parityShards int) (*Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ServeConn serves KCP protocol for a single packet connection.
func ServeConn(block BlockCrypt, dataShards, parityShards int, conn net.PacketConn) (*Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func serveConn(block BlockCrypt, dataShards, parityShards int, conn net.PacketConn, ownConn bool) (*Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Dial connects to the remote address "raddr" on the network "udp" without encryption and FEC
func Dial(raddr string) (net.Conn, error) { _ = "STUB: not implemented"; return *new(net.Conn), nil }

// DialWithOptions connects to the remote address "raddr" on the network "udp" with packet encryption
//
// 'block' is the block encryption algorithm to encrypt packets.
//
// 'dataShards', 'parityShards' specify how many parity packets will be generated following the data packets.
//
// Check https://github.com/klauspost/reedsolomon for details
func DialWithOptions(raddr string, block BlockCrypt, dataShards, parityShards int) (*UDPSession, error) {
	_ = "STUB: not implemented"
	// network type detection
	return nil, nil
}

// NewConn4 establishes a session and talks KCP protocol over a packet connection.
func NewConn4(convid uint32, raddr net.Addr, block BlockCrypt, dataShards, parityShards int, ownConn bool, conn net.PacketConn) (*UDPSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewConn3 establishes a session and talks KCP protocol over a packet connection.
func NewConn3(convid uint32, raddr net.Addr, block BlockCrypt, dataShards, parityShards int, conn net.PacketConn) (*UDPSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewConn2 establishes a session and talks KCP protocol over a packet connection.
func NewConn2(raddr net.Addr, block BlockCrypt, dataShards, parityShards int, conn net.PacketConn) (*UDPSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewConn establishes a session and talks KCP protocol over a packet connection.
func NewConn(raddr string, block BlockCrypt, dataShards, parityShards int, conn net.PacketConn) (*UDPSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
