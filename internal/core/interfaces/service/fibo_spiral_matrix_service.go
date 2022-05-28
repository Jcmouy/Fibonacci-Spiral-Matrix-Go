package service

import (
	"context"
	"fibonacci-spiral-matrix-go/internal/core/domain/fibonacci"
)

type FiboSpiralMatrixService interface {
	GetSpiralMatrix(ctx context.Context, matrix fibonacci.Matrix) (fibonacci.SpiralMatrix, error)
}
