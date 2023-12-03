// internal/handler/weather.go
package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/jazaltron10/Golang/weatherFC_APP/configs"
	"github.com/jazaltron10/Golang/weatherFC_APP/internal/cache"
	// "github.com/jazaltron10/Golang/weatherFC_APP/internal/handler"
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

	// TODO - Jasper - Implement the functions below 

			/*

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

			cityForecast, err := wh.getWeatherForecast(coordinates)
			if err != nil {
				wh.logger.Errorf("Error fetching forecast for city %s: %v", city, err)
				continue
			}

			// Cache the forecast data
			err = wh.store.Set(city, cityForecast)
			if err != nil {
				wh.logger.Errorf("Error caching forecast for city %s: %v", city, err)
			}

			// Append the forecast to the list
			forecasts = append(forecasts, cityForecast...)
		}
	}

	// Construct the response JSON
	responseJSON := map[string]interface{}{"forecast": forecasts}

	
	// Send the response
	return c.JSON(http.StatusOK, responseJSON)
	*/
	return nil
}

// Additional helper functions...

// The additional helper functions (e.g., getCoordinates, getWeatherForecast) should be taken from the modified handler.go file.
