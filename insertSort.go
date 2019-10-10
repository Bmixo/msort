package msort

// InsertSort : sort
func InsertSort(array []int, bigMod bool) {
	length := len(array)
	if length <= 1 {
		return
	}
	if bigMod {
		for i := 1; i < length; i++ {
			if array[i] > array[i-1] {
				tmp := array[i]
				j := i
				for j >= 1 && array[j-1] < tmp {
					array[j] = array[j-1]
					j--
				}
				array[j] = tmp
			}
		}
		return
	}
	for i := 1; i < length; i++ {

		if array[i] < array[i-1] {
			tmp := array[i]
			j := i
			for j >= 1 && array[j-1] > tmp {
				array[j] = array[j-1]
				j--
			}
			array[j] = tmp

		}
	}
}
