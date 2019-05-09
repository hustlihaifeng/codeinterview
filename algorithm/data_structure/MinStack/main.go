package main

func main() {

}

/*push pop top 都是普通的栈操作，getMin返回最小值。只有push pop能改变栈，其他的都是读取操作。然后由于是栈，后进先出，后面的大的值还在时获取的最小值一定是前面的已经压栈的最小值，对前面的已经存在的最小值无影响。
也即找最优解，问所有情况下的最优解。由于后进先出的特性，如果一个后来的值没有更新最优解，那么该值对对最优解没有任何影响，在处理最优解部分时直接忽略改制。在某个最优解元素还在的时候，最优解是该最优解；在该最优解元素被pop的时候，需要更新到他压栈时的最优解，终点是第一个入栈的元素。
关键点在于栈的后进先出特性，所以不需要考虑后来的没有更新最优解的值。又由于只需要得到当前还在元素的最优解，如果获取了当前还在元素的完整排序链路，那么是一种浪费。*/
type MinStack struct {
	Buff []int
	Min  *MinNode
}
type MinNode struct {
	Val  int
	Next *MinNode
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Buff: []int{}, Min: nil}
}

func (this *MinStack) Push(x int) {
	this.Buff = append(this.Buff, x)
	if this.Min == nil || x <= this.Min.Val {
		oldMin := this.Min
		this.Min = &MinNode{Val: x, Next: oldMin}
	}
}

func (this *MinStack) Pop() {
	if len(this.Buff) > 0 {
		maxIdx := len(this.Buff) - 1
		elem := this.Buff[maxIdx]
		this.Buff = this.Buff[:maxIdx]
		if elem == this.Min.Val {
			this.Min = this.Min.Next
		}
	}
}

func (this *MinStack) Top() int {
	return this.Buff[len(this.Buff)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min.Val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
