package utils

import (
	"math/rand"
)

// GetRandomNumberBetween is a function that returns a random number between two numbers
func GetRandomNumberBetween(start int, stop int) int {
	return rand.Intn(stop-start) + start
}
