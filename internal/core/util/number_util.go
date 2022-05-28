package util

import (
	"log"
	"strconv"
)

func ValueToInt(val string) int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	return intVal
}
