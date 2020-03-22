# 问题

见leetcode 146 <https://leetcode.com/problems/lru-cache/>
1. 设计一种LRU（least resently used）算法：最近访问的最新，
2. 所有的都是O(1)的时间复杂度

# 思路
1. 在定位好了的情况下，链表的插入、更新、重排都是O(1)操作。（其他的有么：map不能保序，数组重排不是，tree保持结构一般不是）
2. list的关键在于如何在O(1)的时间定位，一把用map。本例中国给定key，找val。可以key作为map的key，val存储在list中。

# 伪代码
```go
# func put(key,val){
	看map中有没有key,没有则checkAdd，有得到node
	更新node的值为val
	将node放到链头
}
## func checkAdd(key,val){
	if 链满了 {
		del链尾 // 删除node，map中node.key置空，长度减一（如果有list.Len()则不需要手动减）
	}
	add到链头
}

# func get(key) {
	map中没有则返回，有则得到node
	将node移动到链头
	返回node.val
}

# func add() {
	val加入链头且长度加一，map中加入key,node
}
# func del() {
	删除node，map中node.key置空，长度减一（如果有list.Len()则不需要手动减）
}
```

# 注意由于需要通过链尾得到map的key，所以node中需要存储key

# 代码

见[main.go](try20200311/main.go)