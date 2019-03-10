# 目录

# 与事物相关的命令
1. `MULTI` 、`DISCARD` 、 `EXEC`、`WATCH`。
2. `MULTI`开启事物, `EXEC`执行事物，multi和exec之间的命令加入到队列中暂存起来，到exec时依次执行。
	- `DISCARD`可以将事物队列清空，回到非事物状态（注意到EXEC之前，事物都还没有开始执行）。
	- WATCH 命令用于在事务开始之前监视任意数量的键： 当调用 EXEC 命令执行事务时， 如果任意一个被监视的键已经被其他客户端修改了， 那么整个事务不再执行， 直接返回失败。（compare and set）
	- MULTI之后，执行MULTI或WATCH，会返回错误，但是事物队列不会清空，可以继续添加命令或者`DISCARD` 、 `EXEC`。

# ACID
1. A 原子性。redis事务在执行（EXEC）中途遇到错误，不会回滚，而是继续执行后续命令，**违反原子性**。
2. D 持久性。即使在AOF 的“总是 SYNC ”模式下，事务的每条命令在执行成功之后，会立即调用 fsync 或 fdatasync 将事务数据写入到 AOF 文件。但是，这种保存是由后台线程进行的，主线程不会阻塞直到保存成功，所以从命令执行成功到数据保存到硬盘之间，还是有一段非常小的间隔，所以这种模式下的事务也是**不持久的**。
# 参考资料
- [Redis 设计与实现](https://redisbook.readthedocs.io/en/latest/feature/transaction.html)
- [Redis事务的分析及改进](https://segmentfault.com/a/1190000002594059)