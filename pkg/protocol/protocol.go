package protocol

import "net"

type Protocol interface {
	HandleConnection(conn net.Conn) error
}
