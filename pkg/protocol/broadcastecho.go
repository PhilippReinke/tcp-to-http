package protocol

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type BroadcastEchoProtocol struct{}

var _ Protocol = (*BroadcastEchoProtocol)(nil)

func (e BroadcastEchoProtocol) HandleConnection(
	conn net.Conn,
	broadcaster Broadcaster,
) error {
	reader := bufio.NewReader(conn)

	welcome := "Broadcast echo server ready. Send me a message!\n"
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

		broadcaster.Receive <- []byte("Broadcast Echo: " + message)
	}
}
