# 问题
给定n个区间，将重叠的区间合并

# 思路
1. 按照x排序，从第一个开始，找到第一个比y大的x，那么此时左边的xl,yl. 如果yl>y，那么以yl为y向右找，直到yl==y。
那么此时x,yl是一个区间。然后就是新的x了，注意最后的情况（判断一下idx）

# 伪代码
```go
func (){
	1. 数组按照x排序
	2. 初始化x,y为0，注意判断边界
	2. 从1开始遍历排序后的数组
		if cx <= y {
			y = max(y,cy)
		}else {
			append(rst,[x,y])
			x,y = cx,cy
		}
	3. append(rst,[x,y])
	return rst
}
```

# 详见
1. [main.go](main.go)