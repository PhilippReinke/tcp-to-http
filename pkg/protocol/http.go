package protocol

import "net"

type HTTPProtocol struct{}

var _ Protocol = (*HTTPProtocol)(nil)

func (c *HTTPProtocol) HandleConnection(_ net.Conn) error {
	return nil
}
