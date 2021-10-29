package data

import (
	"context"
	"weather-router/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type weatherRepo struct {
	heWeather *HeWeather
	log       *log.Helper
}

func (w *weatherRepo) GetNowWeather(ctx context.Context, weather *biz.Weather) error {
	return nil
}

func NewWeatherRepo(weather *HeWeather, logger log.Logger) biz.WeatherRepo {
	return &weatherRepo{
		heWeather: weather,
		log:       log.NewHelper(logger),
	}
}
