package data

import (
	"context"
	"fmt"
	"net/http"
	"weather-router/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type weatherRepo struct {
	heWeather *HeWeather
	log       *log.Helper
}

func (w *weatherRepo) GetNowWeather(ctx context.Context, weather *biz.Weather) error {
	ub := "https://devapi.qweather.com/v7/weather/now?"
	body, err := http.Get(ub)
	if err != nil {
		return err
	}
	fmt.Printf("body: %v\n", body.Body)
	return nil
}

func NewWeatherRepo(weather *HeWeather, logger log.Logger) biz.WeatherRepo {
	return &weatherRepo{
		heWeather: weather,
		log:       log.NewHelper(logger),
	}
}

func getSignature(params map[string]string, secret string) string {
	return ""
}
