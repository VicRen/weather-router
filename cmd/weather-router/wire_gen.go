// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"weather-router/internal/biz"
	"weather-router/internal/conf"
	"weather-router/internal/data"
	"weather-router/internal/server"
	"weather-router/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase, logger)
	heWeather, cleanup2, err := data.NewHeWeather(logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	nowWeatherRepo := data.NewWeatherRepo(heWeather, logger)
	getNowWeatherUsecase := biz.NewGetNowWeatherUsecase(nowWeatherRepo, logger)
	weatherService := service.NewWeatherService(getNowWeatherUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, weatherService, logger)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
