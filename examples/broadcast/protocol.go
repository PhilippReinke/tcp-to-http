package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/PhilippReinke/tcp-to-http/pkg/protocol"
)

type Broadcast struct{}

var _ protocol.Protocol = (*Broadcast)(nil)

func (Broadcast) HandleConnection(
	conn net.Conn,
	broadcaster protocol.Broadcaster,
) error {
	reader := bufio.NewReader(conn)

	welcome := "Broadcast server ready. Send me a message!\n"
	if _, err := conn.Write([]byte(welcome)); err != nil {
		return err
	}

	// handle broadcast receive
	done := make(chan struct{})
	defer close(done)
	go func() {
		for {
			select {
			case message := <-broadcaster.Send:
				if _, err := conn.Write(message); err != nil {
					return
				}
			case <-done:
				return
			}
		}
	}()

	// read loop
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("read message: %w", err)
		}

		broadcaster.Receive <- []byte("Broadcast: " + message)
	}
}
