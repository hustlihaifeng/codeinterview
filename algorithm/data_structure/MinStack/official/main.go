package main

func main() {

}

// MinStack me
type MinStack struct {
	s   []int
	min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}

}

func (this *MinStack) Push(x int) {
	this.s = append(this.s, x)
	length := len(this.min)
	if length == 0 {
		this.min = append(this.min, this.s[0])
		return
	}
	if this.min[length-1] > x {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[length-1])
	}
}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
