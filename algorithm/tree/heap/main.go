package main

import "fmt"

func IntLess(i, j interface{}) bool {
	return i.(int) < j.(int)
}

type IntHeap []int

func (s IntHeap) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s IntHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntHeap) Len() int {
	return len(s)
}
func (s *IntHeap) Push(e interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*s = append(*s, e.(int))
}
func (s *IntHeap) Pop() interface{} {
	old := *s
	if old.Len() == 0 {
		return nil
	}

	rst := old[old.Len()-1]
	*s = old[0 : old.Len()-1]

	return rst
}

func main() {
	ph := &IntHeap{}
	Push(ph, 3)
	fmt.Printf("%v %v\n", (*ph).Len(), (*ph)[0])
	Push(ph, 2)
	fmt.Printf("%v %v\n", (*ph).Len(), (*ph)[0])
	Push(ph, 1)
	fmt.Printf("%v %v\n", (*ph).Len(), (*ph)[0])
	Push(ph, 5)
	fmt.Printf("%v %v\n", (*ph).Len(), (*ph)[0])
	Push(ph, 4)
	fmt.Printf("%v %v\n", (*ph).Len(), (*ph)[0])

	fmt.Println()

	fmt.Println(Pop(ph))
	fmt.Println(Pop(ph))
	fmt.Println(Pop(ph))
	fmt.Println(Pop(ph))
	fmt.Println(Pop(ph))
	fmt.Println(Pop(ph))
	fmt.Println()

	Push(ph, 3)
	Push(ph, 2)
	Push(ph, 1)
	Push(ph, 4)
	Push(ph, 5)
	fmt.Println(Remove(ph, 0))
	fmt.Println(Remove(ph, 0))
	fmt.Println(Remove(ph, 0))
	fmt.Println(Remove(ph, 0))
	fmt.Println(Remove(ph, 0))
	fmt.Println(Remove(ph, 0))
}
