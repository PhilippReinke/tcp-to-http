package protocol

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type EchoProtocol struct{}

var _ Protocol = (*EchoProtocol)(nil)

func (e EchoProtocol) HandleConnection(conn net.Conn) error {
	reader := bufio.NewReader(conn)

	welcome := "Echo server ready. Send me a message!\n"
	if _, err := conn.Write([]byte(welcome)); err != nil {
		return err
	}

	// read loop
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("read message: %w", err)
		}

		// echo message back
		echo := fmt.Sprintf("Echo: %s", message)
		_, err = conn.Write([]byte(echo))
		if err != nil {
			return fmt.Errorf("write response: %w", err)
		}
	}
}
