# 问题
见[https://leetcode.com/problems/longest-consecutive-sequence/](https://leetcode.com/problems/longest-consecutive-sequence/)
# 分析
hash map读取都是`O(1)`的操作，那么可以把数组中元素都唯一的加入到map中，然后取出元素。从map中找其相邻元素，没找到则不连续；找到则删除该key。每个连续块会有首尾两个空位，但总体次数不超过元素个数的三倍。所以是O(n). 

**关键点**: **map的增删改查都可以看做是O(1)的时间复杂度**。

# 伪代码
```go
for 数组里面的每一个元素{
    覆盖式的加入 map，值设为 true
}

最大计数设置为0
for map中的每一对kv元祖 {
    计数设为0

    计数加一并从 map中删除k
    khigh,klow设为k
    for ++khigh在map中存在 {
        计数加一并从map中删除khigh
    }

    for --klow在map中存在 {
        计数加一并从 map中删除klow
    }

    if 计数大于最大计数{
        最大计数设为当前计数
    }
}
```
# 代码
见[main.go](main.go)
