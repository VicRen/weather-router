package service

import (
	"context"

	pb "weather-router/api/weather"
)

type WeatherService struct {
	pb.UnimplementedWeatherServer
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) CreateWeather(ctx context.Context, req *pb.CreateWeatherRequest) (*pb.CreateWeatherReply, error) {
	return &pb.CreateWeatherReply{}, nil
}
func (s *WeatherService) UpdateWeather(ctx context.Context, req *pb.UpdateWeatherRequest) (*pb.UpdateWeatherReply, error) {
	return &pb.UpdateWeatherReply{}, nil
}
func (s *WeatherService) DeleteWeather(ctx context.Context, req *pb.DeleteWeatherRequest) (*pb.DeleteWeatherReply, error) {
	return &pb.DeleteWeatherReply{}, nil
}
func (s *WeatherService) GetWeather(ctx context.Context, req *pb.GetWeatherRequest) (*pb.GetWeatherReply, error) {
	return &pb.GetWeatherReply{}, nil
}
func (s *WeatherService) ListWeather(ctx context.Context, req *pb.ListWeatherRequest) (*pb.ListWeatherReply, error) {
	return &pb.ListWeatherReply{}, nil
}
