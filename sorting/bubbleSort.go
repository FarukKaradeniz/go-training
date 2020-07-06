package sorting

// BubbleSort is a simple but slow algorithm for large datasets
func BubbleSort(arr *[]int64, size int) {
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
