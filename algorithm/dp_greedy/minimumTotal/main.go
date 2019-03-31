package main

func main() {

}

/*
分析：三角形可以看做一棵树，每个点有很多子树，如果每个子树的最小路径清楚了，那么当前点的最短路径就是当前点加上子树中的最短路径。可以采用递归，想要节省空间采用递推，从下往上递推。由于每一层的空间在计算完其上一层后就不会再使用，所以空间复杂度可以限制在最下面一层的长度上，也即the total number of rows in the triangle。
伪代码：
申请一个层数长度的数组rst，初始化为三角形最下面一层的值
for 从倒数第二层开始向上的每一层 {
    for 当前层的每一个元素{
        rst[i]=sli[i]+min(rst[i],rst[i+1]) // 在计算完当前层的i后，下一层的i就可以不使用了，所以可以被覆盖
    }
}
return rst[0]
*/
func minimumTotal(triangle [][]int) int {
	if minimumTotalInputInvalid(triangle) {
		return 0
	}

	length := len(triangle)
	rst := make([]int, length)
	for i := 0; i < length; i++ {
		rst[i] = triangle[length-1][i]
	}

	for col := length - 2; col >= 0; col-- {
		sli := triangle[col]
		for i, val := range sli {
			rst[i] = val + min(rst[i], rst[i+1])
		}
	}

	return rst[0]
}

func minimumTotalInputInvalid(triangle [][]int) bool {
	if triangle == nil {
		return true
	}
	if len(triangle) == 0 {
		return true
	}
	return false
}
func min(first int, sli ...int) int {
	rst := first
	for _, val := range sli {
		if val < rst {
			rst = val
		}
	}
	return rst
}
