package scenarios

import (
	"testing"
	"weather-router/internal/data"

	"github.com/go-kratos/kratos/v2/log"
)

type dummyLogger struct {
}

func (d *dummyLogger) Log(level log.Level, keyvals ...interface{}) error {
	return nil
}

func TestWeatherGetNow(t *testing.T) {
	_, cleanup, err := data.NewHeWeather(nil)
	defer cleanup()
	if err != nil {
		return
	}
}
