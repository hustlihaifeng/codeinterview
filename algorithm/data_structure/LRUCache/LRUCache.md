# 问题

见leetcode 146 <https://leetcode.com/problems/lru-cache/>

# 分析

1. 全O(1)，需要使用散列函数，即map
2. LRU，使用list。最新put和get的放到list的头部。
3. 两者如何联系起来：map的value设置为list的节点。
4. 注意点：淘汰list尾部的元素时，map中的对应项要清理掉，所以需要通过节点获取key，所以节点的Value里面不仅要包含输入的value，也要包含输入的key，即一个KV结构。

# 代码

见[main.go](main.go)