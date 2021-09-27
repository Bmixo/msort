package normal

import (
	"msort/normal/bubble_sort"
	"msort/normal/insert_sort"
	"msort/normal/merge_sort"
	"msort/normal/quick_sort"
	"msort/normal/select_sort"
)

var Normal *MNormal

type MNormal struct {
	BubbleSort *bubble_sort.MBubbleSort
	InsertSort *insert_sort.MInsertSort
	MergeSort  *merge_sort.MMergeSort
	QuickSort  *quick_sort.MQuickSort
	SelectSort *select_sort.MSelectSort
}

func Init() *MNormal {
	Normal = &MNormal{
		BubbleSort: bubble_sort.Init(),
		InsertSort: insert_sort.Init(),
		MergeSort:  merge_sort.Init(),
		QuickSort:  quick_sort.Init(),
		SelectSort: select_sort.Init(),
	}
	return Normal
}
