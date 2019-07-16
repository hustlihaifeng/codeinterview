# 问题

1. 见leetcode [155. Min Stack](https://leetcode-cn.com/problems/min-stack/)

# 分析

1. push pop top 都是普通的栈操作，`GetMin`返回最小值。只有push pop能改变栈，其他的都是读取操作。然后由于是栈，后进先出，后面的大的值还在时获取的最小值一定是前面的已经压栈的最小值，对前面的已经存在的最小值无影响。

2. 也即找最优解，问所有情况下的最优解。由于后进先出的特性，如果一个后来的值没有更新最优解，那么该值对对最优解没有任何影响，在处理最优解部分时直接忽略改值。在某个最优解元素还在的时候，最优解是该最优解；在该最优解元素被pop的时候，需要更新到他压栈时的最优解，终点是第一个入栈的元素。

3. 关键点在于栈的后进先出特性，所以不需要考虑后来的且没有更新最优解的元素。又由于只需要得到当前还在元素的最优解，如果获取了当前还在元素的完整排序链路（如排序），那么是一种浪费。

# 代码

见[main.go](main.go)