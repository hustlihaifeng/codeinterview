/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
/*
 num
*/
func twoSum(nums []int, target int) []int {
	num := len(nums)
	for i := 0; i < num-1; i++ {
		for j := i + 1; j < num; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
