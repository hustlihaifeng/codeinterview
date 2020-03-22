# 问题

1. 实现一个数据结构支持以下操作：

   1. Inc(key) - 插入一个新的值为 1 的 key。或者使一个存在的 key 增加一，保证 key 不为空字符串。
   2. Dec(key) - 如果这个 key 的值是 1，那么把他从数据结构中移除掉。否者使一个存在的 key 值减一。如果这个 key 不存在，这个函数不做任何事情。key 保证不为空字符串。
   3. GetMaxKey() - 返回 key 中值最大的任意一个。如果没有元素存在，返回一个空字符串`""`。
   4. GetMinKey() - 返回 key 中值最小的任意一个。如果没有元素存在，返回一个空字符串`""`。

   挑战：以 O(1) 的时间复杂度实现所有操作。

2. 详见：<https://leetcode-cn.com/explore/interview/card/bytedance/245/data-structure/1033/>

# 分析

1. Inc和Dec用hash即可实现`O(1)`.
2. `GetMaxKey()`和`GetMinKey()`则需要有一个结构，保存最大最小值。而且在Inc和Dec的时候还要以O(1)的时间变化。如果用一个链表，将key按照大小链接起来，获取最大最小值是O(1)，哈希（key是string，value是链表节点）里面只存链表节点，那么对节点的索引也是O(1)；在Dec/Inc时，如果能以O(1)的复杂度相同大小的收尾元素，将变化的元素与之交换位置，就可以实现全部O(1)。那么在加一个哈希，记录各个长度的链表的的首元素（key是长度，value是链表节点），就可以实现定位。
3. 整体上看，就是一个有序二维链表加上两个哈希，一个哈希指向每个元素，一个哈希指向值相同的子链的首元素。

# 关键点

```go
type AllOne struct{
    KeyElem map[string] list.Element
    LengthFirst []list.Element
    Elems list.List
}
```

1. Inc

```go
func (this *AllOne) Inc(key string) {
    if key 在 KeyElem中不存在 {
        新建一个长度为1的节点NewElem，加入成为ValueElem[1]的前置节点，将ValueElem[1]设置为新加入的节点
    } else {
        将当前节点和下一长度值（可能需要循环）的头结点的前置节点进行替换，并将下一长度值链表的头结点更换为当前节点。
        如果当前是头节点，将当前长度的值的头节点置空
        当前节点值加1
    }
}
```

2. Dec

```go
func (this *AllOne) Dec(key string) {
    if key 在 KeyElem中不存在 {
        返回
    } else {
        if 当前节点值为1 {
            删除当前节点
            如果是1的头节点，将1的头结点置空
        }else{
            删除当前节点，如果是头节点，将二维中的当前值头结点置空
            将当前节点插入成为当前值减一的头节点
        }
    }
}
```





# 代码

