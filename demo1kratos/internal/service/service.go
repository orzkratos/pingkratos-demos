package service

import (
	"github.com/google/wire"
	"github.com/orzkratos/pingkratos/serverpingkratos"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewGreeterService,
	serverpingkratos.NewPingService,
)
