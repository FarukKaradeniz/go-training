package sorting

func QuickSort(arr *[]int64) {
	qSort(arr, 0, len(*arr)-1)
}

func qSort(arr *[]int64, left int, right int) {
	if left < right {
		p := partition(arr, left, right)
		qSort(arr, left, p-1)
		qSort(arr, p+1, right)
	}
}

func partition(arr *[]int64, left int, right int) int {
	pivot := (*arr)[right] // Select last element as pivot
	i := left - 1          // index of smaller element

	for j := left; j < right; j++ {
		if (*arr)[j] <= pivot {
			i++
			(*arr)[j], (*arr)[i] = (*arr)[i], (*arr)[j]
		}
	}

	(*arr)[i+1], (*arr)[right] = (*arr)[right], (*arr)[i+1]

	return i + 1
}
