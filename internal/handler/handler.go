// internal/handler/handler.go
package handler

import (
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jazaltron10/Golang/weatherFC_APP/internal/cache"
	"github.com/jazaltron10/Golang/weatherFC_APP/internal/forecast"
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
	var forecasts []forecast.CityForcast

	// Iterate through each city and retrieve the forecast
	for _, city := range cities {
		// Fetch coordinates for the city
		coordinates, err := h.getCoordinates(city)
		if err != nil {
			h.l.Errorf("Error fetching coordinates for city %s: %v", city, err)
			continue
		}

		// Fetch the weather forecast for the coordinates
		cityForecast, err := h.getWeatherForecast(coordinates)
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
	/*
	endpoint := &forecast.CityCountryEndpoint{
		City:   city,
		Format: "json", // Assuming JSON format for coordinates
	}

	
	link, err := endpoint.GetOpenStreetMapLink()
	if err != nil {
		return nil, fmt.Errorf("error getting coordinates for city %s: %v", city, err)
	}

	return link, nil
	*/
	return nil, nil
}

func (h *Handler) getWeatherForecast(coordinates *url.URL) (forecast.CityForcast, error) {
	/*
	forecastLink, err := h.getForecastURL(coordinates)
	
	if err != nil {
		return forecast.CityForecast{}, fmt.Errorf("error getting forecast for coordinates %s: %v", coordinates.String(), err)
	}

	forecastData, err := h.fetchData(forecastLink)
	if err != nil {
		return forecast.CityForecast{}, fmt.Errorf("error fetching forecast data for coordinates %s: %v", coordinates.String(), err)
	}

	cityForecast, err := forecast.ParseForecast(forecastData)
	if err != nil {
		return forecast.CityForecast{}, fmt.Errorf("error parsing forecast data for coordinates %s: %v", coordinates.String(), err)
	}

	return cityForecast, nil
	*/
	// TODO
	return forecast.CityForcast{}, nil
}

func (h *Handler) getForecastURL(coordinates *url.URL) (*url.URL, error) {
	/*
	forecastCoordinates := &forecast.ForecastCoordinates{
		Latitude:  coordinates.Query().Get("lat"),
		Longitude: coordinates.Query().Get("lon"),
	}

	link, err := forecastCoordinates.GetForecastCoordinatesLink()
	if err != nil {
		return nil, fmt.Errorf("error getting forecast URL for coordinates %s: %v", coordinates.String(), err)
	}

	return link, nil
}

func (h *Handler) fetchData(link *url.URL) ([]byte, error) {
	response, err := h.c.Get(link.String())
	if err != nil {
		return nil, fmt.Errorf("error fetching data from URL %s: %v", link.String(), err)
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading data from response body: %v", err)
	}

	return data, nil
	*/

	return nil, nil
}
