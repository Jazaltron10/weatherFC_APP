// internal/handler/weather.go
package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/jazaltron10/Golang/weatherFC_APP/configs"
	"github.com/jazaltron10/Golang/weatherFC_APP/internal/cache"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// WeatherHandler handles HTTP requests related to weather forecasts.
type WeatherHandler struct {
	client *http.Client
	store  cache.Cache
	logger *logrus.Logger
}

// NewWeatherHandler creates a new instance of WeatherHandler.
func NewWeatherHandler(client *http.Client, store cache.Cache, logger *logrus.Logger) *WeatherHandler {
	return &WeatherHandler{
		client: client,
		store:  store,
		logger: logger,
	}
}

// GetWeatherForecastHandler handles the /weather endpoint and returns weather forecasts for the given cities.
func (wh *WeatherHandler) GetWeatherForecastHandler(c echo.Context) error {
	// Parse the list of cities from the query parameter
	citiesParam := c.QueryParam("city")
	cities := strings.Split(citiesParam, ",")

	// Initialize a slice to store the forecast for each city
	var forecasts []configs.ForecastPeriod

	// Iterate through each city and retrieve the forecast
	for _, city := range cities {
		// Check if forecast data is available in the cache
		cachedForecast, err := wh.store.Get(city)
		if err == nil && cachedForecast != nil {
			// Use cached forecast data if available
			forecasts = append(forecasts, cachedForecast...)
		} else {
			// Fetch and cache the forecast if not available in the cache
			coordinates, err := wh.getCoordinates(city)
			if err != nil {
				wh.logger.Errorf("Error fetching coordinates for city %s: %v", city, err)
				continue
			}

			cityForecast, err := wh.getWeatherForecast(&coordinates)
			if err != nil {
				// Handle unexpected EOF errors
				if strings.Contains(err.Error(), "unexpected EOF") {
					wh.logger.Errorf("Unexpected EOF error fetching forecast for city %s. Skipping.", city)
					continue
				}

				wh.logger.Errorf("Error fetching forecast for city %s: %v", city, err)
				continue
			}

			// Cache the forecast data
			err = wh.store.Set(city, []configs.ForecastPeriod{cityForecast})
			if err != nil {
				wh.logger.Errorf("Error caching forecast for city %s: %v", city, err)
			}

			// Append the forecast to the list
			forecasts = append(forecasts, cityForecast)
		}
	}

	// Send the response
	return c.JSON(http.StatusOK, forecasts)
}

// getCoordinates fetches the coordinates for a given city
func (wh *WeatherHandler) getCoordinates(city string) (configs.ForecastCoordinates, error) {
	// Construct the API endpoint for OpenStreetMap Nominatim without
	endpoint := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1", url.QueryEscape(city))
	// https://nominatim.openstreetmap.org/search?q=new%20york&format=json&limit=1
	// Parse the endpoint URL.
	coords:= []configs.ForecastCoordinates{{}}

	link, err := url.Parse(endpoint)
	if err != nil {
		return coords[0], err
	}
	req, _ :=http.NewRequest(http.MethodGet, link.String(), nil)
	res, _ :=wh.client.Do(req)

	defer res.Body.Close()

	jsonReps, _ := io.ReadAll(res.Body)

	
	_ = json.Unmarshal(jsonReps, &coords)
	
	return coords[0], nil
}

// getWeatherForecast fetches the weather forecast for a given set of coordinates
func (wh *WeatherHandler) getWeatherForecast(coordinates *configs.ForecastCoordinates) (configs.ForecastPeriod, error) {
	// Construct the API endpoint for a weather forecast service
	weatherEndpoint := fmt.Sprintf("https://api.weather.gov/points/%s,%s", coordinates.Latitude, coordinates.Longitude)
	// https://api.weather.gov/points?lat=34.0536909&lon=-118.242766
	// https://api.weather.gov/points/lat=34.0536909&lon=-118.242766
	// https://api.weather.gov/points/34.0537,-118.2428
	// Make a request to the weather API.
	response, err := http.Get(weatherEndpoint)
	if err != nil {
		return configs.ForecastPeriod{}, fmt.Errorf("error fetching data from weather API (%s): %v", weatherEndpoint, err)
	}
	defer response.Body.Close()

	// Decode the response JSON.
	var forecastData configs.ForecastPeriod
	err = json.NewDecoder(response.Body).Decode(&forecastData)
	if err != nil {
		return configs.ForecastPeriod{}, fmt.Errorf("error decoding JSON from weather API (%s): %v", weatherEndpoint, err)
	}

	return forecastData, nil
}

// Additional helper functions...

// The additional helper functions (e.g., getCoordinates, getWeatherForecast) should be taken from the modified handler.go file.

// http test in golang

// for testing handler
// new recorder
// new client
// new server
