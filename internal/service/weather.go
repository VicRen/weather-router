package service

import (
	"context"
	"weather-router/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

	pb "weather-router/api/weather/v1"
)

type WeatherService struct {
	pb.UnimplementedWeatherServer

	uc  *biz.GetNowWeatherUsecase
	log *log.Helper
}

func NewWeatherService(uc *biz.GetNowWeatherUsecase, logger log.Logger) *WeatherService {
	return &WeatherService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *WeatherService) GetNowWeather(ctx context.Context, req *pb.GetNowWeatherRequest) (*pb.GetNowWeatherResponse, error) {
	return &pb.GetNowWeatherResponse{}, nil
}
