package cache

import (
	"sync"

	"github.com/jazaltron10/Golang/weatherFC_APP/configs"
)

type Cache interface {
	Get(string) ([]configs.ForecastPeriod, error)
	Set(string, []configs.ForecastPeriod) error
}

// InternalCache is an in-memory cache database.
type InternalCache struct {
	data map[string][]configs.ForecastPeriod
	mu   sync.RWMutex
}

// NewInternalCache creates a new instance of InternalCache.
func NewInternalCache() *InternalCache {
	return &InternalCache{
		data: make(map[string][]configs.ForecastPeriod),
	}
}

// Get retrieves forecast data from the cache based on the key.
func (c *InternalCache) Get(key string) ([]configs.ForecastPeriod, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	data, exists := c.data[key]
	if !exists {
		return nil, nil
	}

	return data, nil
}

// Set stores forecast data in the cache based on the key.
func (c *InternalCache) Set(key string, forecastData []configs.ForecastPeriod) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = forecastData

	return nil
}
