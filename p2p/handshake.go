package p2p

type HandShakefunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }