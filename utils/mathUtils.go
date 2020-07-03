package utils

import (
	"crypto/rand"
	"math/big"
)

// GetRandomNumberBetween is a function that returns a random number between two numbers
func GetRandomNumberBetween(start int64, stop int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(stop))
	//return rand.Intn(stop-start) + start
	return n.Int64() + start
}
