package merge_sort

type MMergeSort struct {
}

var MergeSort *MMergeSort

func Init() *MMergeSort {
	MergeSort = &MMergeSort{}
	return MergeSort
}

//Merge : sort
func mergeB(array []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	left := make([]int, n1)
	right := make([]int, n2)
	for i, k := 0, p; i < n1; i, k = i+1, k+1 {
		left[i] = array[k]
	}
	for i, k := 0, q+1; i < n2; i, k = i+1, k+1 {
		right[i] = array[k]
	}
	var i, k, j int
	for i, k, j = 0, p, 0; i < n1 && j < n2; k++ {
		if left[i] > right[j] {
			array[k] = left[i]
			i++
		} else {
			array[k] = right[j]
			j++
		}
	}
	if i < n1 {
		for j = i; j < n1; j, k = j+1, k+1 {
			array[k] = left[i]
		}
	}
	if j < n2 {
		for i = j; i < n2; i, k = i+1, k+1 {
			array[k] = right[i]
		}
	}
}

//Merge : sort
func mergeS(array []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	left := make([]int, n1)
	right := make([]int, n2)
	for i, k := 0, p; i < n1; i, k = i+1, k+1 {
		left[i] = array[k]
	}
	for i, k := 0, q+1; i < n2; i, k = i+1, k+1 {
		right[i] = array[k]
	}
	var i, k, j int
	for i, k, j = 0, p, 0; i < n1 && j < n2; k++ {
		if left[i] < right[j] {
			array[k] = left[i]
			i++
		} else {
			array[k] = right[j]
			j++
		}
	}
	if i < n1 {
		for j = i; j < n1; j, k = j+1, k+1 {
			array[k] = left[i]
		}
	}
	if j < n2 {
		for i = j; i < n2; i, k = i+1, k+1 {
			array[k] = right[i]
		}
	}
}

// MergeSort : sort
func (m *MMergeSort) MergeSort(array []int, p int, r int, bigMode bool) {

	if p < r {

		q := (p + r) / 2
		m.MergeSort(array, p, q, bigMode)
		m.MergeSort(array, q+1, r, bigMode)
		if bigMode {
			mergeB(array, p, q, r)
		} else {
			mergeS(array, p, q, r)
		}
	}

}
