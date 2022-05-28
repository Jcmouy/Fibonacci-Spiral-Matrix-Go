//go:build wireinject
// +build wireinject

package wired

import (
	"github.com/google/wire"
	"fibonacci-spiral-matrix-go/internal/api/handler"
	provider3 "fibonacci-spiral-matrix-go/internal/api/handler/provider"
	provider2 "fibonacci-spiral-matrix-go/internal/core/repository/provider"
	"fibonacci-spiral-matrix-go/internal/core/services/provider"
)

var AuthService = wire.NewSet(provider.ProvideAuthService, provider2.ProvideAuthRepository)

func InitializeFiboSpiralMatrixHandler() (handler.FiboSpiralMatrixHandler, error) {
	wire.Build(provider3.ProvideFiboSpiralMatrixHandler, provider.ProvideFiboSpiralMatrixService)
	return handler.FiboSpiralMatrixHandler{}, nil
}

func InitializeAuthHandler() (handler.AuthHandler, error) {
	wire.Build(provider3.ProvideAuthHandler, AuthService)
	return handler.AuthHandler{}, nil
}

func InitializeHealthHandler() handler.HealthHandler {
	return handler.HealthHandler{}
}
