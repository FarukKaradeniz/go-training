package main

import (
	"fmt"

	"./utils"
)

func main() {
	fmt.Println("Hello, World")
	var size int                 // size of the array
	var isRandomGenerated string // 'y' or 'n'. if 'y' populate array with random numbers
	var arr []int64

	fmt.Print("Please Enter the size of the array: ")
	fmt.Scan(&size)
	//	_, errSize := fmt.Scan(&size)

	fmt.Print("Please type 'y' if you want to populate array with random numbers, type 'n' if you want to assign all the values: ")
	fmt.Scan(&isRandomGenerated)

	// TODO Ask repeatedly to user if inputs are invalid
	arr = utils.CreateArray(size)

	if isRandomGenerated == "y" || isRandomGenerated == "Y" {
		setRandom(&arr, size)
	} else if isRandomGenerated == "n" || isRandomGenerated == "N" {
		setManuel(&arr, size)
	} else {
		fmt.Println("You had to type 'N' or 'Y'")
	}

	fmt.Println(arr)
}

func setRandom(arr *[]int64, size int) {
	for i := 0; i < size; i++ {
		var randNumber = utils.GetRandomNumberBetween(1, 50)
		utils.SetValueAtIndex(arr, i, randNumber)
	}
}

func setManuel(arr *[]int64, size int) {
	var value int64
	for i := 0; i < size; i++ {
		fmt.Print("Please Enter the ", (i + 1), ". value: ")
		fmt.Scan(&value)
		utils.SetValueAtIndex(arr, i, value)
	}
}
