package main

import (
	"log"

	p2p "github.com/Hexarage/GoDFS/peer2peer"
)

func main() {
	tcpOptions := p2p.TCPTransportOptions{
		ListenAddress: ":3000",
		ShakeHands:    p2p.NOPHandshakeFunc,
		Decoder:       p2p.GOBDecoder{},
	}
	transport := p2p.NewTCPTransport(tcpOptions)

	if err := transport.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
