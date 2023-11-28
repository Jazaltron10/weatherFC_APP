// server.go
package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/jazaltron10/Golang/weatherFC_APP/internal/cache"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/jazaltron10/Golang/weatherFC_APP/internal/handler"
)

type Server struct {
	e *echo.Echo
	h *handler.WeatherHandler
	l *logrus.Logger
}

func NewServer(store cache.InternalCache, l *logrus.Logger) *Server {
	eRouter := echo.New()

	eRouter.Use(middleware.Logger())
	eRouter.Use(middleware.Recover())

	// Create WeatherHandler with dependencies
	client := &http.Client{} // Customize the HTTP client as needed
	weatherHandler := handler.NewWeatherHandler(client, store, l)

	eRouter.GET("/weather", weatherHandler.GetWeatherForecastHandler)

	return &Server{
		e: eRouter,
		h: weatherHandler,
		l: l,
	}
}

func (s *Server) BeginServer(quit <-chan os.Signal) {
	s.e.Logger.SetLevel(logrus.INFO)

	go func() {
		if err := s.e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			s.l.Fatal("shutting down the server")
		}
	}()

	<-quit
	s.gracefulShutdown()
}

func (s *Server) gracefulShutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.e.Shutdown(ctx); err != nil {
		s.l.Fatal(err)
	}
}
