package main

func main() {

}

/*
# 问题
1. 一维有序数组，拆分成二维的，查找某个元素
# 思路
1. 先找二维，后找一维
2. 直接把二维的当做一维的来找
# 伪代码
1. 直接把二维的当做一维的来找
```go
特殊值处理
llen,clen := len(matrix),len(matrx[0]) // 长尾为0的异常
l,r := 0,llen*clen-1
for l<=r {
    m = (l+r)/2
    line:= m/clen
    col:=m%clen
    val := matrix[line][col]
    if val==target{
        return true
    }else if val > target {
        r = m -1
    }else{
        l=m+1
    }
}
return false
```
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	llen, clen := len(matrix), len(matrix[0]) // 长尾为0的异常
	l, r := 0, llen*clen-1
	for l <= r {
		m := (l + r) / 2
		line := m / clen
		col := m % clen
		val := matrix[line][col]
		if val == target {
			return true
		} else if val > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}
