# 问题

在一个数字数组中，找三个数，是只和为0

# 思路关键点

1. 在有序数组中，找两个和为指定值的数，从前后夹击，复杂度为O(n)，这是两数之和中得到的结论。所以两数之和的复杂度为O(nlog(n))
2. 那么对于三数之和，定一个数，变成两数之和问题，复杂度为O(nlog(n)+n^2)

# 伪代码
```go
sort.Ints(nums)
var rst [][]int
for i:=0;i<len(nums);i++ {
	targ := -nums[i]
	for lidx,ridx := i+1,len(nums)-1; lidx<ridx; {
		sum := nums[lidx]+nums[ridx]
		if sum == targ {
			rst = append(rst,[]int{i,lidx,ridx})
			for ; lidx<ridx && nums[lidx]==nums[lidx+1]; lidx++{} // 这里因为要去重，所以需要这个，不然直接一个外层就可以
			for ; ridx>lidx && nums[ridx]==nums[ridx-1]; ridx--{}
		}else{
			if sum > targ {
				ridx--
			}else{
				lidx++
			}
		}
	}
}
return rst
```

# 详见
1. [try20200223](try20200223/main.go)