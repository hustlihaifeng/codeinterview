package main

func main() {

}

/*
# 问题
1. 找二维数组中最大正方形

# 思路
1. 递推，O(n^2)的时空复杂度，每一个点，存其左上部分的正方形长度，列长度，行长度
2. 可以优化成O(n)的空间复杂度，只需要保存上一行的就可以了（实际上是本行的左边和上一行的当前位置和右边）,row-1,col-1的存在
所以需要两行

# 伪代码
type len3 struct{
	len int
	collen int
	rowlen int
}
new := &(make([]len3,len(matrix[0]),len(matrix[0])))
pre := &(make([]len3,len(matrix[0]),len(matrix[0])))
for row:=0;row<len(matrix);row++{
	for col:=0;col<len(mtriax[0]);col++{
		if matrix[row][col]=1{
			if row==0 || col == 0 {
				(*new)[col].len=1
				(*new)[col].rowlen=1
				(*new)[col].collen=1
			}else{
				(*new)[col].len=max((*pre)[col].rowlen, (*pre)[col-1].len, (*new)[col-1].collen)+1
				(*new)[col].rowlen=(*pre)[col].rowlen+1
				(*new)[col].collen=(*new)[col-1].colen+1
			}
		}else{
			(*new)[col].len=0
			(*new)[col].rowlen=0
			(*new)[col].collen=0
		}
	}
	*pre,*new = *new,*pre
}
*/
func maximalSquare(matrix [][]byte) int {
	type len3 struct {
		len    int
		collen int
		rowlen int
	}
	if len(matrix) == 0 {
		return 0
	}
	news := (make([]len3, len(matrix[0]), len(matrix[0])))
	pres := (make([]len3, len(matrix[0]), len(matrix[0])))
	new, pre := &news, &pres
	max := 0
	// fmt.Println(strings.Replace(fmt.Sprintf("%v", matrix), "[", "\n[", -1))
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == '1' {
				if row == 0 || col == 0 {
					(*new)[col].len = 1
					(*new)[col].rowlen = 1
					(*new)[col].collen = 1

				} else {
					(*new)[col].len = min3((*pre)[col].rowlen, (*pre)[col-1].len, (*new)[col-1].collen) + 1
					(*new)[col].rowlen = (*pre)[col].rowlen + 1
					(*new)[col].collen = (*new)[col-1].collen + 1
				}
			} else {
				(*new)[col].len = 0
				(*new)[col].rowlen = 0
				(*new)[col].collen = 0
			}
			if (*new)[col].len > max {
				max = (*new)[col].len
			}
		}
		// fmt.Printf("row:%v pre:%v new:%v\n", row, *pre, *new)
		*pre, *new = *new, *pre
		// fmt.Printf("after pre:%v new:%v\n", *pre, *new)
	}
	return max * max
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
