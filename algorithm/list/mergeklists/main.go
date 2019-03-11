/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (32.64%)
 * Total Accepted:    347.5K
 * Total Submissions: 1M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * Merge k sorted linked lists and return it as one sorted list. Analyze and
 * describe its complexity.
 *
 * Example:
 *
 *
 * Input:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * Output: 1->1->2->3->4->4->5->6
 *
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
)

/*
将链表数组二分，分到两个时进行合并,每个需要合并log(k)次（k指链表个数），所以是O（nlog(k))的复杂度。 伪代码：

if low==high{
    return sli[low]
}
if low+1==high{
    return merge2(sli[low],sli[high])
}
mid := (low+high)/2
lowSli := mergeKLists()
highSli := mergeKLists()
return merge2(lowSli,highSli)
*/
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	} else if length == 1 {
		return lists[0]
	} else if length == 2 {
		return merge2(lists[0], lists[1])
	}

	middle := length / 2
	p1 := mergeKLists(lists[0:middle])
	p2 := mergeKLists(lists[middle:length])
	return merge2(p1, p2)
}

func merge2(p1, p2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			tail, p1 = addAndNext(tail, p1)
		} else {
			tail, p2 = addAndNext(tail, p2)
		}
	}
	if p1 != nil {
		tail.Next = p1
	} else {
		tail.Next = p2
	}

	return dummy.Next
}

func addAndNext(tail, p *ListNode) (*ListNode, *ListNode) {
	pnext := p.Next
	tail.Next = p
	p.Next = nil
	return p, pnext
}
