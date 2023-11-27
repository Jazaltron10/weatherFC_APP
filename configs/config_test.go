package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCityCountryEndpoint_GetOpenStreetMapLink(t *testing.T) {
	// Arrange
	endpoint := &CityCountryEndpoint{
		City:    "Los Angeles",
		Country: "USA",
		Format:  "json",
	}

	// Act
	link, err := endpoint.GetOpenStreetMapLink()

	// Assert
	assert.NoError(t, err, "GetOpenStreetMapLink should not return an error")
	assert.NotNil(t, link, "GetOpenStreetMapLink should return a valid URL")

	// Additional assertions on the generated URL if needed
	assert.Contains(t, link.String(), "Los Angeles", "URL should contain the city name")
	assert.Contains(t, link.String(), "USA", "URL should contain the country name")
	assert.Contains(t, link.String(), "json", "URL should contain the specified format")
}

func TestForecastCoordinates_GetForecastCoordinatesLink(t *testing.T) {
	// Arrange
	coordinates := &ForecastCoordinates{
		Latitude:  "34.0522",
		Longitude: "-118.2437",
	}

	// Act
	link, err := coordinates.GetForecastCoordinatesLink()

	// Assert
	assert.NoError(t, err, "GetForecastCoordinatesLink should not return an error")
	assert.NotNil(t, link, "GetForecastCoordinatesLink should return a valid URL")

	// Additional assertions on the generated URL if needed
	assert.Contains(t, link.String(), "34.0522", "URL should contain the latitude")
	assert.Contains(t, link.String(), "-118.2437", "URL should contain the longitude")
}

