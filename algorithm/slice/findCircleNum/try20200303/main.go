package main

import (
	"fmt"
)

func main() {
	fmt.Printf("in:%v out:%v my:%v\n", in, out, findCircleNum(in))
}

/*
# 问题
1. 数组中为1的边表示朋友关系，找有多少个朋友圈子

# 思路
1. 与岛屿面积不同的是，这个题的二维数组是无向邻接表（对称的）。
2. 邻接表不只有上下左右四种可能
3. 领接表的bfs或者dfs需要从一个个断点为起点找
*/
func findCircleNum(M [][]int) (rst int) {
	visited := make([]int, len(M), len(M))
	for i := 0; i < len(M); i++ {
		if visited[i] == 0 {
			dfs(M, visited, i)
			rst++
		}
	}
	return rst
}
func dfs(M [][]int, visited []int, row int) {
	visited[row] = 1
	for col := 0; col < len(M[0]); col++ {
		if M[row][col] == 1 && visited[col] == 0 {
			dfs(M, visited, col)
		}
	}
}
