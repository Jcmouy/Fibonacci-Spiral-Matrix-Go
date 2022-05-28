package provider

import (
	"fibonacci-spiral-matrix-go/internal/api/handler"
	"fibonacci-spiral-matrix-go/internal/core/interfaces/service"
)

func ProvideFiboSpiralMatrixHandler(fiboSpiralMatrixService service.FiboSpiralMatrixService) handler.FiboSpiralMatrixHandler {
	return handler.NewFiboSpiralMatrixHandler(fiboSpiralMatrixService)
}

func ProvideAuthHandler(authService service.AuthService) handler.AuthHandler {
	return handler.NewAuthHandler(authService)
}
