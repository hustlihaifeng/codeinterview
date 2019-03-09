
# 目录
- [翻转](#翻转)
- [合并](#合并)
- [两数相加](#两数相加)

# 翻转
1. 关键点：将`newHead`初始化为nil，新链表从前往后涨。
2. 代码见[reverse/main.go](reverse/main.go)

# 合并
1. 关键点：头结点为空会让代码变得复杂，我们可以新建一个`dummy`头结点，那么此后头结点都不为空，结果返回`dummy.Next`.
2. 代码见[merge/main.go](merge/main.go)

# 两数相加
1. 可以用来实现大数相加
2. 关键点
- `dummy`虚拟头避免判断链头是否为空
- 最后一个进位需要加上
- 为了避免两个都为不为空阶段、一个不为空阶段、进位不为空阶段这三段判断的重复代码。可以把循环条件设为`l1!=nil || l2!=nil || add!=0`
3. 代码见:[addtwonum/main.go](addtwonum/main.go)