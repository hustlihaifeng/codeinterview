/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 *
 * https://leetcode.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (40.43%)
 * Total Accepted:    206.9K
 * Total Submissions: 511.9K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * Given an unsorted array of integers, find the length of longest increasing
 * subsequence.
 *
 * Example:
 *
 *
 * Input: [10,9,2,5,3,7,101,18]
 * Output: 4
 * Explanation: The longest increasing subsequence is [2,3,7,101], therefore
 * the length is 4.
 *
 * Note:
 *
 *
 * There may be more than one LIS combination, it is only necessary for you to
 * return the length.
 * Your algorithm should run in O(n^2) complexity.
 *
 *
 * Follow up: Could you improve it to O(n log n) time complexity?
 *
 */
package main

import "fmt"

func main() {
	var in []int
	var out int
	in = []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}
	out = 6
	fmt.Printf("%v %v expect %v\n", in, lengthOfLIS(in), out)
}

/*
暴力破解：
对每个点，后面所有的大点中，结果最大的一个加1。没有大点，则是1。这是一个递推的过程。O(n^2)。
反过来的话，就是前面所有的小点的结果中的最大值加1，没有小点则是1.O(n^2)
优化：
1. 如果将前面的点排序，那么排序复杂度是O(nlog(n))，这样去找每个点好像依然是O(n)。
2. 记录下长度为i的序列中，最小尾值，那么这些尾值应该是递增的（不是的话用后面的序列的前一个值代替前一个长度的尾值即可），此时找比当前点小的那个点，其长度加一即可。关键点事尾值如何更新。当前得到一个值，与已有的尾值进行比较即可。
*/
/* 伪代码
minValueOfLenI[0]设置为数组里面的第一个值
for 数组里面的每一个值 {
    在minValueOfLenI里面找到第一个小于等于当前值的值的下标i
    if 没找到 { // 当前值是最小值
        minValueOfLenI[0]设置为当前值
    } else if 所找到的值等于当前值 {
        忽略
    } else { // 找到比当前值小的值
        if i == 当前最大下标 {
            当前值append到minValueOfLenI上 // 长度加一
        }else {
            if 当前值 < minValueOfLenI[i+1] {
                minValueOfLenI[i+1] 设置为当前值
            }
        }
    }

}

*/
func lengthOfLIS(nums []int) int {
	if lengthOfLISInputInvalid(nums) {
		return 0
	}
	minValueOfLenI := make([]int, 1, 10)
	minValueOfLenI[0] = nums[0]
	for _, val := range nums {
		lowIdx := FindLE(minValueOfLenI, val)
		// fmt.Printf("%v %v %v ", val, lowIdx, minValueOfLenI)
		if lowIdx == -1 {
			// fmt.Printf("1 ")
			minValueOfLenI[0] = val
		} else if minValueOfLenI[lowIdx] == val {
			// fmt.Printf("2 ")
		} else {
			// fmt.Printf("3 ")
			if lowIdx == len(minValueOfLenI)-1 {
				// fmt.Printf("3.1 ")
				minValueOfLenI = append(minValueOfLenI, val)
			} else {
				// fmt.Printf("3.2 ")
				if val < minValueOfLenI[lowIdx+1] {
					// fmt.Printf("3.2.1 ")
					minValueOfLenI[lowIdx+1] = val
				}
			}
		}
		// fmt.Printf("%v\n", minValueOfLenI)
	}
	return len(minValueOfLenI)
}

func FindLE(sortedSli []int, target int) int {
	low := 0
	maxIdx := len(sortedSli) - 1
	high := maxIdx
	for 0 <= low && low <= high && high <= maxIdx {
		middle := (low + high) / 2
		if sortedSli[middle] == target {
			return middle
		} else if sortedSli[middle] > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return high
}

func lengthOfLISInputInvalid(sli []int) bool {
	if sli == nil {
		return true
	} else if len(sli) == 0 {
		return true
	}
	return false
}
