package utils

import (
	"fmt"
)

// CreateArray is a function that returns an array
func CreateArray(size int) []int {
	return make([]int, size)
}

// SetValueAtIndex is a function that sets a value of an array at spesific index
func SetValueAtIndex(arr *[]int, index int, value int) {
	if len(*arr) < index {
		fmt.Println("Index shouldn't be bigger than array's length")
	}
	(*arr)[index] = value
}
