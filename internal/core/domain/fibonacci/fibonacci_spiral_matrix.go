package fibonacci

type SpiralMatrix struct {
	Spiral                           [][]int
	Top, Bottom, Left, Right, Number int
}

func NewSpiralMatrix(rows, columns int) SpiralMatrix {
	return SpiralMatrix{Spiral: createMatrix[int](rows, columns),
		Top:    0,
		Bottom: rows - 1,
		Left:   0,
		Right:  columns - 1,
		Number: 0}
}

func IncreseTop(sm *SpiralMatrix) {
	sm.Top++
}

func DecreseBottom(sm *SpiralMatrix) {
	sm.Bottom--
}

func IncreseLeft(sm *SpiralMatrix) {
	sm.Left++
}

func DecreseRight(sm *SpiralMatrix) {
	sm.Right--
}

func IncreseNumber(sm *SpiralMatrix) {
	sm.Number++
}

func createMatrix[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		matrix[i] = rows[startRow : startRow+m]
	}
	return matrix
}
