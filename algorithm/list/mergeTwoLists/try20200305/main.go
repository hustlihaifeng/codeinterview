package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
# 问题
1. 合并两个有序数组

# 思路
1. 用一个dummy节点来处理数组头问题。
2. l1 != nil || l2 != nil作为判断条件
3. 其实在一个为nil时，可以把结果直接指向另一个。此时循环的判断条件为：l1!=nil && l2!=nil
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	rst := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil || (l1 != nil && l2 != nil && l1.Val > l2.Val) {
			rst.Next = l2
			l2 = l2.Next
			rst = rst.Next
		} else {
			rst.Next = l1
			l1 = l1.Next
			rst = rst.Next
		}
	}
	return dummy.Next
}
