# 目录
- [MyISAM和InnoDB对比](#MyISAM和InnoDB对比)
- [参考资料](#参考资料)


# MyISAM和InnoDB对比

|对比项|MyISAM|InnoDB|
|-|-|-|
|事物|不支持|支持|
|外键|不支持|支持|
|锁|表级锁|行级锁、表级锁等等|
|应用场景|非事物场景，读多写少，有大量的select，MyISAM能提供高速检索和存储的能力|1、事物场景 2、有大量的insert、update，InnoDB行级锁能提高多用户并发操作性能|
|索引结构|B+树，所有索引叶子节点都是存的索引值+数据物理地址指针。索引文件`.MYI`，数据文件`.MYD`|[B+树](../algorithm/tree/01%20平衡二分查找树、B树、B+树、R树.md),聚集索引的叶子节点存数据，其他索引叶子节点存聚集索引的索引值。聚集索引有主键则主键，没主键则唯一标识列，没唯一标识列则InnoDB会加一个6字节的ID字段|

# 参考资料
- [MyISAM和InnoDB的主要区别和应用场景](https://blog.csdn.net/u013408431/article/details/71270464)
- [平衡二叉树、B树、B+树、B*树 理解其中一种你就都明白了](https://zhuanlan.zhihu.com/p/27700617)
- [从B树、B+树、B*树谈到R 树](https://blog.csdn.net/v_JULY_v/article/details/6530142/)
