package main

func main() {

}

/*
# 问题
1. 找买卖股票的最佳时机

# 思路
1. 因为要先买后卖，所以不能直接用最大值减去最小值，因为最大值可能在最小值前面
2. 暴力解法：已每个点为买入点，后面的每个点为卖出点。O(n^2)的复杂度
3. 计算每一天的增加值，如果累计增加值大于0，那么目前依然是对整体有利的；如果小于0，那么从新开始比较好。每一次都尝试更新最大值。
O(n)的时间复杂度。

# 伪代码
if len(prices)<2{
	return 0
}
max := 0
cur := 0
for i:=1;i<len(prices);i++ {
	offset := prices[i]-prices[i-1]
	cur += offset
	if cur > max {
		max = cur
	}
	if cur < 0 {
		cur = 0
	}
}
*/
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := 0
	cur := 0
	for i := 1; i < len(prices); i++ {
		offset := prices[i] - prices[i-1]
		cur += offset
		if cur > max {
			max = cur
		}
		if cur < 0 {
			cur = 0
		}
	}
	return max
}
