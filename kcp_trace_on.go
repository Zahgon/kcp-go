//go:build debug

// only build tag debug is set, then debugLog will be enabled in compile time
package kcp

func (kcp *KCP) debugLog(logtype KCPLogType, args ...any) { _ = "STUB: not implemented"; return }
