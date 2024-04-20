package p2p

// remote node
type Peer interface {
}

// transport handles comms between nodes
type Transport interface {
	ListenAndAccept() error
}
