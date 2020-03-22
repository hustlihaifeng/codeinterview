package main

import (
	"sort"
)

func main() {

}

/*
关键点：将找三数转化为，在增序数组中，找两数之和（O(n)）。固定一个，剩下两数之和与-item比较，这样增减方向固定。
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	rs, nlen := make([][]int, 0), len(nums)
	for i := 0; i < nlen; i++ {
		t := -nums[i]
		for l, r := i+1, nlen-1; l < r; {
			sum := nums[l] + nums[r]
			if sum == t {
				rs = append(rs, []int{nums[i], nums[l], nums[r]})
				for ; l < r && nums[l] == nums[l+1]; l++ {
				}
				l++
				for ; r > l && nums[r] == nums[r-1]; r-- {
				}
				r--
			} else if t < sum {
				r--
			} else {
				l++
			}
		}
	}
	return rs
}
