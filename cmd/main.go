package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jazaltron10/Golang/weatherFC_APP/internal/server"
)

func main() {
	// Create a channel to gracefully shutdown your application
	gcQuit := make(chan os.Signal, 1)
	signal.Notify(gcQuit, os.Interrupt, syscall.SIGTERM)

	// Set up the Echo server
	s := server.NewServer()

	// Start the server in a goroutine
	go func() {
		if err := s.BeginServer(); err != nil {
			s.Logger.Errorf("Error starting server: %v", err)
			close(gcQuit)
		}
	}()

	// Wait for an interrupt signal (Ctrl+C) or SIGTERM to gracefully shutdown the server
	<-gcQuit

	// Gracefully shutdown the server
	if err := s.gracefulShutdown(); err != nil {
		s.Logger.Errorf("Error shutting down server: %v", err)
	}

	// Additional cleanup or logging if needed
}
