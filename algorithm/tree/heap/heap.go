package main

import "sort"

///////////////////////////////////////public interface//////////////////
type Heap interface {
	sort.Interface
	Push(interface{}) // Push and Pop 后进先出
	Pop() interface{}
}

func Push(s Heap, e interface{}) {
	s.Push(e)
	idx := s.Len() - 1
	FixToRoot(s, idx)
}

func Pop(s Heap) interface{} {
	if s.Len() == 0 {
		return nil
	}

	s.Swap(0, s.Len()-1)
	rst := s.Pop()
	FixToLeaf(s, 0)

	return rst
}

func Fix(s Heap, idx int) {
	FixToRoot(s, idx)
	FixToLeaf(s, idx)
}

func Remove(s Heap, idx int) interface{} {
	if s.Len() < idx+1 {
		return nil
	}

	s.Swap(idx, s.Len()-1)
	rst := s.Pop()

	if s.Len() > idx {
		Fix(s, idx)
	}
	return rst
}

///////////////////////////////////////internal interface////////////////
func FixToRoot(s Heap, idx int) {
	parIdx := (idx - 1) / 2

	if parIdx >= 0 && s.Less(idx, parIdx) {
		s.Swap(idx, parIdx)
		FixToRoot(s, parIdx)
	}
}

func FixToLeaf(s Heap, idx int) {
	left := 2*idx + 1
	right := 2*idx + 2
	if s.Len() >= right+1 {
		var lessIdx int
		if s.Less(left, right) {
			lessIdx = left
		} else {
			lessIdx = right
		}

		if s.Less(lessIdx, idx) {
			s.Swap(lessIdx, idx)
			FixToLeaf(s, lessIdx)
		}
	} else if s.Len() == left+1 {
		if s.Less(left, idx) {
			s.Swap(left, idx)
		}
	}
}
