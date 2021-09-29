package weather

import "github.com/google/wire"

// ProviderSet is weather providers.
var ProviderSet = wire.NewSet(NewGetNowWeatherUsecase)
