package p2p

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

// remote node
type TCPPeer struct {
	conn     net.Conn
	outbound bool //dial to retrieve a connection: outbound = true, if accept to retrieve: outbound = false;
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	handshakeFunc HandshakeFunc
	decoder       Decoder
	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

func NOPHanshakeFunc(any) error { return nil }

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		// handshakeFunc: func(any) error { return nil },
		handshakeFunc: NOPHanshakeFunc,
		listenAddress: listenAddr,
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
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
			continue // Continue accepting connections even if there's an error
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if t.handshakeFunc(conn); err != nil {

	}

	buf := new(bytes.Buffer)
	for {
		n, err := conn.Read(buf)
		//msg = buf[:n]
	}
	fmt.Printf("new incoming connection %+v\n", peer)
}
