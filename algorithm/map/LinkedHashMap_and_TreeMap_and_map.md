1. Java中HashMap的实现是hash函数、取模、链表解决冲突。
2. Java中LinkedHashMap在HashMap的基础上，维护了一个所有节点的双向链表（添加了after、before两个指针），来存储总体的顺序。双向链表里面维持LRU，链表维持LRU的问题是直接查找时O(n)的复杂度，但是结合了HashMap可以将查找复杂度控制在O(1)，双向链表也使得链表的维护变得相对简单。**所以实现了O(1)的LRU算法**！！！

