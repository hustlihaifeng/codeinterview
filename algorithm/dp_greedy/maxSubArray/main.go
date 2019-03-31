package main

func main() {

}

/*
sum、maxSum设置为第一个元素，起点终点设置为0
for 数组里面从第2个元素开始的没一个元素 {
    if sum小于0，那么加上sum只会使后面的更小{
        当前查找的起点更新为当前索引
        sum=当前点的值
    }else{
        sum加上当前点的值
    }
    if sum > maxSum {
        更新maxSum为当前sum
        maxSum对应的起点设置为当前的起点
        maxSum对应的终点设置为当前点
    }
}
*/
func maxSubArray(nums []int) int {
	if maxSubArrayInputInvalid(nums) {
		return 0
	}

	maxSum, sum := nums[0], nums[0]
	//maxBgnIdx,maxEndIdx := 0,0
	//curBgnIdx := 0
	length := len(nums)
	for i := 1; i < length; i++ {
		if sum < 0 {
			//curBgnIdx=i
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		if sum > maxSum {
			maxSum = sum
			//maxBgnIdx=curBgnIdx
			//maxEndIdx=i
		}
	}
	return maxSum
}
func maxSubArrayInputInvalid(nums []int) bool {
	if nums == nil {
		return true
	}
	if len(nums) == 0 {
		return true
	}
	return false
}
