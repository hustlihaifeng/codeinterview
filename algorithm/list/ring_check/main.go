package main

import (
	"fmt"
	"lhf/leetcode/list"
)

func main() {
	fmt.Printf("%v is cyclic list ? %v\n", "NilList", isCyclicList(list.NilList))
	fmt.Printf("%v is cyclic list ? %v\n", "CyclicList1", isCyclicList(list.CyclicList1))
	fmt.Printf("%v is cyclic list ? %v\n", "CyclicList2", isCyclicList(list.CyclicList2))
	fmt.Printf("%v is cyclic list ? %v\n", "CyclicList3", isCyclicList(list.CyclicList3))
	fmt.Printf("%v is cyclic list ? %v\n", "CyclicList4", isCyclicList(list.CyclicList4))
	fmt.Printf("%v is cyclic list ? %v\n", "CyclicList5", isCyclicList(list.CyclicList5))
	fmt.Printf("%v is cyclic list ? %v\n", "CyclicList6", isCyclicList(list.CyclicList6))
	fmt.Printf("%v is cyclic list ? %v\n", "LineList1", isCyclicList(list.LineList1))
	fmt.Printf("%v is cyclic list ? %v\n", "LineList2", isCyclicList(list.LineList2))
	fmt.Printf("%v is cyclic list ? %v\n", "LineList6", isCyclicList(list.LineList6))
}

func isCyclicList(l *list.List) bool {

	for p1, p2 := l.Front, l.Front; p2 != nil && p1 != nil; p1, p2 = p1.Next, p2.Next {
		p2 = p2.Next
		if p2 == nil {
			break
		}

		if p2 == p1 || p2.Next == p1 { // p1 != nil
			return true
		}
	}

	return false
}
