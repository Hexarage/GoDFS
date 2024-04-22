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

type TCPTransportOptions struct {
	ListenAddress string
	ShakeHands    HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOptions
	listener net.Listener

	mut  sync.RWMutex
	peer map[net.Addr]Peer
}

func NewTCPTransport(conf TCPTransportOptions) *TCPTransport {

	return &TCPTransport{
		TCPTransportOptions: conf,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddress)
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
	if err := t.ShakeHands(peer); err != nil {
		connection.Close()
		fmt.Printf("TCP handshake error %s\n", err)
		return
	}

	// Read loop
	msg := &Temp{}
	for {
		if err := t.Decoder.Decode(connection, msg); err != nil {
			fmt.Printf("TCP error %s\n", err)
			continue
		}
	}
}
