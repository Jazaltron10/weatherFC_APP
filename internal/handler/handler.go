// internal/handler/handler.go
package handler

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jazaltron10/Golang/weatherFC_APP/configs"
	"github.com/jazaltron10/Golang/weatherFC_APP/internal/cache"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	c     *http.Client
	store cache.Cache
	l     *logrus.Logger
}

func NewHandler() *Handler {
	return &Handler{}
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

func (h *Handler) GetWeatherForecastHandler(c echo.Context) error {
	// Parse the list of cities from the query parameter
	citiesParam := c.QueryParam("city")
	cities := strings.Split(citiesParam, ",")

	// Initialize a slice to store the forecast for each city
	var forecasts []*configs.CityCountryEndpoint

	// Iterate through each city and retrieve the forecast
	for _, city := range cities {
		// Fetch URL for coordinates for each city
		coordinatesLink, err := h.getCoordinates(city)
		if err != nil {
			h.l.Errorf("Error fetching coordinates for city %s: %v", city, err)
			continue
		}

		// Fetch the weather forecast for the coordinates
		cityForecast, err := h.getWeatherForecast(coordinatesLink)
		if err != nil {
			h.l.Errorf("Error fetching forecast for city %s: %v", city, err)
			continue
		}

		// Append the forecast to the list
		forecasts = append(forecasts, cityForecast)
	}

	// Construct the response JSON
	responseJSON := map[string]interface{}{"forecast": forecasts}

	// Send the response
	return c.JSON(http.StatusOK, responseJSON)
}

func (h *Handler) getCoordinates(city string) (*url.URL, error) {
	endpoint := &configs.CityCountryEndpoint{
		City:   city,
		Format: "json", // Assuming JSON format for coordinates
	}

	link, err := endpoint.GetOpenStreetMapLink()
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (h *Handler) getWeatherForecast(coordinatesURL *url.URL) (*configs.CityCountryEndpoint, error) {
	forecastLink, err := h.getForecastURL(coordinatesURL)
	if err != nil {
		return nil, err
	}

	forecastData, err := h.fetchData(forecastLink)
	if err != nil {
		return nil, err
	}

	forcastInfo := &configs.CityCountryEndpoint{}

	err =json.Unmarshal(forecastData, forcastInfo)
	if err != nil{
		return nil, err
	}

	return forcastInfo, nil
}

func (h *Handler) getForecastURL(coordinates *url.URL) (*url.URL, error) {
	forecastCoordinates := &configs.ForecastCoordinates{
		Latitude:  coordinates.Query().Get("lat"),
		Longitude: coordinates.Query().Get("lon"),
	}

	link, err := forecastCoordinates.GetForecastCoordinatesLink()
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (h *Handler) fetchData(link *url.URL) ([]byte, error) {
	response, err := h.c.Get(link.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
