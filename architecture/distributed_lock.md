[TOC]



# 1. 分布式锁是什么

1. 分布式环境中的锁

# 2. 好的分布式锁要解决哪些问题

1. 互斥性

- 分布式锁需要保证在不同节点的不同线程的互斥

2. 可重入性

- 同一个节点上的同一个线程如果获取了锁之后那么也可以再次获取这个锁

3. 锁超时

- 和本地锁一样支持锁超时，防止死锁

4. 高效，高可用

- 加锁和解锁需要高效，同时也需要保证高可用防止分布式锁失效

5. 支持阻塞和非阻塞

- 支持`lock`和`rylock`以及`tryLock(long timeOut)`

# 3. 常见分布式锁方案对比

## 3.1 Mysql分布式锁



## 3.2 RedLock

1. <https://juejin.im/post/5bbb0d8df265da0abd3533a5>
2. <https://www.cnblogs.com/rgcLOVEyaya/p/RGC_LOVE_YAYA_1003days.html>



# 参考资料

1. [再有人问你分布式锁，这篇文章扔给他](<https://juejin.im/post/5bbb0d8df265da0abd3533a5>)

2. [Redlock（redis分布式锁）原理分析](https://www.cnblogs.com/rgcLOVEyaya/p/RGC_LOVE_YAYA_1003days.html)