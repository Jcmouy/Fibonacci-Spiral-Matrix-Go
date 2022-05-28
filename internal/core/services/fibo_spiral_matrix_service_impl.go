package services

import (
	"context"
	"fibonacci-spiral-matrix-go/internal/core/domain/fibonacci"
	"github.com/opentracing/opentracing-go"
)

type FiboSpiralMatrixServiceImpl struct {
}

var fibNum = fibonacci.Numbers{}
var fibSpM = fibonacci.SpiralMatrix{}

func NewFiboSpiralMatrixService() *FiboSpiralMatrixServiceImpl {
	return &FiboSpiralMatrixServiceImpl{}
}

func (fsm *FiboSpiralMatrixServiceImpl) GetSpiralMatrix(ctx context.Context, matrix fibonacci.Matrix) (fibonacci.SpiralMatrix, error) {
	span, spanContext := opentracing.StartSpanFromContext(ctx, "GetSpiralMatrix")
	defer span.Finish()

	InitFibonacciNumbers(spanContext, matrix)
	SpiralMatrix(spanContext, matrix)
	return fibSpM, nil
}

func InitFibonacciNumbers(ctx context.Context, matrix fibonacci.Matrix) {
	span, _ := opentracing.StartSpanFromContext(ctx, "InitFibonacciNumbers")
	defer span.Finish()

	fibonacci.MapFibonacciSequence(matrix.Row, matrix.Column, &fibNum)
}

func SpiralMatrix(ctx context.Context, matrix fibonacci.Matrix) {
	span, _ := opentracing.StartSpanFromContext(ctx, "SpiralMatrix")
	defer span.Finish()

	InitFibonacciSpiral(matrix)
	for outerCondition() {
		getValuesFromTopRow()
		getValuesFromRightColumn()
		getValuesFromBottomRow()
		getValuesFromLeftColumn()
	}
}

func InitFibonacciSpiral(matrix fibonacci.Matrix) {
	fibSpM = fibonacci.NewSpiralMatrix(matrix.Row, matrix.Column)
}

func outerCondition() bool {
	return fibSpM.Left <= fibSpM.Right && fibSpM.Top <= fibSpM.Bottom
}

func getValuesFromTopRow() {
	for i := fibSpM.Left; i <= fibSpM.Right; i++ {
		fibSpM.Spiral[fibSpM.Top][i] = fibonacci.GetElementAtPosition(fibSpM.Number, fibNum)
		increaseValueNumber()
	}
	fibonacci.IncreseTop(&fibSpM)
}

func getValuesFromRightColumn() {
	for i := fibSpM.Top; i <= fibSpM.Bottom; i++ {
		fibSpM.Spiral[i][fibSpM.Right] = fibonacci.GetElementAtPosition(fibSpM.Number, fibNum)
		increaseValueNumber()
	}
	fibonacci.DecreseRight(&fibSpM)
}

func getValuesFromBottomRow() {
	if fibSpM.Bottom >= fibSpM.Top {
		for i := fibSpM.Right; i >= fibSpM.Left; i-- {
			fibSpM.Spiral[fibSpM.Bottom][i] = fibonacci.GetElementAtPosition(fibSpM.Number, fibNum)
			increaseValueNumber()
		}
		fibonacci.DecreseBottom(&fibSpM)
	}
}

func getValuesFromLeftColumn() {
	for i := fibSpM.Bottom; i >= fibSpM.Top; i-- {
		fibSpM.Spiral[i][fibSpM.Left] = fibonacci.GetElementAtPosition(fibSpM.Number, fibNum)
		increaseValueNumber()
	}
	fibonacci.IncreseLeft(&fibSpM)
}

func increaseValueNumber() {
	fibonacci.IncreseNumber(&fibSpM)
}
