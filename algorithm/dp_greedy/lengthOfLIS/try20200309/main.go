package main

func main() {

}

/*
# 问题
1. 最长递增子序列的长度（不需要连续）

# 思路
1. 用一个数组minh保存当前长度的最小值，这个值后续可更新（新的分支），但是数组有序性不变（因为只会把前边的变得更小），通过二分查找找到第一个小于当前值的点，然后进行更新。
*/
func lengthOfLIS(nums []int) int {
	var minh []int
	for _, nu := range nums {
		idx := findFirstSmallBinSearch(minh, nu)
		if idx == -1 {
			if len(minh) == 0 {
				minh = append(minh, nu)
			} else {
				minh[0] = nu
			}
		} else {
			if idx < len(minh)-1 {
				if minh[idx+1] > nu {
					minh[idx+1] = nu
				}
			} else {
				minh = append(minh, nu)
			}
		}
		// fmt.Printf("%v %v\n",nu,minh)
	}
	return len(minh)
}

func findFirstSmallBinSearch(sli []int, x int) int {
	if len(sli) == 0 {
		return -1
	}
	l, r := 0, len(sli)-1
	for l < r {
		m := (l + r) / 2
		if sli[m] >= x {
			r = m - 1
		} else { // sli[m] < x
			if sli[m+1] >= x {
				return m
			}
			l = m + 1
		}
	}
	if sli[l] < x {
		return l
	}
	return -1
}
