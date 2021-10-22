package data

import (
	"context"

	"weather-router/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	g.Hello = "hehehe"
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

type weatherRepo struct {
	heWeather *HeWeather
	log       *log.Helper
}

func (w *weatherRepo) GetNowWeather(ctx context.Context, weather *biz.Weather) error {
	return nil
}

func NewWeatherRepo(weather *HeWeather, logger log.Logger) biz.NowWeatherRepo {
	return &weatherRepo{
		heWeather: weather,
		log:       log.NewHelper(logger),
	}
}
