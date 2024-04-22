package p2p

// The remote node
type Peer interface{}

// Anything that handles the communication between
// the network nodes (UDP, TCP, Websockets, etc...)
type Transport interface {
	ListenAndAccept() error
}
