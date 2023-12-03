package main
// Import necessary packages
import (
    "os"
    "os/signal"
    "syscall"
    "github.com/sirupsen/logrus"
    "github.com/jazaltron10/Golang/weatherFC_APP/internal/server"
    "github.com/jazaltron10/Golang/weatherFC_APP/internal/cache" // Import the package for your cache implementation
)

func main() {
    // Create a channel for graceful shutdown
    quit := make(chan os.Signal, 1)
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
