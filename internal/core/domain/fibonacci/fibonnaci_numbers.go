package fibonacci

import "fibonacci-spiral-matrix-go/internal/core/util"

type Numbers struct {
	numbers []int
}

func MapFibonacciSequence(column, row int, fibNum *Numbers) {
	fibNum.numbers = util.GetFibonacciSequence(column, row)
}

func GetElementAtPosition(position int, fibNum Numbers) int {
	return find(position, fibNum.numbers)
}

func find(searchPosition int, array []int) (idx int) {
	for position, element := range array {
		if position == searchPosition {
			return element
		}
	}
	return -1
}
