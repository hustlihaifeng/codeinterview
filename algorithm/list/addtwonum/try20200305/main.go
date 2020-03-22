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
1. 数字用链表反序表示，计算两数之和

# 思路
1. 循环条件：l1 != nil || l2 != nil ,然后在里面对l1和l2进行判断操作
2. 因为有进位，所以不能直接指向剩下的非nil的链
3. 注意最后的进位add需要单独处理，简单起见，可以把循环条件改为：`l1 != nil || l2 != nil || add != 0`
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	rst := dummy
	add := 0
	for l1 != nil || l2 != nil || add != 0 {
		n1 := 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		n2 := 0
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + add
		add = sum / 10
		rst.Next = &ListNode{Val: sum % 10}
		rst = rst.Next
	}
	return dummy.Next
}
