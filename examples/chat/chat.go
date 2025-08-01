package main

import (
	"net"

	"github.com/PhilippReinke/tcp-to-http/pkg/protocol"
)

type Chat struct{}

var _ protocol.Protocol = (*Chat)(nil)

func (c *Chat) HandleConnection(_ net.Conn, _ protocol.Broadcaster) error {
	return nil
}
