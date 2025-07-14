package protocol

import "net"

type Protocol interface {
	HandleConnection(conn net.Conn, broadcaster Broadcaster) error
}

type Broadcaster struct {
	Receive chan<- []byte
	Send    <-chan []byte
}
