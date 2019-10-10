package msort

// BubbleSort : sort nunmber from array
func BubbleSort(array []int, bigMod bool) {
	length := len(array)
	if bigMod {
		for i := 0; i < length; i++ {
			for j := 1; j < length; j++ {
				if array[j] > array[j-1] {
					tmp := array[j-1]
					array[j-1] = array[j]
					array[j] = tmp
				}
			}
		}
	} else {
		for i := 0; i < length; i++ {
			for j := 1; j < length; j++ {
				if array[j] < array[j-1] {
					tmp := array[j-1]
					array[j-1] = array[j]
					array[j] = tmp

				}

			}
		}
	}
}
