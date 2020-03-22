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
1. 链表相交点检测，要求O(n)的时间，O(1)的空间，不能修改链表结构

# 思路
1. O(1)的空间，只有走遍历模式。找交点，一般需要通过某种方式让两者同时到达交点，或者从同时到达交点后开始走公共路线，也即要求两者在某种情况下走的长度一样。
2. 本例中a走完了走b，b走完了走a，这样长度相等。且最后一段公共的路径是一样的，那么就会在交点相遇。
3. 注意边界条件预留TODO:
4. 在等式双方改变后，注意检查等式是否满足

# 伪代码
func (){
	pa,pb := heada,headb
	for pb != pa {
		if pb == nil {
			pb = heada
		}
		if pa == nil {
			pa = headb
		}
		pa = pa.Next
		pb = pb.Next
	}
	return pb
}
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB // [3] [2 3]
	for pb != pa {         // 3,2
		if pb == nil {
			pb = headA
		}
		if pa == nil {
			pa = headB // 3
		}
		if pa == pb { // TODO:在等式双方改变后，注意检查等式是否满足
			return pa
		}
		pa = pa.Next // nil 3
		pb = pb.Next // 3   nil
	}
	return pb
}
