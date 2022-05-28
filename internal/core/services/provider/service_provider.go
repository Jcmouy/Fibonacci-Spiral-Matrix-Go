package provider

import (
	"fibonacci-spiral-matrix-go/internal/core/interfaces/repository"
	"fibonacci-spiral-matrix-go/internal/core/interfaces/service"
	"fibonacci-spiral-matrix-go/internal/core/services"
)

func ProvideFiboSpiralMatrixService() service.FiboSpiralMatrixService {
	return services.NewFiboSpiralMatrixService()
}

func ProvideAuthService(authRepository repository.AuthRepository) service.AuthService {
	return services.NewAuthServiceService(authRepository)
}
