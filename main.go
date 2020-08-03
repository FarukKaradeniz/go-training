package main

import (
	"fmt"

	"go-training/sorting"
	"go-training/utils"
)

func main() {
	fmt.Println("Hello, World")
	var size int                 // size of the array
	var isRandomGenerated string // 'y' or 'n'. if 'y' populate array with random numbers
	var arr []int64
	var sortAlgorithm int

	fmt.Print("Please Enter the size of the array: ")
	_, errSize := fmt.Scan(&size)
	if nil != errSize {
		panic("invalid input")
	}

	fmt.Print("Please type 'y' if you want to populate array with random numbers, type 'n' if you want to assign all the values: ")
	_, errGen := fmt.Scan(&isRandomGenerated)
	if nil != errGen {
		panic("invalid input")
	}

	// TODO Ask repeatedly to user if inputs are invalid
	arr = utils.CreateArray(size)

	if isRandomGenerated == "y" || isRandomGenerated == "Y" {
		setRandom(&arr, size)
	} else if isRandomGenerated == "n" || isRandomGenerated == "N" {
		setManual(&arr, size)
	} else {
		panic("You had to type 'N' or 'Y'")
	}

	fmt.Println("1- BubbleSort\n2- SelectionSort")
	fmt.Print("Please select the sorting algoritm: ")
	_, errAlgorithmPick := fmt.Scan(&sortAlgorithm)
	if nil != errAlgorithmPick {
		panic("invalid input")
	}

	fmt.Println("Not Sorted: ", arr)
	switch sortAlgorithm {
	case 1:
		sorting.BubbleSort(&arr, size)
	case 2:
		sorting.SelectionSort(&arr, size)
	}
	fmt.Println("Sorted:     ", arr)
}

func setRandom(arr *[]int64, size int) {
	for i := 0; i < size; i++ {
		var randNumber = utils.GetRandomNumberBetween(1, 50)
		utils.SetValueAtIndex(arr, i, randNumber)
	}
}

func setManual(arr *[]int64, size int) {
	var value int64
	for i := 0; i < size; i++ {
		fmt.Print("Please Enter the ", i+1, ". value: ")
		_, errValue := fmt.Scan(&value)
		if nil != errValue {
			panic("invalid error")
		}

		utils.SetValueAtIndex(arr, i, value)
	}
}
