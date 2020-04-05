package main

import "container/heap"

func main() {

}

/*
# 问题
1. 合并k个有序数组，详见 <https://leetcode.com/problems/merge-k-sorted-lists/>
2. 要求分析复杂度

# 思路
1. 将非空链都加入链堆中。
2. for 堆非空 {
    取出堆头链，删除链头并将链头加入结果链。
    如果该链非空，将该链加入堆中。
}
3. 时间复杂度：O(nlog(k))， 空间复杂度：k

*/
func mergeKLists(lists []*ListNode) *ListNode {
	var dummy ListNode
	tail := &dummy
	h := &ListHeap{}
	for _, lst := range lists {
		if lst != nil {
			heap.Push(h, lst)
		}
	}
	for h.Len() > 0 {
		lst := heap.Pop(h).(*ListNode)
		tail.Next = lst
		tail = tail.Next
		if lst.Next != nil {
			heap.Push(h, lst.Next)
		}
	}
	return dummy.Next
}

type ListHeap []*ListNode

func (s ListHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ListHeap) Less(i, j int) bool {
	return s[i].Val < s[j].Val
}
func (s ListHeap) Len() int {
	return len(s)
}
func (s *ListHeap) Push(e interface{}) {
	*s = append(*s, e.(*ListNode)) // &&
}
func (s *ListHeap) Pop() interface{} {
	rst := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return rst
}
