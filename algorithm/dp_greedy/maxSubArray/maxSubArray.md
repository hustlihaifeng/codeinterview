# 问题

leetcode 59：<https://leetcode-cn.com/problems/maximum-subarray/>，求最大连续子数组之和。

# 分析

1. 有可能有负数
2. 伪代码：

```go
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
```



# 源码

1. 见[main.go](main.go)