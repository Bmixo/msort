package main

import (
	"fmt"
	"msort"
	"time"
)

func main() {
	m := msort.Init()
	tree := m.Tree.IntervalTree.NewIntervalTree()
	node1, err := m.Tree.IntervalTree.NewInterval(time.Unix(10, 0), time.Unix(20, 0))
	if err != nil {
		panic(err)
	}
	node2, err := m.Tree.IntervalTree.NewInterval(time.Unix(15, 0), time.Unix(30, 0))
	if err != nil {
		panic(err)
	}
	node3, err := m.Tree.IntervalTree.NewInterval(time.Unix(21, 0), time.Unix(30, 0))
	if err != nil {
		panic(err)
	}
	tree.Upsert(node1, nil)
	resultList, err := tree.FindAllOverlapping(node2)
	fmt.Println(resultList, err)
	resultList, err = tree.FindAllOverlapping(node3)
	fmt.Println(resultList, err)
}
