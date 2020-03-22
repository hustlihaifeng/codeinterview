package main

import (
	"fmt"
	"sort"
)

func main() {

}

/*
问题：在一个数字数组中，找三个数，是只和为0
思路关键点：在有序数组中，找两个和为指定值的数，从前后夹击，复杂度为O(n)，这是两数之和中得到的结论。所以两数之和的复杂度为O(nlog(n))
那么对于三数之和，定一个数，变成两数之和问题，复杂度为O(nlog(n)+n^2)
伪代码：
sort.Ints(nums)
var rst [][]int
for i:=0;i<len(nums);i++ {
	targ := -nums[i]
	for lidx,ridx := i+1,len(nums)-1; lidx<ridx; {
		sum := nums[lidx]+nums[ridx]
		if sum == targ {
			rst = append(rst,[]int{i,lidx,ridx})
			for ; lidx<ridx && nums[lidx]==nums[lidx+1]; lidx++{} // 这里因为要去重，所以需要这个，不然直接一个外层就可以
			for ; ridx>lidx && nums[ridx]==nums[ridx-1]; ridx--{}
		}else{
			if sum > targ {
				ridx--
			}else{
				lidx++
			}
		}
	}
}
return rst
*/
func threeSum(nums []int) [][]int {
	// TODO:边界检查 [0,0,0]
	sort.Ints(nums)
	fmt.Println(nums)
	var rst [][]int
	for i := 0; i < len(nums); i++ {
		targ := -nums[i]
		for lidx, ridx := i+1, len(nums)-1; lidx < ridx; {
			sum := nums[lidx] + nums[ridx]
			if sum == targ {
				rst = append(rst, []int{nums[i], nums[lidx], nums[ridx]})
				for ; lidx < ridx && nums[lidx] == nums[lidx+1]; lidx++ {
				} // 这里因为要去重，所以需要这个，不然直接一个外层就可以. 这里由于先有lidx<irdx，所以lidx+1不会超标
				// 此时nums[lidx]依然没变，两数之和，左边没变，右边依然没变
				lidx++
				for ; ridx > lidx && nums[ridx] == nums[ridx-1]; ridx-- {
				}
				ridx--
			} else {
				if sum > targ {
					ridx--
				} else {
					lidx++
				}
			}
		}
		for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {
		} // 这个去重需要理解
	}
	return rst
}
