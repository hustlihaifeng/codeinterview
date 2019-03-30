package main

import "fmt"

func main() {
	var in [][]byte
	var out int

	in = in4
	out = 4
	fmt.Printf("%v %v expect %v\n", in, maximalSquare(in), out)
}

/*
1. 初始思路：找最大面积的题目，用广度优先搜索，只不过扩展条件变成了正方形扩展。最后发现，有c41 c42 c43 c44总共10中情况需要处理，代码复杂。很可能有重复劳动。最关键的点是，搜索过的点不能简单的不搜索了，所以用深度优先、广度优先这类方法不好。复杂度感觉很复杂，没去算。

2. 第二种思路：对每一个没有访问的点，构建一个正方形，进行左上方扩展。如果结果大于最大值，更新最大值。（因为左上方的都已经被访问，如果其他方向可以扩展的话，接下来访问的点会覆盖这些情况）。 这种直接思路可以让问题变简单，因为只有一个方向扩展了。

3. 第三中思路：在第二种的基础上，发现左上方的点都是已经被搜索过的，那么我们可以记录下每个点的左上方的最大正方形，来减少第二种方法中的递归搜索。对于一个全是1的图，复杂度为O(n^3)，第二种方法中可能是O(n^4)

4. 更进一步，如果我们记录下水平和纵向有多少个连续点，那么可以将第三种的每个点的时间变成常数级别，复杂度可以变为O(n^2)。直接做法是另外弄一个数组，记录水平和纵向在左上方向上有多少个连续点。优雅做法是直接用第三种方法中记录的每个点左上方的正方形变长就可以，因为如果能满足正方形，就一定能满足单方向连续。那么此时问题变成知道左上方点的正方形变长，知道左点和上点的连续点长度，求本点在左上方方向的能组成的最大的正方形长度：本点非0时，左上、左、上点的最小值加一（本点是1）。此时时间复杂度是O(n^2)，空间复杂那个度能变成第3步的1/4（每个点记录左上、右下点坐标，到只用记录边长（其实第三种方法也可以只记录变长，这样的话，似乎很容易就能想到第四种方法））
*/
/////////////////////////////////////////////////////////////////////////////////////////
func maximalSquare(matrix [][]byte) int {
	if maximalSquareInputInvalid(matrix) {
		return 0
	}

	colLen := len(matrix)
	rowLen := len(matrix[0])
	maxSquare := make([][]int, colLen)
	for i := 0; i < colLen; i++ {
		maxSquare[i] = make([]int, rowLen)
	}

	rst := 0
	for i := 0; i < colLen; i++ {
		for j := 0; j < rowLen; j++ {
			if matrix[i][j] == '1' {
				area := dp(xy{y: i, x: j}, maxSquare)
				if area > rst {
					rst = area
				}
			}
		}
	}

	return rst * rst
}
func dp(s xy, maxSquare [][]int) int {
	if s.x == 0 || s.y == 0 {
		maxSquare[s.y][s.x] = 1
	} else {
		maxSquare[s.y][s.x] = min(maxSquare[s.y-1][s.x-1], maxSquare[s.y][s.x-1], maxSquare[s.y-1][s.x]) + 1
	}
	return maxSquare[s.y][s.x]
}

func min(first int, nums ...int) int {
	rst := first
	for _, num := range nums {
		if num < rst {
			rst = num
		}
	}
	return rst
}

func maximalSquareInputInvalid(matrix [][]byte) bool {
	if matrix == nil {
		return true
	}
	if len(matrix) == 0 {
		return true
	}
	if len(matrix[0]) == 0 {
		return true
	}
	return false
}

type xy struct {
	x int
	y int
}

//////////////////////////////////////////////////////////////////////////////////////////////

type square struct {
	lt xy
	rb xy
}

func (s square) Area() int {
	rowNum := s.rb.x - s.lt.x + 1
	colNum := s.rb.y - s.lt.y + 1
	return rowNum * colNum
}

func maximalSquareOld(matrix [][]byte) int {
	if maximalSquareInputInvalid(matrix) {
		return 0
	}

	colLen := len(matrix)
	rowLen := len(matrix[0])
	maxSquare := make([][]*square, colLen)
	for i := 0; i < colLen; i++ {
		maxSquare[i] = make([]*square, rowLen)
	}

	rst := 0
	for i := 0; i < colLen; i++ {
		for j := 0; j < rowLen; j++ {
			if matrix[i][j] == '1' {
				area := squareDfs(xy{y: i, x: j}, maxSquare, matrix)
				if area > rst {
					rst = area
				}
			}
		}
	}

	fmt.Println(maxSquare)
	return rst
}

func squareDfs(s xy, maxSquare [][]*square, matrix [][]byte) int {
	if s.x == 0 || s.y == 0 {
		maxSquare[s.y][s.x] = &square{s, s}
		return 1
	}

	ltsquare := maxSquare[s.y-1][s.x-1]
	if ltsquare == nil {
		maxSquare[s.y][s.x] = &square{s, s}
		return 1
	}

	newSquare := square{lt: ltsquare.lt, rb: s}
	GetMaxSquare(newSquare, maxSquare, matrix)
	return maxSquare[s.y][s.x].Area()
}

func GetMaxSquare(newSquare square, maxSquare [][]*square, maxtrx [][]byte) {
	x1, y1 := newSquare.lt.x, newSquare.lt.y
	x2, y2 := newSquare.rb.x, newSquare.rb.y
	rowOffset := 0
	colOffset := 0

	for ; x2-rowOffset >= x1 && maxtrx[y2][x2-rowOffset] == '1'; rowOffset++ {
	}
	for ; y2-colOffset >= y1 && maxtrx[y2-colOffset][x2] == '1'; colOffset++ {
	}

	minOffset := rowOffset
	if colOffset < minOffset {
		minOffset = colOffset
	}

	maxSquare[y2][x2] = &square{rb: xy{x2, y2}, lt: xy{x2 - minOffset + 1, y2 - minOffset + 1}}
}
