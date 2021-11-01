package data

import (
	"context"
	"weather-router/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type mockWeatherRepo struct {
	heWeather *HeWeather
	log       *log.Helper
}

func (w *mockWeatherRepo) GetNowWeather(ctx context.Context, weather *biz.Weather) error {
	weather.Hello = "Sunny"
	return nil
}

func NewMockWeatherRepo(weather *HeWeather, logger log.Logger) biz.WeatherRepo {
	return &weatherRepo{
		heWeather: weather,
		log:       log.NewHelper(logger),
	}
}
