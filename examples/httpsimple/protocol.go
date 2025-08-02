package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"

	"github.com/PhilippReinke/tcp-to-http/pkg/protocol"
)

type HTTPSimple struct{}

var _ protocol.Protocol = (*HTTPSimple)(nil)

func (HTTPSimple) HandleConnection(
	conn net.Conn,
	_ protocol.Broadcaster,
) error {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	_, err := readRequest(reader)
	if err != nil {
		return fmt.Errorf("read request: %w", err)
	}

	// always respond with 200 OK
	body := []byte("Hello from httpsimple server <3\n")
	res := &Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(body)),
		},
		Body: body,
	}

	if err = writeResponse(writer, res); err != nil {
		return fmt.Errorf("write response: %w", err)
	}

	return nil
}
