package service

import (
	"context"
	"weather-router/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// NowWeatherService is a weather service.
type NowWeatherService struct {
	uc  *biz.GetNowWeatherUsecase
	log *log.Helper
}

// NewNowWeatherService returns a new weather service.
func NewNowWeatherService(uc *biz.GetNowWeatherUsecase, logger log.Logger) *NowWeatherService {
	return &NowWeatherService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *NowWeatherService) GetNow(ctx context.Context, in string) string {
	return ""
}
