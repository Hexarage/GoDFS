package main

import (
	"log"

	p2p "github.com/Hexarage/GoDFS/peer2peer"
)

func main() {
	transport := p2p.NewTCPTransport(":3000")

	if err := transport.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
