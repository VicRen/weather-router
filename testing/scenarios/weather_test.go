package scenarios

import (
	"testing"
	"weather-router/internal/data"
)

func TestWeatherGetNow(t *testing.T) {
	_, cleanup, err := data.NewHeWeather(&dummyLogger{})
	defer cleanup()
	if err != nil {
		return
	}
}
