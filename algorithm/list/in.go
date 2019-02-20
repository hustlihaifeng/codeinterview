package list

var NilList = NewList()
var CyclicList1 = NewCyclicList(6, 1)
var CyclicList2 = NewCyclicList(6, 2)
var CyclicList3 = NewCyclicList(6, 3)
var CyclicList4 = NewCyclicList(6, 4)
var CyclicList5 = NewCyclicList(6, 5)
var CyclicList6 = NewCyclicList(6, 6)
var LineList1 = NewLineList(1)
var LineList2 = NewLineList(2)
var LineList6 = NewLineList(6)

func NewCyclicList(total, cyclicPoint int) *List {
	lst := NewList()
	var cyclicElem *Elem
	for i := 1; i <= total; i++ {
		e := lst.PushBack(i)
		if i == cyclicPoint {
			cyclicElem = e
		}
	}
	lst.Back.Next = cyclicElem

	return lst
}

func NewLineList(num int) *List {
	lst := NewList()
	for i := 1; i <= num; i++ {
		lst.PushBack(i)
	}

	return lst
}
