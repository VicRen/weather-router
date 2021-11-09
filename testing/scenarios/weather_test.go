package scenarios

import (
	"net"
	"testing"
	"weather-router/internal/biz"
	"weather-router/internal/data"
	"weather-router/internal/server"
	"weather-router/internal/service"
)

func TestWeatherGetNow(t *testing.T) {
	logger := &dummyLogger{}
	heweather, cleanup, err := data.NewHeWeather(logger)
	defer cleanup()
	if err != nil {
		return
	}
	weatherRepo := data.NewWeatherRepo(heweather, logger)
	getNowWeatherUsecase := biz.NewGetNowWeatherUsecase(weatherRepo, logger)
	weatherService := service.NewWeatherService(getNowWeatherUsecase, logger)
	httpServer := server.NewHTTPServer(nil, weatherService, logger)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	httpServer.Serve(l)
}
