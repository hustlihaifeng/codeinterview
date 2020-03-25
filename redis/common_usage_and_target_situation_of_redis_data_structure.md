# 整理常用数据结构和使用场景
## string
### 常见用法

1. 

```shell
[lhf@VM_130_94_centos ~]$ redis-cli 
127.0.0.1:6379> SET runoobkey redis
OK
127.0.0.1:6379> get runoobkey
"redis"
127.0.0.1:6379> ttl runoobkey
(integer) -1
127.0.0.1:6379> expire runoobkey 20
(integer) 1
127.0.0.1:6379> ttl runoobkey
(integer) 17
127.0.0.1:6379> del runoobkey
(integer) 1
127.0.0.1:6379> ttl runoobkey
(integer) -2
127.0.0.1:6379> incr nu1 # 默认是数字0
(integer) 1
127.0.0.1:6379> incr nu1
(integer) 2
127.0.0.1:6379> decr nu2
(integer) -1
127.0.0.1:6379> decr nu2
(integer) -2
```
## hash
### 常见用法
1. 
```go
127.0.0.1:6379> HSET hkey a 1
(integer) 1
127.0.0.1:6379> HSET hkey a 2
(integer) 0
127.0.0.1:6379> HSET hkey b 2
(integer) 1
127.0.0.1:6379> HGETALL hkey
1) "a"
2) "2"
3) "b"
4) "2"
127.0.0.1:6379> Hkeys hkey
1) "a"
2) "b"
127.0.0.1:6379> Hvals hkey
1) "2"
2) "2"
127.0.0.1:6379> hget hkey a
"2"
127.0.0.1:6379> hdel hkey a
(integer) 1
127.0.0.1:6379> hget hkey a
(nil)
127.0.0.1:6379> del hkey
(integer) 1
127.0.0.1:6379> HGETALL hkey
(empty list or set)
```
### 应用场景

1. 常用来存储属性，对应一个map：有一个名字，然后有多个kv对。

## list

### 常见用法

1. 就是一个双端链表（head里面有一个last指针，执行tail节点，是单向链表，并不是双向链表。在尾部加元素比较省时间）

```shell
127.0.0.1:6379> LPUSH runoobkey redis
(integer) 1
127.0.0.1:6379> LPUSH runoobkey mongodb
(integer) 2
127.0.0.1:6379> LPUSH runoobkey mysql
(integer) 3
127.0.0.1:6379> LRANGE runoobkey 0 2
1) "mysql"
2) "mongodb"
3) "redis"
127.0.0.1:6379> LRANGE runoobkey 0 1
1) "mysql"
2) "mongodb"
127.0.0.1:6379> LRANGE runoobkey 1 2
1) "mongodb"
2) "redis"
127.0.0.1:6379> LRANGE runoobkey 0 10
1) "mysql"
2) "mongodb"
3) "redis"
127.0.0.1:6379> LINDEX runoobkey 0
"mysql"
127.0.0.1:6379> BLPOP runoobkey 1 # fifo B指block，没有则等待1秒。简单的是LPOP
1) "runoobkey"
2) "mysql"
127.0.0.1:6379> LRANGE runoobkey 0 10
1) "mongodb"
2) "redis"
127.0.0.1:6379> BRPOP runoobkey 1 # filo，B指block，没有则等待1秒。简单的是RPOP
1) "runoobkey"
2) "redis"
127.0.0.1:6379> LPUSH runoobkey right # 向左push
(integer) 2
127.0.0.1:6379> RPUSH runoobkey left # 向右push
(integer) 3
127.0.0.1:6379> LRANGE runoobkey 0 10
1) "right"
2) "mongodb"
3) "left"
127.0.0.1:6379> LINDEX runoobkey 0
"right"
```



### 应用场景

1. ```
   1.微博 TimeLine
   2.消息队列
   ```
## set
### set常见用法
```shell
127.0.0.1:6379> del runoobkey
(integer) 1
127.0.0.1:6379> SADD runoobkey redis
(integer) 1
127.0.0.1:6379> SADD runoobkey mongodb
(integer) 1
127.0.0.1:6379> SADD runoobkey mysql
(integer) 1
127.0.0.1:6379> SMEMBERS runoobkey
1) "mysql"
2) "redis"
3) "mongodb"
```
### 常见使用场景
1.共同好友、二度好友
2.利用唯一性，可以统计访问网站的所有独立 IP

## Sorted Set
### 常见操作
```shell
127.0.0.1:6379> zadd runoobkey 1 mysql
(integer) 1
127.0.0.1:6379> zadd runoobkey 2 redis
(integer) 1
127.0.0.1:6379> zadd runoobkey 3 mongodb
(integer) 1
127.0.0.1:6379> zadd runoobkey 3 mongodb
(integer) 0
127.0.0.1:6379> zadd runoobkey 4 mongodb
(integer) 0
127.0.0.1:6379> ZRANGE runoobkey 1 4 WITHSCORES
1) "redis"
2) "2"
3) "mongodb"
4) "4"
127.0.0.1:6379> ZRANGE runoobkey 0 4 WITHSCORES
1) "mysql"
2) "1"
3) "redis"
4) "2"
5) "mongodb"
6) "4"
```

### 常见应用场景
1.带有权重的元素，比如一个游戏的用户得分排行榜
2.比较复杂的数据结构，一般用到的场景不算太多

# 订阅发布

1. Pub/Sub 从字面上理解就是发布（Publish）与订阅（Subscribe），在 Redis 中，你可以设定对某一个 key 值进行消息发布及消息订阅，当一个 key 值上进行了消息发布后，所有订阅它的客户端都会收到相应的消息。这一功能最明显的用法就是用作实时消息系统，比如普通的即时聊天，群聊等功能。

2. 例子

   1. pub:

   ```shell
   127.0.0.1:6379> PUBLISH redisChat "Redis is a great caching technique"
   (integer) 0 # 接收到信息的订阅者数量
   ```

   2. client1

   ```shell
   127.0.0.1:6379> SUBSCRIBE redisChat
   Reading messages... (press Ctrl-C to quit)
   1) "subscribe"
   2) "redisChat"
   3) (integer) 1
   ```

   3. pub

   ```shell
   127.0.0.1:6379> PUBLISH redisChat "Learn redis by runoob.com"
   (integer) 1
   ```

   - client1

   ```shell
   1) "message"
   2) "redisChat"
   3) "Learn redis by runoob.com"
   ```

   4. client2

   ```shell
   127.0.0.1:6379> SUBSCRIBE redisChat
   Reading messages... (press Ctrl-C to quit)
   1) "subscribe"
   2) "redisChat"
   3) (integer) 1
   ```

   5. pub

   ```shell
   127.0.0.1:6379> PUBLISH redisChat "Redis is a great caching technique"
   (integer) 2 # 接收到信息的订阅者数量
   ```

   - client1

   ```shell
   1) "message"
   2) "redisChat"
   3) "Redis is a great caching technique"
   ```

   - client2

   ```shell
   1) "message"
   2) "redisChat"
   3) "Redis is a great caching technique"
   ```

   

# 参考文献

1. [Redis中5种数据结构的使用场景介绍](https://segmentfault.com/a/1190000004012214)

