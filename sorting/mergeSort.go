package sorting

func MergeSort(arr *[]int64) {
	sort(arr, 0, len(*arr)-1)
}

func sort(arr *[]int64, left int, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2
	sort(arr, left, mid)
	sort(arr, mid+1, right)

	merge(arr, left, mid, right)
}

func merge(arr *[]int64, left int, mid int, right int) {
	tmp := make([]int64, right-left+1)

	i := left
	j := mid + 1
	k := 0

	for i <= mid && j <= right {
		if (*arr)[i] < (*arr)[j] {
			tmp[k] = (*arr)[i]
			i++
		} else {
			tmp[k] = (*arr)[j]
			j++
		}
		k++
	}

	for i <= mid {
		tmp[k] = (*arr)[i]
		k++
		i++
	}

	for j <= right {
		tmp[k] = (*arr)[j]
		k++
		j++
	}

	for idx := left; idx <= right; idx++ {
		(*arr)[idx] = tmp[idx-left]
	}
}
