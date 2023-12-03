// server.go
package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/jazaltron10/Golang/weatherFC_APP/internal/cache"
	"github.com/jazaltron10/Golang/weatherFC_APP/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type Server struct {
	e *echo.Echo
	h *handler.WeatherHandler
	l *logrus.Logger
}

func NewServer(store cache.Cache, l *logrus.Logger) *Server {
	eRouter := echo.New()

	eRouter.Use(middleware.Logger())
	eRouter.Use(middleware.Recover())

	// Create WeatherHandler with dependencies
	client := &http.Client{} // Customize the HTTP client as needed
	weatherHandler := handler.NewWeatherHandler(client, store, l)

	eRouter.GET("/weather", func(c echo.Context) error {
		return weatherHandler.GetWeatherForecastHandler(c)
	})

	eRouter.GET("/test",  func(c echo.Context) error {
		return  c.JSON(http.StatusOK, []byte("hello this is Jasper!! "))
	})

	return &Server{
		e: eRouter,
		h: weatherHandler,
		l: l,
	}
}

func (s *Server) BeginServer(quit chan os.Signal) {
	s.e.Logger.SetLevel(log.INFO)

	go func() {
		if err := s.e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			s.l.Fatal("shutting down the server")
		}
	}()
	// Wait for an interrupt signal (Ctrl+C) or SIGTERM to gracefully shutdown the server
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


/*
checkout 
go Panic -> fatal handles panics 
graceful shutdown

*/

// "aGVsbG8gdGhpcyBpcyBKYXNwZXIhISA="

// https://www.base64decode.org/
