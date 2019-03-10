package list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (s *ListNode) String() string {
	var rst string
	head := s
	for head != nil {
		rst += fmt.Sprintf("%v ", head.Val)
		head = head.Next
	}

	return rst
}

func NewListNode(nums []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, val := range nums {
		tail.Next = &ListNode{Val: val, Next: nil}
		tail = tail.Next
	}

	return dummy.Next
}
