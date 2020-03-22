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
找到环点，没有则是nil

# 思路
1. 先判断是否有环，用快慢指针法，有环会追赶，没环会走到nil。
伪代码：
func getMeetPoint(){
    for p2!=nil && p2.Next !=nil && p2.Next != p1 {
        p1=p1.Next
        p2 = p2.Next.Next
    }
    if p2 == nil || p2.Next == nil{
        return nil
    }
    // p2.Next==p1
    return p1.Next
}
2. 找到环点后，假设第一次相遇时，慢指针走了S步,环长为L，那么：2S=S+NL；则S=NL,那么从相遇点开始，再走S步依然能到相遇点，从链表头到相遇点也是S步；那么从相遇开始，一个从相遇点开始走，一个从链表头开始走，每次都走一步，S步后，两者会在相遇点相遇；由于两者速度一样，所以从环点到相遇点两者重合；所以这样两者会在环点第一次相遇。
伪代码：
func getCrossPoint(){
    获取相遇点pfirst
    for p1,p2:= head,pfirst;p2!=p1;{
        p1=p1.Next
        p2=p2.Next
    }
    return p2
}
*/
func detectCycle(head *ListNode) *ListNode {
	return getCrossPoint(head)
}

// 1. 先判断是否有环，用快慢指针法，有环会追赶，没环会走到nil。
func getMeetPoint(head *ListNode) (rst *ListNode) {
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil && p2.Next != p1 {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	if p2 == nil || p2.Next == nil {
		return nil
	}
	// p2.Next==p1
	return p1.Next
}

// 2. 找到环点后，假设第一次相遇时，慢指针走了S步,环长为L，那么：2S=S+NL；则S=NL,那么从相遇点开始，再走S步依然能到相遇点，从链表头到相遇点也是S步；那么从相遇开始，一个从相遇点开始走，一个从链表头开始走，每次都走一步，S步后，两者会在相遇点相遇；由于两者速度一样，所以从环点到相遇点两者重合；所以这样两者会在环点第一次相遇。
func getCrossPoint(head *ListNode) (rst *ListNode) {
	pfirst := getMeetPoint(head)
	if pfirst == nil {
		return nil
	}

	p1, p2 := head, pfirst
	for p2 != p1 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
