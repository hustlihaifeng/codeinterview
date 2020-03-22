package main

func main() {

}

/*
# 问题
1. 求竖方块之间能存多少水。

# 思路
1. 从左和右分别开始，找递增块，两个递增块之间有可能储水，低块是水位线。分为三个部分，左，右，中间来分别计算水。

# 伪代码
func (){
	从左到右找递增块
	从右到左找递增块
	从左到计算水量，到左边最大值为止
	从右到左计算水量，到右边最大值为止
	如果左右最大值位置不相等，从左最大值到右最大值计算水量
}

*/

func trap(height []int) int {
	rst := 0
	lIncSli := findIncrease(height, 0, 1, len(height))
	rIncSli := findIncrease(height, len(height)-1, -1, -1)
	for i := 0; i < len(lIncSli)-1; i++ {
		rst += calWater(height, lIncSli[i], 1, lIncSli[i+1])
	}
	for i := 0; i < len(rIncSli)-1; i++ {
		rst += calWater(height, rIncSli[i], -1, rIncSli[i+1])
	}
	// 判断incSli长度为0时
	if len(lIncSli) > 0 && len(rIncSli) > 0 {
		lMaxIdx, rMaxIdx := lIncSli[len(lIncSli)-1], rIncSli[len(rIncSli)-1]
		if lMaxIdx != rMaxIdx {
			rst += calWater(height, lMaxIdx, 1, rMaxIdx)
		}
	}

	return rst
}

// sli 长度为0，或者sli全为0，则返回nil
func findIncrease(sli []int, bgnIdx, offset, endIdx int) (rst []int) {
	for i := bgnIdx; i != endIdx; i = i + offset {
		if sli[i] > 0 {
			if len(rst) == 0 {
				rst = append(rst, i)
			} else {
				lastIdx := rst[len(rst)-1]
				if sli[i] > sli[lastIdx] {
					rst = append(rst, i)
				}
			}
		}
	}
	return rst
}

func calWater(sli []int, bgnIdx, offset, endIdx int) (rst int) {
	rst = 0
	for i := bgnIdx; i != endIdx; i += offset { // endIdx大
		rst += sli[bgnIdx] - sli[i]
	}
	return rst
}
