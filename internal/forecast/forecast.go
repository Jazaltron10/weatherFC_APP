// internal/forecast/forecast.go
package forecast

import (
	// "encoding/json"
	// "fmt"
	// "io"
	"net/http"
	// "net/url"
	"os"
	"time"

	// "github.com/jazaltron10/Golang/weatherFC_APP/configs"
	"github.com/jazaltron10/Golang/weatherFC_APP/internal/cache"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.InfoLevel)
}

type ForecastCoordinates struct {
	Latitude  string
	Longitude string
}

type ForecastPeriod struct {
	DetailedForecast string    `json:"detailedForecast"`
	StartTime        time.Time `json:"startTime"`
	EndTime          time.Time `json:"endTime"`
}

type PropertiesForecastInfo struct {
	Periods []ForecastPeriod `json:"periods"`
}

type Handler struct {
	c     *http.Client
	store cache.Cache
	l     *logrus.Logger
}

func NewHandler(client *http.Client, store cache.Cache, logger *logrus.Logger) *Handler {
	return &Handler{
		c:     client,
		store: store,
		l:     logger,
	}
}

// func (h *Handler) GetCoordinates(city string) (*url.URL, error) {
// 	endpoint := &configs.CityCountryEndpoint{
// 		Country: "USA",
// 		City:    city,
// 		Format:  "json", // Assuming JSON format for coordinates
// 	}

// 	link, err := endpoint.GetOpenStreetMapLink()
// 	if err != nil {
// 		return nil, fmt.Errorf("error getting coordinates for city %s: %v", city, err)
// 	}

// 	return link, nil
// }

// func (h *Handler) GetWeatherForecast(city string) (PropertiesForecastInfo, error) {
// 	coordinatesLink, err := h.GetCoordinates(city)
// 	if err != nil {
// 		return PropertiesForecastInfo{}, fmt.Errorf("error getting coordinates: %v", err)
// 	}

// 	forecastData, err := h.fetchData(coordinatesLink)
// 	if err != nil {
// 		return PropertiesForecastInfo{}, fmt.Errorf("error fetching forecast data: %v", err)
// 	}

// 	return forecastData, nil
// }

// func (h *Handler) fetchData(link *url.URL) (PropertiesForecastInfo, error) {
// 	response, err := h.c.Get(link.String())
// 	if err != nil {
// 		return PropertiesForecastInfo{}, fmt.Errorf("error fetching data from URL %s: %v", link.String(), err)
// 	}
// 	defer response.Body.Close()

// 	data, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		return PropertiesForecastInfo{}, fmt.Errorf("error reading data from response body: %v", err)
// 	}

// 	var forecastData PropertiesForecastInfo
// 	err = json.Unmarshal(data, &forecastData)
// 	if err != nil {
// 		return PropertiesForecastInfo{}, fmt.Errorf("error unmarshalling forecast data: %v", err)
// 	}

// 	return forecastData, nil
// }
