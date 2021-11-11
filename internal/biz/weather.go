package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Weather struct {
	Hello string
}

type WeatherRepo interface {
	GetNowWeather(ctx context.Context, greeter *Weather) error
}

type GetNowWeatherUsecase struct {
	repo WeatherRepo
	log  *log.Helper
}

func NewGetNowWeatherUsecase(repo WeatherRepo, logger log.Logger) *GetNowWeatherUsecase {
	return &GetNowWeatherUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (s *GetNowWeatherUsecase) GetNow(ctx context.Context) (Weather, error) {
	s.repo.GetNowWeather(ctx, &Weather{})
	return Weather{
		"HeiHei",
	}, nil
}
