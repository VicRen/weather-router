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
	heWeather, cleanup, err := data.NewHeWeather(logger)
	if err != nil {
		return nil, nil, err
	}
	weatherRepo := data.NewWeatherRepo(heWeather, logger)
	getNowWeatherUsecase := biz.NewGetNowWeatherUsecase(weatherRepo, logger)
	weatherService := service.NewWeatherService(getNowWeatherUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, weatherService, logger)
	grpcServer := server.NewGRPCServer(confServer, weatherService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
