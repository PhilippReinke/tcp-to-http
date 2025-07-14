package connection

import (
	"errors"
	"net"
	"sync"
	"time"

	"github.com/PhilippReinke/tcp-to-http/pkg/logger"
)

var (
	ErrNotFound         = errors.New("connection not found")
	ErrAlreadyRegisterd = errors.New("connection already registered")
	ErrNoNewConnections = errors.New("no new connections are permitted")
)

type Manager struct {
	logger              *logger.Logger
	conns               map[net.Conn]ConnectionInfo
	allowNewConnections bool
}

type ConnectionInfo struct {
	SendToProtocol chan []byte
}

// NewManager creates new connection manager.
func NewManager(logger *logger.Logger) *Manager {
	return &Manager{
		logger:              logger,
		conns:               make(map[net.Conn]ConnectionInfo),
		allowNewConnections: true,
	}
}

// Register registers a new connection.
func (c *Manager) Register(conn net.Conn) (ConnectionInfo, error) {
	if !c.allowNewConnections {
		return ConnectionInfo{}, ErrNoNewConnections
	}

	_, ok := c.conns[conn]
	if ok {
		return ConnectionInfo{}, ErrAlreadyRegisterd
	}
	connectionInfo := ConnectionInfo{
		SendToProtocol: make(chan []byte),
	}
	c.conns[conn] = connectionInfo

	return connectionInfo, nil
}

// Broadcast a message to all known connections.
func (c *Manager) Broadcast(data []byte) {
	for conn, connInfo := range c.conns {
		go func(ch chan []byte) {
			select {
			case ch <- data:
			case <-time.After(2 * time.Second):
				// after 2 seconds the message is dropped
				c.logger.WithConnection(conn).Error("Broadcast timeout.")
			}
		}(connInfo.SendToProtocol)
	}
}

// CloseConnection closes a connection.
func (c *Manager) CloseConnection(conn net.Conn) error {
	_, ok := c.conns[conn]
	if !ok {
		return ErrNotFound
	}

	err := conn.Close()
	delete(c.conns, conn)
	return err
}

// Close will attempt to close all connections.
func (c *Manager) Close() error {
	c.allowNewConnections = false

	var wg sync.WaitGroup
	for conn := range c.conns {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := c.CloseConnection(conn); err != nil {
				c.logger.WithError(err).Error("Failed to close connection.")
				return
			}
			c.logger.WithConnection(conn).Info("Connection closed.")
		}()
	}

	wg.Wait()
	c.logger.Info("Closed all connections.")

	return nil
}
