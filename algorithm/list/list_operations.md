
# 目录
- [翻转](#翻转)
- [合并](#合并)
- [两数相加](#两数相加)
- [排序](#排序)

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