package msort

import (
	"github.com/Bmixo/msort/normal"
	trees "github.com/Bmixo/msort/tree"
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
