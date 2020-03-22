package main

func main() {

}

/*
# 问题
1. 找数组中和最大的字串

# 思路
1. 从左到右扫描，在和为正时继续，为负时重新开始。O(n)
2. 分治法
*/
func maxSubArray(nums []int) int {
	max := nums[0]
	cur := max
	for i := 1; i < len(nums); i++ {
		if cur < 0 {
			cur = 0
		}
		cur += nums[i]
		if cur > max {
			max = cur
		}
	}
	return max
}
