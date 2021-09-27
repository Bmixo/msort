package msort

import (
	"msort/normal"
	trees "msort/tree"
)

var Sore *MMSort

type MMSort struct {
	Normal *normal.MNormal
	Tree   *trees.MTree
}

func Init() *MMSort {
	Sore = &MMSort{
		Normal: normal.Init(),
		Tree:   trees.Init(),
	}
	return Sore
}
