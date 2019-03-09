
# 目录
- [翻转](#翻转)
- [合并](#合并)

# 翻转
1. 关键点：将`newHead`初始化为nil，新链表从前往后涨。
2. 代码见[reverse/main.go](reverse/main.go)

# 合并
1. 关键点
- 头结点为空会让代码变得复杂，我们可以新建一个`dummy`头结点，那么此后头结点都不为空，结果返回`dummy.Next`.
- 循环结束后，可以直接让新链尾指向剩下的非空链

2. 代码见[merge/main.go](merge/main.go)
