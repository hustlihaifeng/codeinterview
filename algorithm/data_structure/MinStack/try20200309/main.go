package main

func main() {

}

/*
# 问题
1. 最小栈，push、pop、top、getMin都是O(1)的时间复杂度

# 思路
1. push pop top这几个都是常规操作，是O(1)的复杂度
2. getMin要实现O(1)的话，那么需要把每次操作后的最小值都保存下来。由于push pop只会对栈顶修改，所有还好。
3. 用一个slice来作为栈，用一个链表来保存最小值。push pop可能对最小值有改变。
4. 用val来标志当前push  pop的是否是最小值，这样有几个相等值就需要截个节点。
*/
type MinStack struct {
	sli []int
	min *node
}
type node struct {
	val  int
	next *node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.sli = append(this.sli, x)
	if this.min == nil || x <= this.min.val {
		new := &node{val: x, next: this.min}
		this.min = new
	}
}

func (this *MinStack) Pop() {
	val := this.sli[len(this.sli)-1]
	this.sli = this.sli[:len(this.sli)-1]
	if val == this.min.val {
		this.min = this.min.next
	}
}

func (this *MinStack) Top() int {
	return this.sli[len(this.sli)-1]
}

func (this *MinStack) GetMin() int {
	return this.min.val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
