package main

import (
	"go-training/sorting"
	"go-training/utils"
	"testing"
)

func createRandomArray() []int64 {
	var size = 10
	arr := make([]int64, size)
	for i := 0; i < size; i++ {
		var randNumber = utils.GetRandomNumberBetween(1, 50)
		utils.SetValueAtIndex(&arr, i, randNumber)
	}
	return arr
}

func TestBubbleSort(t *testing.T) {
	array := createRandomArray()

	sorting.BubbleSort(&array)
	if !utils.IsSorted(array) {
		t.Error(array)
	}
}

func TestSelectionSort(t *testing.T) {
	array := createRandomArray()

	sorting.SelectionSort(&array)
	if !utils.IsSorted(array) {
		t.Error(array)
	}
}

func TestMergeSort(t *testing.T) {
	array := createRandomArray()

	sorting.MergeSort(&array)
	if !utils.IsSorted(array) {
		t.Error(array)
	}
}

func TestQuickSort(t *testing.T) {
	array := createRandomArray()

	sorting.QuickSort(&array)
	if !utils.IsSorted(array) {
		t.Error(array)
	}
}
