// internal/forecast/forecast.go
package forecast

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jazaltron10/Golang/weatherFC_APP/configs"
)

// FetchForecast fetches the weather forecast for a given set of coordinates.
func FetchForecast(link *url.URL, client *http.Client) ([]configs.ForecastPeriod, error) {
	// Fetch data using the HTTP client
	response, err := client.Get(link.String())
	if err != nil {
		return nil, fmt.Errorf("error fetching forecast data from URL %s: %v", link.String(), err)
	}
	defer response.Body.Close()

	// Read the response body
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading forecast data from response body: %v", err)
	}

	// Parse the forecast data
	return ParseForecast(data)
}

// ParseForecast parses the raw forecast data and returns the relevant forecast periods.
func ParseForecast(data []byte) ([]configs.ForecastPeriod, error) {
	var propertiesInfo configs.PropertiesForecastInfo

	// Unmarshal JSON data into the propertiesInfo struct
	if err := json.Unmarshal(data, &propertiesInfo); err != nil {
		return nil, fmt.Errorf("error unmarshalling forecast data: %v", err)
	}

	// Extract the forecast periods
	return propertiesInfo.Periods.Periods, nil
}
