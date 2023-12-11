package configs

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	openStreetMapWebLink      = "https://nominatim.openstreetmap.org"
	forecastPairOfCoordinates = "https://api.weather.gov/points"
)

var openStreetFormats = []string{"xml", "geojson", "geocodejson", "json", "jsonv2"}

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.InfoLevel)
}

type CityCountryEndpoint struct {
	City    string
	Country string
	Format  string
}

type ForecastCoordinates struct {
	Longitude string `json:"lon"`
	Latitude  string `json:"lat"`
}

type PropertiesInfo struct {
	Properties PropertyInfo `json:"properties"`
}

type PropertyInfo struct {
	ForecastURL string `json:"forecast"`
}

type ForecastPeriod struct {
	DetailedForecast string    `json:"detailedForecast"`
	StartTime        time.Time `json:"startTime"`
	EndTime          time.Time `json:"endTime"`
}

type PropertiesForecastInfo struct {
	Periods ForecastPeriodsInfo `json:"properties"`
}

type ForecastPeriodsInfo struct {
	Periods []ForecastPeriod `json:"periods"`
}

func (c *CityCountryEndpoint) GetOpenStreetMapLink() (*url.URL, error) {
	if !c.formatIsValid() {
		return nil, errors.New("invalid format")
	}

	link := fmt.Sprintf("%s/search?q=%s,%s&format=%s", openStreetMapWebLink, c.City, c.Country, c.Format)
	return getURL(link)
}

func (c *CityCountryEndpoint) formatIsValid() bool {
	for _, format := range openStreetFormats {
		if c.Format == format {
			return true
		}
	}
	return false
}

func (f *ForecastCoordinates) GetForecastCoordinatesLink() (*url.URL, error) {
	link := fmt.Sprintf("%s/%s,%s", forecastPairOfCoordinates, f.Latitude, f.Longitude)
	return getURL(link)
}

func getURL(link string) (*url.URL, error) {
	return url.Parse(link)
}

// In this completion, the `GetOpenStreetMapLink` and `GetForecastCoordinatesLink` functions construct the URLs for the OpenStreetMap and weather forecast services, respectively. The `getURL` function is a utility function to parse a string into a URL.
