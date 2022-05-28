package dto

import (
	"fibonacci-spiral-matrix-go/internal/core/domain/fibonacci"
	"fibonacci-spiral-matrix-go/internal/core/util"
)

type FibonacciSpiralMatrixDto struct {
	Rows [][]int `json:"rows"`
}

func (fibDto *FibonacciSpiralMatrixDto) FromModel(spiralMatrix fibonacci.SpiralMatrix) FibonacciSpiralMatrixDto {
	fibDto.Rows = spiralMatrix.Spiral
	return *fibDto
}

type MatrixInput struct {
	Row    string `json:"row" binding:"required"`
	Column string `json:"column" binding:"required"`
}

func (dto *MatrixInput) ToModel() fibonacci.Matrix {
	intRow := util.ValueToInt(dto.Row)
	intColumn := util.ValueToInt(dto.Column)

	return fibonacci.Matrix{
		Row:    intRow,
		Column: intColumn,
	}
}