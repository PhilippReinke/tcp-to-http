package main

import (
	"flag"
	"os"

	"github.com/PhilippReinke/tcp-to-http/pkg/connection"
	"github.com/PhilippReinke/tcp-to-http/pkg/logger"
	"github.com/PhilippReinke/tcp-to-http/pkg/protocol"
	"github.com/PhilippReinke/tcp-to-http/pkg/server"
)

func main() {
	host := flag.String("host", "localhost", "hostname for server")
	port := flag.Int("port", 8080, "port for server")
	flag.Parse()

	appLogger := logger.New(logger.Debug)
	appLogger.Info("Created logger.")

	connManager := connection.NewManager(appLogger)
	appLogger.Info("Created connection manager.")

	srv, err := server.New(
		*host, *port,
		appLogger,
		connManager,
		protocol.EchoProtocol{},
	)
	if err != nil {
		appLogger.WithError(err).Error("Failed to create server.")
		os.Exit(1)
	}
	appLogger.Info("Created server.")

	if err := srv.Serve(); err != nil {
		appLogger.WithError(err).Error("Failed to serve.")
		os.Exit(1)
	}
}
