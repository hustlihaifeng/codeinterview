/*
 * @lc app=leetcode id=354 lang=golang
 *
 * [354] Russian Doll Envelopes
 *
 * https://leetcode.com/problems/russian-doll-envelopes/description/
 *
 * algorithms
 * Hard (33.67%)
 * Total Accepted:    42.6K
 * Total Submissions: 126.5K
 * Testcase Example:  '[[5,4],[6,4],[6,7],[2,3]]'
 *
 * You have a number of envelopes with widths and heights given as a pair of
 * integers (w, h). One envelope can fit into another if and only if both the
 * width and height of one envelope is greater than the width and height of the
 * other envelope.
 *
 * What is the maximum number of envelopes can you Russian doll? (put one
 * inside other)
 *
 * Note:
 * Rotation is not allowed.
 *
 * Example:
 *
 *
 *
 * Input: [[5,4],[6,4],[6,7],[2,3]]
 * Output: 3
 * Explanation: The maximum number of envelopes you can Russian doll is 3
 * ([2,3] => [5,4] => [6,7]).
 *
 *
 *
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	var in [][]int
	var out int

	in = in3
	out = out3
	fmt.Printf("%v %v, expect %v\n", in, maxEnvelopes(in), out)
}

/*
1. 暴力破解：将套娃按照长度排序（O(nlog(n))），后面就是高度的最大递增子序列长度问题了（注意对于同长度的套娃选择一个长度比前面套娃打的最小的那个就可以）
2. 最大递增子序列长度问题：记录下到当前为止，长度为i的序列的尾点（这些尾点一定是递增的）。此时对于一个新点，找前面长宽都不超过它的点，如果没找到，说明当前点有可能成为第一个点（长度递增，要么此点与0点长度相等，此时此点高度不小，忽略；要么此点比0点长度大，如果此点的高度比0点小，那么此点有可能成为0点）
*/
/*
1,100 100,1 2,2 3,3 4,4 5,100 6,6 7,100 8,8
1,1 2,2 3,3 4,100 100,4 5,100 100,5 6,100 100,6
1,1 2,2 2,4 4,3
*/

/*暴力破解：
1. 先按照长度增加排序（O(nlog(n))）
2. 找到长宽都比当前点小的点的长度，本点长度为该点长度加1(O(n^2))，并同步更新整体最大值*/

/* 暴力破解优化：
步骤：
1. 将点按照长度递增、长度相等则高度递减的方式来排序
2. 对于一个点，找到高度值大于等于该点的第一个点，那么此时新点能将原来的点替换掉（高度更小的点：也即后面的点，长度一定能满足，此时找高度更小的点来降低后面高度，进而得到最大的深度）。此时思路从找比前面点大的最小点；变成了找满足条件的最小点，满足条件即比前面点大；进一步是高度等于或者大于当前点的最小点能被当前点覆盖（从能套住那些，变成哪些点能被当前点覆盖）。
*/
/*伪代码：
将点按照深度递增、深度相同则高度递减排序
最大下标设置为0
for 每一个点 {
	在已有的当前深度值的最小高度点数组里面找等于或者大于当前点的第一个点的下标
	将该下标的点覆盖为当前点
}
返回深度数组长度
*/

/* 思路整理20190502：
这里的找第一个比当前点高的点的下标，然后赋值为当前点的操作。就是对后面的长度大高度小的点的一次二次查找。因为后面的一定是长度大或者等于的点，所以长度覆盖了没啥，只要找到某个深度下高度最小的点就行。按长度递增排序，保证了后面可以只考虑高度。同长度时大的在前面vs同长度是小的在前面？如果高度小点在前，那么后面同长度的高度高点就可以增加一个深度，而实际上由于长度相同，是不能增加高度的；所以要把同长度下高度大的点放在前面，这样不会有同长度是高度覆盖导致深度错误增加的问题；同时，高度大的在前后，如果同一高度的最后被放到同一位置，那么就是找了满足条件的的高度最低点；如果不是，那也不会被误以为发生了同长度是高度覆盖增加深度的问题，因为原来高度长的前面是有对应的小套娃的，而不是这个新加入的同长度小套娃。那么现在就确定了先按长度递增排序，再按高度递减排序，然后找第一个高度比当前点大的点的索引，将当前点覆盖到该索引处。特殊情况下，与当前点高度相等的点呢？因为后面高度是一直增加的，所以覆盖了不会有问题；如果不覆盖呢？那么当前点只有覆盖到前一个位置，深度要小1，而实际上当前点是可以覆盖到已经存在的高度相同的点的，这样会导致当前点的深度值算错，所以必须覆盖。
*/
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool { // [0] x [1] y
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	var minPointOfHeight [][]int
	for _, point := range envelopes {
		gePos := FindGePos(minPointOfHeight, point[1])
		if gePos < len(minPointOfHeight) {
			minPointOfHeight[gePos] = point
		} else {
			minPointOfHeight = append(minPointOfHeight, point)
		}
	}

	return len(minPointOfHeight)
}

func FindGePos(minPointOfHeight [][]int, height int) int {
	left := 0
	right := len(minPointOfHeight) - 1
	for left <= right {
		middle := (left + right) / 2
		if minPointOfHeight[middle][1] == height {
			return middle
		} else if minPointOfHeight[middle][1] < height {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return left
}
