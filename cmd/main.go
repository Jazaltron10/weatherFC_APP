package main

// Import necessary packages
import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jazaltron10/Golang/weatherFC_APP/internal/cache" // Import the package for your cache implementation
	"github.com/jazaltron10/Golang/weatherFC_APP/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	// Create a channel for graceful shutdown
	quit := make(chan os.Signal, 1) // channel of type os.signal that can store only 1 message at a time
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Initialize logger
	logger := logrus.New()

	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	// Initialize cache instance (replace with your cache initialization logic)
	cacheInstance := cache.NewInternalCache()

	// Create server with the initialized logger and cache instance
	s := server.NewServer(cacheInstance, logger)
	s.BeginServer(quit)

}

// Hello, world or 你好， 世界 or Καλημέρα κόσμε or こ んに ち は世界\n")
// }
