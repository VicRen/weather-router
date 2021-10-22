package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Weather struct {
	Hello string
}

type NowWeatherRepo interface {
	GetNowWeather(ctx context.Context, greeter *Weather) error
}

type GetNowWeatherUsecase struct {
	repo NowWeatherRepo
	log  *log.Helper
}

func NewGetNowWeatherUsecase(repo NowWeatherRepo, logger log.Logger) *GetNowWeatherUsecase {
	return &GetNowWeatherUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (s *GetNowWeatherUsecase) GetNow() (Weather, error) {
	return Weather{
		"HeiHei",
	}, nil
}
