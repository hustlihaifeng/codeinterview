package main

import "fmt"

func main() {
	fmt.Printf("in:%v out:%v my:%v\n", in, out, ToSlice(sortList(GetList(in))))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func GetList(in []int) (rst *ListNode) {
	dummy := &ListNode{}
	tail := dummy
	for _, nu := range in {
		tail.Next = &ListNode{Val: nu}
		tail = tail.Next
		tail.Next = nil
	}
	return dummy.Next
}

func ToSlice(in *ListNode) (rst []int) {
	for in != nil {
		rst = append(rst, in.Val)
		in = in.Next
	}
	return rst
}

/*
# 问题
1. 链表排序，O(nlogn)时间复杂度（归并或者快排（平均）），O(1)空间复杂度（不能用递归）
# 思路
1. 先获取链表长度，然后根据链表长度遍历来进行切分和合并操作，刚好O(nlogn)

# 伪代码
func(){
	获取链表长度length
	for clen:=1;clen<length;clen=2*clen {
		循环找两个clen长度的有序链表，进行合并。（先两个nil），合并完之后在指向next
	}
}
*/
func sortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	// 获取链表长度length
	dummy.Next = head
	length := getLength(head)
	for clen := 1; clen < length; clen = 2 * clen {
		// 	循环找两个clen长度的有序链表，进行合并。（先两个nil），合并完之后在指向next
		// fmt.Printf("clen:%v length:%v\n", clen, length)
		p := dummy.Next
		pre := dummy
		var l1, l2 *ListNode
		for p != nil {
			l1, p = getSubList(p, clen)
			l2, p = getSubList(p, clen)
			l3, tail := merge(l1, l2)
			tail.Next = p
			pre.Next = l3
			pre = tail
		}
	}
	return dummy.Next
}

func getLength(head *ListNode) (rst int) {
	// fmt.Printf("enter getLength\n")
	p := head
	for p != nil {
		rst++
		p = p.Next
	}
	return rst
}

func getSubList(head *ListNode, clen int) (sub, remain *ListNode) {
	// fmt.Printf("enter getSubList，clen:%v\n", clen)
	tail := &ListNode{}
	remain = head
	cnt := 0
	for remain != nil && cnt < clen {
		tail.Next = remain
		remain = remain.Next
		tail = tail.Next
		tail.Next = nil
		cnt++
	}
	return head, remain
}

func merge(l1, l2 *ListNode) (rst, tail *ListNode) {
	// fmt.Printf("enter merge\n")
	dummy := &ListNode{}
	tail = dummy
	for l1 != nil || l2 != nil {
		if l1 == nil || (l1 != nil && l2 != nil && l1.Val > l2.Val) {
			tail.Next = l2
			l2 = l2.Next
		} else {
			tail.Next = l1
			l1 = l1.Next
		}
		tail = tail.Next
		tail.Next = nil
	}
	return dummy.Next, tail
}
