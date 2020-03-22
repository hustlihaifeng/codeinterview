package main

import "fmt"

func main() {
	in := []int{4, 0, -4, -2, 2, 5, 2, 0, -8, -8, -8, -8, -1, 7, 4, 5, 5, -4, 6, 6, -3}
	out := 5
	fmt.Printf("in:%v out:%v my:%v\n", in, out, longestConsecutive(in))
}

/*
# 问题
1. 非排序数组中，找出排序后的最长连续序列长度，要求O(n)的复杂度
# 思路
1. 用一个map，保存到目前为止，比该点大的序列长度和比该点小的序列长度。新加的元素，看加减1是否有匹配的：没匹配则1；单匹配则原来的加1；双匹配则串联两边. 问题在于只更新了相邻的！！！！如果所有的都更更新的话，那么复杂度高，但是知道长度，我们就可以计算出端点值，然后更新端点值（这是由于key是数字的特性导致的。这样相邻值也就不用更新了）。
2. 如果map能保序列，那么用map来实现桶排序即可。
3. 以往的思路是，所有加入map，然后遍历map，对一个值，找上下连续的，并删除。那么总体似乎也是一个被访问常数次。

# 伪代码
lenMap都置位0
for _,nu := range nums {
    已经存在则continue
    upLen := lenMap[nu+1]
	downLen := lenMap[nu-1]
	upKey := nu+upLen
	downKey := nu-downLen
	curLen := upLen+downLen+1
	lenMap[upKey]=curLen
	lenMap[downKey]=curLen
	if curLen > maxLen {
		maxLen=curLen
	}
	return maxLen
}
遍历map，去最大值
*/

func longestConsecutive(nums []int) int {
	lenMap := make(map[int]int)
	for nu := range nums {
		lenMap[nu] = 0
	}
	maxLen := 0
	for _, nu := range nums {
		if lenMap[nu] > 0 {
			continue
		}
		upLen, ok := lenMap[nu+1]
		if !ok {
			upLen = 0
		}
		downLen, ok := lenMap[nu-1]
		if !ok {
			downLen = 0
		}
		upKey := nu + upLen
		downKey := nu - downLen
		curLen := upLen + downLen + 1
		lenMap[upKey] = curLen
		lenMap[downKey] = curLen
		lenMap[nu] = curLen // 上面的重复判断是按照大小来的，所以这部必须有
		if curLen > maxLen {
			maxLen = curLen
		}
		// fmt.Printf("%v map:%v\n", nu, lenMap)
	}
	return maxLen
}
