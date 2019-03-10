/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 *
 * https://leetcode.com/problems/linked-list-cycle-ii/description/
 *
 * algorithms
 * Medium (30.69%)
 * Total Accepted:    198K
 * Total Submissions: 638.4K
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * Given a linked list, return the node where the cycle begins. If there is no
 * cycle, return null.
 *
 * To represent a cycle in the given linked list, we use an integer pos which
 * represents the position (0-indexed) in the linked list where tail connects
 * to. If pos is -1, then there is no cycle in the linked list.
 *
 * Note: Do not modify the linked list.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: head = [3,2,0,-4], pos = 1
 * Output: tail connects to node index 1
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * second node.
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2], pos = 0
 * Output: tail connects to node index 0
 * Explanation: There is a cycle in the linked list, where tail connects to the
 * first node.
 *
 *
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1], pos = -1
 * Output: no cycle
 * Explanation: There is no cycle in the linked list.
 *
 *
 *
 *
 *
 *
 * Follow up:
 * Can you solve it without using extra space?
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
环形链表检测，使用快慢指针，没有额外的存储，快指针每次走2步，慢指针每次走一步，在快指针追慢指针的过程中，两者距离每走一步会减小1，所以如果有环，那么两者一定能相遇。

假设慢指针目前在环点，根据上面分析的每走一步两者距离减1，假设环长度为r，那么最多需要r-1步，快慢指针就能相遇。也即最多n步，两者就能相遇。

设相遇时慢指针走了k步，那么k+mr=2k，mr=k。由于k是r整数倍，那么再走k步的话，一定能回到当前点。这样在相遇点，让两个指正的速度都变为1，一个从当前相遇点开始走，一个从链表起始点开始走，两者在k步后一定都能在当前点相遇。由于两者步速相同，那么从环点到当前点的这段路程，两者是重合的，那么两者一定会在环点相遇。

伪代码：
p2,p1设为head
first
for p2!=nil && (first || p2!=p1) {
    first=false
    p1=p1.next
    p2=p2.next
    if p2==nil{
        break
    }
    p2=p2.next
}
if p2==ni{
    return nil
}
p1=head
for p2!=p1 {
    p2=p2.next
    p1=p1.next
}
return p1
*/
func detectCycle(head *ListNode) *ListNode {
	p1 := head
	p2 := head
	first := true
	for p2 != nil && (first || p2 != p1) {
		first = false
		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			break
		}
		p2 = p2.Next
	}

	if p2 == nil {
		return nil
	}

	p1 = head
	for p2 != p1 {
		p2 = p2.Next
		p1 = p1.Next
	}

	return p1
}
