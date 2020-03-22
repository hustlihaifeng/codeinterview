package main

import (
	"fmt"
	"strings"
)

func main() {
	in := [][]int{[]int{1, 1, 0, 0, 0}, []int{1, 1, 0, 0, 0}, []int{0, 0, 0, 1, 1}, []int{0, 0, 0, 1, 1}}
	out := 4
	fmt.Printf("in:%v out:%v my:%v\n", strings.Replace(fmt.Sprintf("%v", in), "[", "\n[", -1), out, maxAreaOfIsland(in))
}

/*
问题：求二维数组中水平或者垂直链接的1的个数
思路：1. 简单的][在每个未查找的点进行宽度优先搜索，这样每个点可能要找很多遍。如何每个点只找一遍呢？
2. 计算是，外层循环是第几行的循环，内存循环是行内第几列的循环，那么一个点计算的时候，只包含了其左边和上边的点。那么
grid[line][col]:=1
grid[line][col]+= grid[line-1][col]-  grid[line-1][col-1] if grid[line-1][col]>0
grid[line][col]+= grid[line][  col-1]-grid[line-1][col-1] if grid[line][  col-1]>0
grid[line][col]+= grid[line-1][col-1]                 if grid[line-1][col]+grid[line][col-1]
*/

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	if len(grid) > 0 {
		max = grid[0][0]
		for col := 1; col < len(grid[0]); col++ {
			if grid[0][col] > 0 {
				grid[0][col] = grid[0][col-1] + 1
			}
			if grid[0][col] > max {
				max = grid[0][col]
			}
		}
	}
	for line := 1; line < len(grid); line++ {
		if grid[line][0] > 0 {
			grid[line][0] = grid[line-1][0] + 1
		}
		if grid[line][0] > max {
			max = grid[line][0]
		}
	}
	for line := 1; line < len(grid); line++ {
		for col := 1; col < len(grid[0]); col++ {
			if grid[line][col] == 1 {
				if grid[line-1][col] > 0 {
					maxcol := col
					for ; maxcol < len(grid[0]) && grid[line-1][maxcol] > 0; maxcol++ {
					}
					if maxcol == len(grid[0]) {
						maxcol = len(grid[0]) - 1
					}
					grid[line][col] += grid[line-1][maxcol] - grid[line-1][col-1]
				}
				if grid[line][col-1] > 0 {
					grid[line][col] += grid[line][col-1] - grid[line-1][col-1]
				}
				if grid[line-1][col]+grid[line][col-1] > 0 {
					grid[line][col] += grid[line-1][col-1]
				}
				if grid[line][col] > max {
					max = grid[line][col]
				}
			}
		}
	}
	fmt.Println(strings.Replace(fmt.Sprintf("%v", grid), "[", "\n[", -1))
	return max
}
