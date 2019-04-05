# 问题
见[https://leetcode.com/problems/longest-increasing-subsequence/description/](https://leetcode.com/problems/longest-increasing-subsequence/description/), 找数组中的递增数组的最大长度
# 分析
## 暴力破解：
对每个点，后面所有的大点中，结果最大的一个加1。没有大点，则是1。这是一个递推的过程。O(n^2)。
反过来的话，就是前面所有的小点的结果中的最大值加1，没有小点则是1.O(n^2)
## 优化：
1. 如果将前面的点排序，那么排序复杂度是O(nlog(n))，这样去找每个点好像依然是O(n)。
2. 记录下长度为i的序列中，最小尾值，那么这些尾值应该是递增的（不是的话用后面的序列的前一个值代替前一个长度的尾值即可），此时找比当前点小的那个点，其长度加一即可。关键点事尾值如何更新。当前得到一个值，与已有的尾值进行比较即可。
## 伪代码
```go
minValueOfLenI[0]设置为数组里面的第一个值
for 数组里面的每一个值 {
    在minValueOfLenI里面找到第一个小于等于当前值的值的下标i
    if 没找到 { // 当前值是最小值
        minValueOfLenI[0]设置为当前值
    } else if 所找到的值等于当前值 {
        忽略
    } else { // 找到比当前值小的值
        if i == 当前最大下标 {
            当前值append到minValueOfLenI上 // 长度加一
        }else {
            if 当前值 < minValueOfLenI[i+1] {
                minValueOfLenI[i+1] 设置为当前值
            }
        }
    }

}
```
# 经验
1. 先想出暴力破解法，不需要考虑时空复杂度，然后一步步优化。否则自己容易迷糊。
# 源码
见[main.go](main.go)