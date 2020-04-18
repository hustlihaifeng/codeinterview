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

### 3.2.1 加锁过程

1. 需要保证setnx和expire的原子性

   1. 使用Lua脚本（包含setnx和expire两条指令）

   2. ### set key value [EX seconds][PX milliseconds][NX|XX] 命令，Redis在 2.6.12 版本开始，为 SET 命令增加一系列选项：

      1. EX seconds: 设定过期时间，单位为秒
      2. PX milliseconds: 设定过期时间，单位为毫秒
      3. NX: 仅当key不存在时设置值
      4. XX: 仅当key存在时设置值

### 3.2.2 释放锁

1. 要判断这个锁是不是自己锁的，不是自己锁的不能释放。这里使用Lua脚本的方式，尽量保证原子性。

```java
public boolean releaseLock_with_lua(String key,String value) {
    String luaScript = "if redis.call('get',KEYS[1]) == ARGV[1] then " +
            "return redis.call('del',KEYS[1]) else return 0 end";
    return jedis.eval(luaScript, Collections.singletonList(key), Collections.singletonList(value)).equals(1L);
}
```

### 3.3.3 分布式锁加锁超时怎么办

1. 要把ttl设置为大于加锁时间，然后所有的节点的加锁过程用同一套实现。那这样就能避免在正常情况下的锁超时。
2. 如果不同点的ttl不一样，那么恶意节点获取锁后一直不释放，那么其他节点就永远获取不到锁了。




# 4. 千万并发下的分布式锁实现

1. 按照id进行一致性哈希分片
2. 单片里面多服务，每个分片里面获取自己的分布式锁。并将分布式锁对应的资源也分片。
3. 千万并发不一定要获取一千万次分布式锁：如果是电商秒杀，可以直接不用分布式锁，用MQ来进行排队，MQ大小设置为商品库存大小，MQ满了就让接下来的请求返回失败。那么此时因为有了队列，分布式锁的并发量可以大大降低。

# 参考资料

1. [再有人问你分布式锁，这篇文章扔给他](<https://juejin.im/post/5bbb0d8df265da0abd3533a5>)

2. [Redlock（redis分布式锁）原理分析](https://www.cnblogs.com/rgcLOVEyaya/p/RGC_LOVE_YAYA_1003days.html)