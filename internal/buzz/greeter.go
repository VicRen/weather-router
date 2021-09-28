package buzz

import "github.com/google/wire"

// ProviderSet is buzz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase)
