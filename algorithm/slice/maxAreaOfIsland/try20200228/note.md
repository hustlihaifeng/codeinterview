# 问题
1. 求二维数组表示的岛屿的最大面积

# 思路
1. 用一个map，记录一个点是否被访问，访问了后就不在访问了。这样可以保证被每个点只访问一次。
2. bfs访问，访问数组初始为0，访问后记录到该点的面积。每一个面积都尝试更新最大值。原始点非0且访问数组为0才访问。

# 伪代码
```go
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
```

# 详见
1. [main.go](main.go)