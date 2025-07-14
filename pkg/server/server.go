package server

import (
	"fmt"
	"net"
	"strconv"

	"github.com/PhilippReinke/tcp-to-http/pkg/connection"
	"github.com/PhilippReinke/tcp-to-http/pkg/logger"
	"github.com/PhilippReinke/tcp-to-http/pkg/protocol"
)

type Server struct {
	logger      *logger.Logger
	listener    net.Listener
	connManager *connection.Manager
	protocol    protocol.Protocol
}

func New(
	host string,
	port int,
	logger *logger.Logger,
	connManager *connection.Manager,
	protocol protocol.Protocol,
) (*Server, error) {
	addr := host + ":" + strconv.Itoa(int(port))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("create listener: %w", err)
	}

	return &Server{
		logger:      logger,
		listener:    listener,
		connManager: connManager,
		protocol:    protocol,
	}, nil
}

func (s *Server) Serve() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			s.logger.WithError(err).Error("Failed to establish connection.")
			continue
		}

		broadcastQueue := make(chan []byte, 1000)
		go func() {
			for data := range broadcastQueue {
				s.connManager.Broadcast(data)
			}
		}()

		go func() {
			defer func() {
				if err := s.connManager.CloseConnection(conn); err != nil {
					s.logger.WithError(err).Error("Failed to close connection.")
				}
			}()

			connectionInfo, err := s.connManager.Register(conn)
			if err != nil {
				s.logger.WithConnection(conn).WithError(err).
					Error("Failed to register connection.")
				return
			}

			broadcaster := protocol.Broadcaster{
				Receive: broadcastQueue,
				Send:    connectionInfo.SendToProtocol,
			}

			if err := s.protocol.HandleConnection(conn, broadcaster); err != nil {
				s.logger.WithError(err).Error("Failed to handle connection.")
				return
			}
		}()
	}
}

func (s *Server) Close() error {
	return s.listener.Close()
}
