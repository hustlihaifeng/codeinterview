
# 目录
- [翻转](#翻转)
- [合并](#合并)
- [两数相加](#两数相加)
- [排序](#排序)
- [相交点检测](#相交点检测)
- [多个链表合并](#多个链表合并)
- [最低公共父节点](#最低公共父节点)
- [之字形遍历二叉树](#之字形遍历二叉树)

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

# 之字形遍历二叉树
1. 关键点：有多个在变，先保持其中一个不变，可以使问题简单。
2. 思路：保持list中都是从左到右的顺序，那么从左到右或从右到左交替变换的将list里面的元素加入rst，就可以实现之字形访问。
    - 从左到右访问时，先左右后就可以，此时需要PushBack来将后添加的放到后面，保证顺序性。
    - 从右到左访问时，先右后左才能保持顺序一致性，但是此时就是反序访问。那么用PushFront将后添加的移到最前，就可以在反序一遍。
3. 伪代码:
```go
新建一个list
非空root加入list
for l.Len()>0{
    新建一个sli
    新建一个list
    if l2r {
        从左到右遍历list，将节点加入sli
        先左后右将非空子节点加入新list（PushBack，下一层是从左到右的方式）
        l2r=false
    }else{ // 链表访问方向必须和子节点添加方向一致，然后反向访问时后进先出就可以了（PushFront）
        从右到左遍历，将节点加入sli
        先右后左将非空子节点pushFront到新list（此时新list中，下一层的节点，还是从左到右的顺序）
        l2r=true
    }
    将rst加入rst
    用新list替换老list
}
```