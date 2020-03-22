package main

func main() {

}

/*
# 问题
1. 找平方根

# 思路
1. 暴力找，从0到x遍历，找到相邻的两个。
2. 因为有序，所以可以二分查找。

# 伪代码
l,r := 0,x
for l < r {
    m = (l+r)/2
    mm = m*m
    mm1= (m+1)*(m+1)
    if mm < x && mm1 >= x {
        return m+1
    }else if mm1 < x {
        l = m+1
    }else {
        r = m-1
    }
}
return l
*/

func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		m := (l + r) / 2
		mm := m * m
		mm1 := (m + 1) * (m + 1)
		if mm < x && mm1 >= x {
			if mm1 == x {
				return m + 1
			}
			return m
		} else if mm1 < x {
			l = m + 1
		} else {
			r = m // 因为在l<r时，m偏向于l，那么此时r一定会变小
		}
	}
	return l
}
