package trees

import "msort/tree/interval"

var Tree *MTree

type MTree struct {
	IntervalTree *interval.MIntervalTree
}

func Init() *MTree {
	Tree = &MTree{
		IntervalTree: interval.Init(),
	}
	return Tree
}
