package main

import (
	"container/list"
)

func main() {

}

/*
思路：
1. 用一个map，记录一个点是否被访问，访问了后就不在访问了。这样可以保证被每个点只访问一次。
2. bfs访问，访问数组初始为0，访问后记录到该点的面积。每一个面积都尝试更新最大值。原始点非0且访问数组为0才访问。

伪代码：
TODO: 初始化visGrid,maxArea和边界值处理
for rowIdx:=0;rowIdx<len(grid);rowIdx++ {
	for colIdx:=0;colIdx<len(grid[0]);colIdx++{
		if grid[rowIdx][colIdx]>0 && visGrid[rowIdx][colIdx]==0{
			bfs(grid,visGrid,rowIdx,colIdx,&maxArea)
		}
	}
	return maxArea
}
type point struct {
	row int
	col int
}
func bfs(grid [][]int,visGrid [][]int,bgnCol,bgnRow int,maxArea *int){
	fifo := list.New()
	// 处理第一个节点
	visit(grid,visGrid,bgnCol,bgnRow,maxArea,fifo)
	for fifo.Len()>0 {
		1. 取出头节点
		2. 对相邻四点，如果在范围内且非0：则计算并设置累加值，尝试更新最大值；并加入队列
	}
}

*/
func maxAreaOfIsland(grid [][]int) int {
	// grid[0]越界问题
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}

	maxArea := 0
	visGrid := make([][]int, len(grid))
	for idx := range visGrid {
		visGrid[idx] = make([]int, len(grid[0]))
	}
	for rowIdx := 0; rowIdx < len(grid); rowIdx++ {
		for colIdx := 0; colIdx < len(grid[0]); colIdx++ {
			if grid[rowIdx][colIdx] > 0 && visGrid[rowIdx][colIdx] == 0 {
				bfs(grid, visGrid, rowIdx, colIdx, &maxArea)
			}
		}
	}
	// fmt.Println(strings.Replace(fmt.Sprintf("%v", grid), "[", "\n[", -1))
	// fmt.Println(strings.Replace(fmt.Sprintf("%v", visGrid), "[", "\n[", -1))
	return maxArea
}

type point struct {
	row int
	col int
}

func bfs(grid, visGrid [][]int, row, col int, maxArea *int) {
	lst := list.New()
	curArea := 0
	visitFirst(row, col, grid, visGrid, point{row: row, col: col}, &curArea, maxArea, lst)
	for lst.Len() > 0 {
		// 	1. 取出头节点
		po := lst.Remove(lst.Front()).(point)
		// 	2. 对相邻四点，如果在范围内且非0：则计算并设置累加值，尝试更新最大值；并加入队列
		visitFirst(po.row, po.col-1, grid, visGrid, po, &curArea, maxArea, lst)
		visitFirst(po.row, po.col+1, grid, visGrid, po, &curArea, maxArea, lst)
		visitFirst(po.row-1, po.col, grid, visGrid, po, &curArea, maxArea, lst)
		visitFirst(po.row+1, po.col, grid, visGrid, po, &curArea, maxArea, lst)
	}
}

// 如果在范围内且非0且没访问过：则计算并设置累加值，尝试更新最大值；并加入队列
func visitFirst(x, y int, grid, visGrid [][]int, po point, curArea, maxArea *int, lst *list.List) {
	if InrangeIsoneNotvisit(x, y, grid, visGrid) {
		(*curArea)++
		visGrid[x][y] = *curArea
		if visGrid[x][y] > *maxArea {
			*maxArea = visGrid[x][y]
		}
		lst.PushBack(point{row: x, col: y})
	}
}

func InrangeIsoneNotvisit(row, col int, grid, visGrid [][]int) (rst bool) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
		return false
	}
	if grid[row][col] != 1 {
		return false
	}
	if visGrid[row][col] > 0 {
		return false
	}
	return true
}
