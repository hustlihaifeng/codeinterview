package main

func main() {

}

/*
# 问题
升序数组轮转了，变成了两段升序数组，没有重复的值，找某个值。相当于一段y=x以某个点为分界，发生了位置平移替换。
# 思路
思路1. 二分查找。应该先和目标值进行比较，来得到左边还是右边。然后和临值比价，看是否是断点。好复杂
# 伪代码
边界判断和初始化
for lidx < ridx {
    midx := (lidx+ridx)/2
    if mid值 就是 {
        return midx
    } else if mid 在左边线中 {
        if target > mid值{
            在mid的右边找
        } else {
            在mid的左边找
        }
    } else { // middle在右边线中
        if target > mid值{

        }
    }
}
if nums[lidx]==target {
    return lidx
}
return -1

思路2. 二分找到断点，然后在其中一个满足条件的递增区间找。

初始化和边界值处理
找断点
if target >= 0值 {
    在0,断点 二分找
}else {
    在断点+1,最右值二分找
}

找断点：
初始化和边界值处理
for lidx < ridx {
    二分得到midx
    if midx值大于0值 {
        断点在[midx,ridx]
    }else if 等于0值{
        当前值就是断点
    }else { // mid值小于0值
        断点在[lidx,midx]
    }
}

二分查找:
for lidx < ridx {
    if mid值就是{
        返回middle值
    } else if target > middle 值{
        [midx+1,lidx]
    }else{
        [lidx,midx-1]
    }
}
判断lidx==ridx时
*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	lidx, ridx := 0, len(nums)-1
	if nums[0] < nums[len(nums)-1] {
		lidx, ridx = 0, len(nums)-1
	} else {
		midx := findMidx(nums)
		if target >= nums[0] {
			lidx, ridx = 0, midx
		} else {
			lidx, ridx = midx+1, len(nums)-1
		}
	}
	return binarySearch(nums, lidx, ridx, target)
}

func findMidx(nums []int) (rst int) {
	lidx, ridx := 0, len(nums)-1
	for lidx < ridx {
		midx := (lidx + ridx) / 2
		if nums[midx] > nums[lidx] {
			lidx = midx
		} else if nums[lidx] == nums[midx] {
			return midx
		} else {
			ridx = midx
		}
	}
	return -1
}

func binarySearch(nums []int, lidx, ridx, target int) int {
	for lidx < ridx {
		midx := (lidx + ridx) / 2
		if target == nums[midx] {
			return midx
		} else if target > nums[midx] {
			lidx = midx + 1
		} else {
			ridx = midx - 1
		}
	}
	if target == nums[lidx] {
		return lidx
	}
	return -1
}
