package util

var (
	ListNumbers []int
	a           int
	b           = 1
	c           int
)

func GetFibonacciSequence(column, row int) []int {
	ResetValuesIfNeeded()
	ListNumbers = append(ListNumbers, 0)
	ListNumbers = append(ListNumbers, 1)
	CalculateFibonacciNumber(column*row - 2)
	return ListNumbers
}

func CalculateFibonacciNumber(n int) {
	if n > 0 {
		c = a + b
		a = b
		b = c
		ListNumbers = append(ListNumbers, c)
		CalculateFibonacciNumber(n - 1)
	}
}

func ResetValuesIfNeeded() {
	if a > 0 {
		a = 0
		b = 1
		ListNumbers = nil
	}
}
