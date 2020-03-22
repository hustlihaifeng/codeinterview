package main

import "sort"

func main() {

}

/*
# 问题
1. 俄罗斯套娃问题

# 思路
1. 先按照X排序，从小到大扫描，扫描过的按照深度排序（同时保存该深度下的最小y值）；那么后面的，按照深度从大到小找，如果y满足，后面的就是深度加一；注意要将同长度不同深度的套娃小的放在后面，这样避免同长度的被当做大的套娃；至于同长度的取最小y，不管同长度的y如何排序，都能实现(同深度但是高度小的可以覆盖。因为x不会下降，所以不会有问题)

# 伪代码
func (){
    按照x升，y降排序
    用数组minheight保存当前深度下的最小y值
    for xy := range envelopes {
        idx:= findFirstSmallFromRight(minheight,xy[1])
        if idx==-1{
            if len(minheight)==0 {
                minheight = append(minheight,xy[1])
            }else{
                minheight[0]=xy[1] // 相当于有了一个新的分支，但是minheight[1]保留了以往的老分支的结果.因为x递增所以对0的这个替换可以，也不用考虑minheight[1]本次的连带更新，应为1的x目前还要小，是另一个分支的。
            }
        }else{
            if idx < len(minheight)-1 {
                if minheight[idx+1]>xy[1] {
                    minheight[idx+1]=xy[1]
                }
            }else{
                minheight=append(minheight,xy[1])
            }
        }
    }
    return len(minheight)
}
*/

func maxEnvelopes(envelopes [][]int) int {
	// 按照x升，y降排序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	// 用数组minheight保存当前深度下的最小y值
	var minheight []int
	for _, xy := range envelopes {
		idx := findFirstSmallBinSearch(minheight, xy[1])
		if idx == -1 {
			if len(minheight) == 0 {
				minheight = append(minheight, xy[1])
			} else {
				minheight[0] = xy[1] // 相当于有了一个新的分支，但是minheight[1]保留了以往的老分支的结果.因为x递增所以对0的这个替换可以，也不用考虑minheight[1]本次的连带更新，应为1的x目前还要小，是另一个分支的。
			}
		} else {
			if idx < len(minheight)-1 {
				if minheight[idx+1] > xy[1] {
					minheight[idx+1] = xy[1]
				}
			} else {
				minheight = append(minheight, xy[1])
			}
		}
	}
	return len(minheight)
}

// 因为这个height只会逐渐变小，不会影响有序性质，所以可以二分查找
func findFirstSmallFromRight(sli []int, x int) int {
	for idx := len(sli) - 1; idx >= 0; idx-- {
		if sli[idx] < x { // 等于套不上
			return idx
		}
	}
	return -1
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
