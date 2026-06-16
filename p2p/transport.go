package p2p

import "net"

// Peer is an interface that represents the remote node
type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

// Could be any transport protocol(TCP, UDP, websockets)
type Transport interface { 
	Addr() string
	Dial(string) error
	ListenAndAccept() error
	Close() error
}