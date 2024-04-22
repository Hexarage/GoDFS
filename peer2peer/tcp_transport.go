package p2p

import (
	"fmt"
	"net"
	"sync"
)

// Represents the remote node over a TCP established connection
type TCPPeer struct {
	// The underlying connection of the peer
	connection net.Conn

	// if we dial and retreive a connection		-> outbound = true
	// if we accept and retreive a connection	-> outbound = false
	outbound bool
}

func NewTCPPeer(connection net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		connection: connection,
		outbound:   outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	decoder       Decoder

	mut  sync.RWMutex
	peer map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {

	return &TCPTransport{
		shakeHands:    NOPHandshakeFunc,
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		connection, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP Accept Error: %s\n", err)
		}

		fmt.Printf("New incoming connection %s\n", connection)
		go t.handleConnection(connection)
	}
}

type Temp struct{} //placeholder

func (t *TCPTransport) handleConnection(connection net.Conn) {
	peer := NewTCPPeer(connection, true)
	if err := t.shakeHands(peer); err != nil {
		connection.Close()
		fmt.Printf("TCP handshake error %s\n", err)
		return
	}

	// Read loop
	msg := &Temp{}
	for {
		if err := t.decoder(connection, msg); err != nil {
			fmt.Printf("TCP error %s\n", err)
			continue
		}
	}
}
