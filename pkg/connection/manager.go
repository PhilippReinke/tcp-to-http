package connection

import (
	"errors"
	"net"
	"sync"

	"github.com/PhilippReinke/tcp-to-http/pkg/logger"
)

var (
	ErrNotFound         = errors.New("connection not found")
	ErrAlreadyRegisterd = errors.New("connection already registered")
	ErrNoNewConnections = errors.New("no new connections are permitted")
)

type Manager struct {
	logger              *logger.Logger
	conns               map[net.Conn]struct{}
	allowNewConnections bool
}

// NewManager creates new connection manager.
func NewManager(logger *logger.Logger) *Manager {
	return &Manager{
		logger:              logger,
		conns:               make(map[net.Conn]struct{}),
		allowNewConnections: true,
	}
}

// Register registers a new connection.
func (c *Manager) Register(conn net.Conn) error {
	if !c.allowNewConnections {
		return ErrNoNewConnections
	}

	_, ok := c.conns[conn]
	if ok {
		return ErrAlreadyRegisterd
	}
	c.conns[conn] = struct{}{}

	return nil
}

// CloseConnection closes a connection.
func (c *Manager) CloseConnection(conn net.Conn) error {
	_, ok := c.conns[conn]
	if !ok {
		return ErrNotFound
	}

	return conn.Close()
}

// Close will attempt to close all connections.
func (c *Manager) Close() error {
	c.allowNewConnections = false

	var wg sync.WaitGroup
	for conn := range c.conns {
		wg.Add(1)
		go func() {
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
