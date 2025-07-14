package protocol

import "net"

type ChatProtocol struct{}

var _ Protocol = (*ChatProtocol)(nil)

func (c *ChatProtocol) HandleConnection(_ net.Conn, _ Broadcaster) error {
	return nil
}
