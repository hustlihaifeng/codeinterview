package main

import (
	"fmt"
)

/*
问题：
单链表，奇数位升序，偶数位降序，想要得到一个全降序的单链表；
1,8,3,4,5,2,7,0
要求：时间复杂度O(n), 空间复杂度O(1)
*/

/*
思路：
1. 将升序和降序链表分离
2. 升序链表翻转
3. 两个链表合并

pup,pdown cnt
for 链表里面的每一个节点{
    idx是偶数，加入到升序
    idx是奇数，加入到降序
}

head=nil
for 每个节点{
    next设为head
    head指向当前节点
    p指向下一个节点
}

注意:将最后一个元素的next指针设置为null（如在add的时候），不然会有意想不到的麻烦。
这里面的`String()`,`addAndNext`,`reverseList`,`mergeList`有代表性，特别是`addAndNext`
*/
func main() {
	var sli []int
	var in *Node

	sli = []int{1, 8, 3, 4, 5, 2, 7, 0}
	in = NewNode(sli)
	fmt.Printf("%v\n", in)
	fmt.Printf("%v\n", reorderList(in))
}

type Node struct {
	val  int
	next *Node
}

func (s *Node) String() string {
	var rst string

	p := s
	for p != nil {
		rst += fmt.Sprintf("%v ", p.val)
		p = p.next
	}

	return rst
}
func NewNode(ary []int) *Node {
	var rst *Node
	var ptail *Node
	for _, val := range ary {
		/*
			if rst == nil {
				rst = &Node{val: val} // 这里是一个局部变量，go的gc应该用来引用计数
				ptail = rst
			} else {
				ptail.next = &Node{val: val}
				ptail = ptail.next
			}*/
		pNode := &Node{val: val}
		addAndNext(&rst, &ptail, &pNode)
	}

	return rst
}

func reorderListInputInvalid(ph *Node) bool {
	if ph == nil {
		return true
	}

	return false
}
func reorderList(ph *Node) *Node {
	if reorderListInputInvalid(ph) {
		return nil
	}

	pup, pdown := splitList(ph)
	pupRev := reverseList(pup)
	rst := mergeList(pdown, pupRev)
	return rst
}

func splitList(ph *Node) (*Node, *Node) {
	var pup *Node
	var pdown *Node
	var pupHead *Node
	var pdownHead *Node
	cnt := 1
	for ph != nil {
		if cnt%2 == 0 { //偶数降序
			addAndNext(&pdownHead, &pdown, &ph)
		} else { // 奇数升序
			addAndNext(&pupHead, &pup, &ph)
		}
		cnt++
	}
	return pupHead, pdownHead
}

func reverseListInputInvalid(ph *Node) bool {
	if ph == nil {
		return true
	}

	return false
}
func reverseList(ph *Node) *Node {
	if reverseListInputInvalid(ph) {
		return nil
	}

	var head *Node
	head = nil
	for ph != nil {
		pnext := ph.next
		ph.next = head
		head = ph
		ph = pnext
	}

	return head
}

func mergeList(p1, p2 *Node) *Node {
	var pall *Node
	var rst *Node
	for p1 != nil && p2 != nil {
		if p1.val == p2.val {
			addAndNext(&rst, &pall, &p1)
			addAndNext(&rst, &pall, &p2)
		} else if p1.val > p2.val {
			addAndNext(&rst, &pall, &p1)
		} else {
			addAndNext(&rst, &pall, &p2)
		}
	}

	var pleft *Node
	if p1 != nil {
		pleft = p1
	} else if p2 != nil {
		pleft = p2
	} else {
		pleft = nil
	}
	for pleft != nil {
		addAndNext(&rst, &pall, &pleft)
	}

	return rst
}

func addAndNext(phead, ptail, pnode **Node) {
	pnext := (*pnode).next
	if *phead == nil {
		*phead = *pnode
		*ptail = *pnode
	} else {
		(*ptail).next = *pnode
		(*ptail) = *pnode
	}
	(*ptail).next = nil
	fmt.Printf("add %v to %v\n", (*pnode).val, *phead)
	*pnode = pnext
}
