package main

func main() {

}

/*
问题：最长连续递增序列长度
*/
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	cur := 1
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] > nums[idx-1] {
			cur++
			if cur > max {
				max = cur
			}
		} else {
			cur = 1
		}
	}
	return max
}
