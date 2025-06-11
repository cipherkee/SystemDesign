package main

import (
	"github.com/wangjia184/sortedset"

	"fmt"
)

func PlayAround() {
	set := sortedset.New()
	//var s *sortedset.SortedSet

	set.AddOrUpdate("a", 89, "Kelly")
	set.AddOrUpdate("b", 100, "Staley")
	set.AddOrUpdate("c", 150, "Jordon")
	set.AddOrUpdate("d", 120, "Bert")
	//set.PopMin()
	//node := set.PeekMax()
	nodes := set.GetByScoreRange(90, 130, &sortedset.GetByScoreRangeOptions{
		ExcludeStart: false,
	})

	for _, node := range nodes {
		fmt.Println(node.Value)
	}
}
