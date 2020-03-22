# 问题
1. 非排序数组中，找出排序后的最长连续序列长度，要求O(n)的复杂度

# 思路
1. 用一个map，保存到目前为止，比该点大的序列长度和比该点小的序列长度。新加的元素，看加减1是否有匹配的：没匹配则1；单匹配则原来的加1；双匹配则串联两边. 问题在于只更新了相邻的！！！！如果所有的都更更新的话，那么复杂度高，但是知道长度，我们就可以计算出端点值，然后更新端点值（这是由于key是数字的特性导致的。这样相邻值也就不用更新了）。
2. 如果map能保序列，那么用map来实现桶排序即可。
3. 以往的思路是，所有加入map，然后遍历map，对一个值，找上下连续的，并删除。那么总体似乎也是一个被访问常数次。

# 伪代码
```go
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
```

# 详见
1. [main.go](main.go)