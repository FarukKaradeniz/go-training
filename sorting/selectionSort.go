package sorting

// SelectionSort is a simple but slow algorithm for large datasets
func SelectionSort(arr *[]int64, size int) {
	for i := 0; i < size-1; i++ {
		min := i
		for j := i + 1; j < size; j++ {
			if (*arr)[min] > (*arr)[j] {
				min = j
			}
		}
		(*arr)[i], (*arr)[min] = (*arr)[min], (*arr)[i]
	}
}
