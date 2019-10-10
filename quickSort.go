package msort

func sortS(array []int, left, right int) {
	if left >= right {
		return
	}
	low, high := left, right
	key := array[low]

	for low < high {
		for low < high && array[high] >= key {
			high--
		}
		if low < high {
			array[low] = array[high]
		}
		for low < high && array[low] <= key {
			low++
		}
		if low < high {
			array[high] = array[low]
		}
	}
	array[low] = key
	sortS(array, left, low)
	sortS(array, low+1, right)

}

func sortB(array []int, left, right int) {

	if left >= right {
		return
	}
	low, high := left, right
	key := array[low]
	for low < high {
		for low < high && array[high] <= key {
			high--
		}
		if low < high {
			array[low] = array[high]
		}
		for low < high && array[low] > key {
			low++
		}
		if low < high {
			array[high] = array[low]
		}

	}
	array[low] = key
	sortB(array, left, low)
	sortB(array, low+1, right)

}
func QuickSort(array []int, bigMod bool) {
	if bigMod {
		sortB(array, 0, len(array)-1)
		return
	}
	sortS(array, 0, len(array)-1)
}
