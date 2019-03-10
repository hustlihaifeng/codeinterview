/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 *
 * https://leetcode.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (33.66%)
 * Total Accepted:    171.3K
 * Total Submissions: 503.5K
 * Testcase Example:  '[4,2,1,3]'
 *
 * Sort a linked list in O(n log n) time using constant space complexity.
 *
 * Example 1:
 *
 *
 * Input: 4->2->1->3
 * Output: 1->2->3->4
 *
 *
 * Example 2:
 *
 *
 * Input: -1->5->3->4->0
 * Output: -1->0->3->4->5
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import (
	. "algorithm/list"
	"fmt"
)

func main() {
	var in []int
	var root *ListNode

	in = []int{4, 2, 1, 3}
	root = NewListNode(in)
	fmt.Println(root)
	fmt.Printf("%v\n", sortList(root))
}

func sortList(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head} // 因为dummy.Next，dummy遍历变化时，可能会使head变化
	listLen := getLen(head)        // O(n)

	for step := 1; step < listLen; step = step << 1 { // O(log(n))
		tail := dummy
		cur := dummy.Next
		for cur != nil {
			left := cur
			cur = split(cur, step)
			right := cur
			cur = split(cur, step)
			tail = merge(left, right, tail) // O(step) O(n)
		}
		//fmt.Println(dummy)
	}
	return dummy.Next
}

func split(cur *ListNode, step int) (rsta *ListNode) {
	/*fmt.Printf("get %v elem from %v", step, cur)
	curin := cur
	defer func() {
		fmt.Printf(" get %v, left %v\n", curin, rsta)
	}()*/

	for cnt := 1; cur != nil && cnt < step; cur = cur.Next { // 1+n-1=n
		cnt++
	}

	if cur == nil {
		return nil
	}
	pnext := cur.Next
	cur.Next = nil
	return pnext
}
func addAndNext(ptail, pnode **ListNode) {
	pnext := (*pnode).Next

	(*ptail).Next = *pnode
	*ptail = (*ptail).Next
	(*ptail).Next = nil

	*pnode = pnext
}
func merge(left, right, tail *ListNode) (rsta *ListNode) {
	/*fmt.Printf("merge %v and %v ", left, right)
	tailin := tail
	defer func() {
		fmt.Printf("get %v\n", tailin.Next)
	}()*/

	for left != nil && right != nil {
		if left.Val <= right.Val {
			addAndNext(&tail, &left)
		} else {
			addAndNext(&tail, &right)
		}
	}

	if left != nil {
		tail.Next = left
	} else {
		tail.Next = right
	}

	for tail.Next != nil {
		tail = tail.Next
	}

	return tail
}

func getLen(head *ListNode) (rsta int) {
	/*defer func() {
		fmt.Printf("len of %v is %v\n", head, rsta)
	}()*/

	cnt := 0
	for head != nil {
		cnt++
		head = head.Next
	}

	return cnt
}
