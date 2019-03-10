
# 问题
见[https://leetcode.com/problems/sort-list/](https://leetcode.com/problems/sort-list/)
# 分析
归并排序`O(nlog(n))`的复杂度，空间复杂度的话，如果使用递归，那么至少是`O(log(n))`，使用递推的话，可以做到O(1)的空间复杂度。

# 关键点
1. 涉及到链的合并操作，使用dummy头结点来简化操作。
2. split的时候，要将两条子链斩断，也即最后设为nil，不然维护起来很复杂。
3. head结点有可能会被替换掉，此时再用head有可能出错。用`dummy.Next`则不会有问题，dummy结点不会变化，dummy后面节点的都是常规操作。
4. `split`函数设计很巧妙，是关键所在。

# 源码
见[sortlist.md](sortlist.md)