// api/handler.go
package api

import (
	"net/http"
	"strings"

	"github.com/jazaltron10/goAPI/weatherAPI/internal/weather"
	"github.com/labstack/echo/v4"
)


type Handler struct {
	c     *http.Client
	store cache.Cache
	l     *logrus.Logger
}


func (h *Handler) CreateClient(store cache.Cache, l *logrus.Logger) {
	h.c = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	h.l = l
	h.store = store
}


// internal/handler/handler.go
// -> type Handler struct {... Objects -> client, store, logger}
// instead of func GetWeatherForecastHandler -> func (h *Handler) GetWeatherForecastHandler
