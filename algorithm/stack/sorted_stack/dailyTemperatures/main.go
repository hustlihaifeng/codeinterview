package main

import "container/list"

func main() {

}

/*
# 1. 问题
1. 找出数组中第一个比当前数大的元素，输出两者之前的距离，详见<https://leetcode.com/problems/daily-temperatures/>
2. 要求O(n)的复杂度

# 思路
1. 暴力破解：一个二维循环即可，`O(n*n)`的复杂度
2. 如果每个点，从后往前找，把比当前点小的都处理掉，那么剩下的点，都会是比当前点大的点，那么所有剩下的点，递减。由于递减，那么每个元素，只会被访问两遍：进和出（如果比当前点大，那么当前点可以结束；否则，可以优先判断下一个点）。
	- 关键点是，剩下的数，递减
	- 把剩下的数另外存储，来节省时间

# 伪代码
func dailyTemperatures(T []int) []int {
    for idx,num := range T {
		for lst.Len() > 0 {
			if num > lst.Back() {
				e:=lst.Remove(lst.Back())
				rst[e.idx]=idx-e.idx
			}else{
				lst.PushBack(num,idx)
				break
			}
		}
		if lst.Len() == 0 {
			lst.PushBack(num,idx)
		}
	}
	for lst.Len()>0 {
		e := lst.Remove(lst.Back())
		rst[e.idx]=0
	}
}
*/

func dailyTemperatures(T []int) []int {
	type kv struct {
		k int
		v int
	}
	rst := make([]int, len(T), len(T))
	lst := list.New()
	for idx, num := range T {
		for lst.Len() > 0 {
			if num > lst.Back().Value.(*kv).v {
				e := lst.Remove(lst.Back())
				oidx := e.(*kv).k
				rst[oidx] = idx - oidx
			} else {
				lst.PushBack(&kv{idx, num})
				break
			}
		}
		if lst.Len() == 0 {
			lst.PushBack(&kv{idx, num})
		}
	}
	return rst
}
