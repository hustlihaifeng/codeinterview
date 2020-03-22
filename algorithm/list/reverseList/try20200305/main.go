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
1. 反转链表

# 思路
1. 直接反过去：cnode = p; p=p.Next; cnode.Next=head;head=cnode
*/
func reverseList(head *ListNode) *ListNode {
	var rst *ListNode
	for p := head; p != nil; {
		cnode := p
		p = p.Next
		cnode.Next = rst
		rst = cnode
	}
	return rst
}
