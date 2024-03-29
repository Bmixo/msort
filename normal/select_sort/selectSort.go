package select_sort

type MSelectSort struct {
}

var SelectSort *MSelectSort

func Init() *MSelectSort {
	SelectSort = &MSelectSort{}
	return SelectSort
}

// SelectSort : sort number from array
func (m *MSelectSort) SelectSort(array []int, bigMod bool) {
	length := len(array)
	if len(array) <= 1 {
		return
	}
	if bigMod {
		for i := 0; i < length; i++ {
			tmp := array[i]
			flag := i
			for j := i + 1; j < length; j++ {
				if array[j] > tmp {
					flag = j
					tmp = array[j]
				}
			}
			if flag != i {
				array[flag] = array[i]
				array[i] = tmp
			}
		}
		return
	}

	for i := 0; i < length; i++ {
		tmp := array[i]
		flag := i
		for j := i + 1; j < length; j++ {
			if array[j] < tmp {
				flag = j
				tmp = array[j]
			}
		}
		if flag != i {
			array[flag] = array[i]
			array[i] = tmp
		}
	}

}
