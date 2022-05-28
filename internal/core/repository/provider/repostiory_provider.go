package provider

import (
	"fibonacci-spiral-matrix-go/internal/core/interfaces/repository"
	repository2 "fibonacci-spiral-matrix-go/internal/core/repository"
)

func ProvideAuthRepository() repository.AuthRepository {
	return repository2.NewAuthRepository()
}
