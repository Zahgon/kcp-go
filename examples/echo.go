package main

import (
	"crypto/sha1"
	"log"

	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
)

func main() {
	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)

	listener, err := kcp.ListenWithOptions("127.0.0.1:12345", block, 10, 3)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()

	// spin-up the client
	go client()
	for {
		s, err := listener.AcceptKCP()
		if err != nil {
			log.Fatal(err)
			return
		}

		go handleEcho(s)
	}
}

// handleEcho send back everything it received
func handleEcho(conn *kcp.UDPSession) { _ = "STUB: not implemented"; return }

func client() { _ = "STUB: not implemented"; return }

// wait for server to become ready

// dial to the echo server

// read back the data
