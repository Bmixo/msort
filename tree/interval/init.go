package interval

var IntervalTree *MIntervalTree

type MIntervalTree struct {
}

func Init() *MIntervalTree {
	IntervalTree = &MIntervalTree{}
	return IntervalTree
}
