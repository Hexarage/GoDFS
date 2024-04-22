package p2p

// TODO: Figure out how to explain this vague idea
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
