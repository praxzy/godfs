package p2p

import (
	"net"
	"sync"
)

// TCPPeer represents the remote node over an established TCP connection.
type TCPPeer struct {
	// The underlying connection of the peer.
	// TCP in this case
	net.Conn

	// Dial() => outbound == true
	// Accept() => outbound == false 
	outbound bool

	wg *sync.WaitGroup
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		Conn: conn,
		outbound: outbound,
		wg: &sync.WaitGroup{},
	}
}

func (p *TCPPeer) CloseStream() {
	p.wg.Done()
}

func (p *TCPPeer) Send(b []byte) error {
	_, err := p.Conn.Write(b)
	return err
}

// type TCPTransport struct {
// 	listenAddress string
// 	listener net.Listener

// 	mu sync.RWMutex
// 	peers map[net.Addr]Peer

// }

// func NewTCPTransport(listenAddr string) *TCPTransport {
// 	return &TCPTransport{
// 		listenAddress: listenAddr,
// 	}
// }
