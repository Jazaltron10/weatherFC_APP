package cache

import "github.com/jazaltron10/Golang/weatherFC_APP/configs"

// MockCache is a mock implementation of the Cache interface for testing.
type MockCache struct {
	GetFunc func(key string) ([]configs.ForecastPeriod, error)
	SetFunc func(key string, forecastData []configs.ForecastPeriod) error
}

// Get calls the GetFunc of the MockCache.
func (m *MockCache) Get(key string) ([]configs.ForecastPeriod, error) {
	return m.GetFunc(key)
}

// Set calls the SetFunc of the MockCache.
func (m *MockCache) Set(key string, forecastData []configs.ForecastPeriod) error {
	return m.SetFunc(key, forecastData)
}
