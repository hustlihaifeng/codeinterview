
# 目录
- [翻转](#翻转)
- [合并](#合并)
- [两数相加](#两数相加)
- [排序](#排序)
- [相交点检测](#相交点检测)
- [多个链表合并](#多个链表合并)
- [最低公共父节点](#最低公共父节点)

# 翻转
1. 关键点：将`newHead`初始化为nil，新链表从前往后涨。
2. 代码见[reverse/main.go](reverse/main.go)

# 合并
1. 关键点
- 头结点为空会让代码变得复杂，我们可以新建一个`dummy`头结点，那么此后头结点都不为空，结果返回`dummy.Next`.
- 循环结束后，可以直接让新链尾指向剩下的非空链

2. 代码见[merge/main.go](merge/main.go)

# 两数相加
1. 可以用来实现大数相加
2. 关键点
- `dummy`虚拟头避免判断链头是否为空
- 最后一个进位需要加上
- 为了避免两个都为不为空阶段、一个不为空阶段、进位不为空阶段这三段判断的重复代码。可以把循环条件设为`l1!=nil || l2!=nil || add!=0`

3. 代码见:[addtwonum/main.go](addtwonum/main.go)

# 排序
## 排序分析
归并排序`O(nlog(n))`的复杂度，空间复杂度的话，如果使用递归，那么至少是`O(log(n))`，使用递推的话，可以做到O(1)的空间复杂度。

## 排序关键点
1. 涉及到链的合并操作，使用dummy头结点来简化操作。
2. split的时候，要将两条子链斩断，也即最后设为nil，不然维护起来很复杂。
3. head结点有可能会被替换掉，此时再用head有可能出错。用`dummy.Next`则不会有问题，dummy结点不会变化，dummy后面节点的都是常规操作。
4. `split`函数设计很巧妙，是关键所在。

## 排序源码
见[sortlist/main.go](sortlist/main.go)

# 相交点检测
1. 方法1： 把headA的尾指针，指向headB。那么如果有交点，两者成环，没有交点两者不成环。实现上，先把两个链表连起来，然后调用环点检测函数。空间复杂度O(1),时间复杂度6m+5n=k+2k+k+k+m=5k+m=5(m+n)+m=6m+5n
2. 方法2：两个指针分别从a、b链头开始遍历，遇到nil后就指向另一条链。那么最后，两个指针走的长度相等。第一次相遇时，要么在交叉点，要么在nil点。nil点说明非相交，交叉点说明相交。时间复杂度2m+2n，空间复杂度O(1)。伪代码：
```go
pa=headA
pb=headB
for pa!=pb {
    pa = (pa==nil)?headB:pa.Next
    pb = (pb==nil)?headA:pb.Next
}
return pa
```
3. 代码见 [intersection_point/main.go](intersection_point/main.go)

# 多个链表合并
1. 将链表数组二分，分到两个时进行合并,每个需要合并log(k)次（k指链表个数），所以是O（nlog(k))的复杂度。 
2. 伪代码：
```go
if low==high{
    return sli[low]
}
if low+1==high{
    return merge2(sli[low],sli[high])
}
mid := (low+high)/2
lowSli := mergeKLists(low,middle-1)
highSli := mergeKLists(middle,high)
return merge2(lowSli,highSli)
```

3. 代码详见：[mergeklists/main.go](mergeklists/main.go)

# 最低公共父节点
1. 思路，先序遍历，得到每个节点的父节点；然后就是找两个链表的交点。先序遍历O(n)的复杂度，找两个链表交点，O(height)的复杂度，总体O(n)
2. 牛人思路：如果pq分布在左右子树中，那么就是当前节点；否则返回其中有的那个子树。没有的，返回nil。那么pq的非LCS点会返回pq本身，LCS点处会返回改点，LCS点的父节点，会返回LCS点。没有的，返回nil，这个如何保证呢？首先是找到nil点还没找到，那么返回nil；其次是两个子节点都没有，返回nil。这也是一个深度优先遍历，时空复杂度都是O(n)。 **关键是，遍历之后，这个子树有没有p或者q就可以知道，知道了的话，第一个左右同时出现p、q的点，就是LCS点。在一个时另一个的ancestor的情况，通过其他分支都没有另一个点，来得出是这个点的子节点**
3. 代码见：[lowest_common_ancestor/main.go](lowest_common_ancestor/main.go)，代码非常简单。牛人思路啊。