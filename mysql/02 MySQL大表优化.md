本文主要是对 [MySQL大表优化方案](https://segmentfault.com/a/1190000006158186)的学习总结补充。
# 1 分析

## 1.1 单表优化
1. 单表性能上限
> 一般以整型值为主的表在千万级以下，字符串为主的表在五百万以下是没有太大问题的 

### 1.1.1 字段
- 尽量使用`TINYINT`、`SMALLINT`、`MEDIUMINT`作为整数类型而非`INT`(分别是1,2,3,4字节)，如果非负则加上`UNSIGNED`。详见：[https://dev.mysql.com/doc/refman/8.0/en/integer-types.html](https://dev.mysql.com/doc/refman/8.0/en/integer-types.html)
- `VARCHAR`的长度只分配真正需要的空间
- 使用枚举或整数代替字符串类型.([枚举](https://dev.mysql.com/doc/refman/8.0/en/enum.html)：会自动把字符串形式的枚举值转化为int型，节省空间)
- 尽量使用`TIMESTAMP`而非`DATETIME`（4字节比8字节，详见[https://stackoverflow.com/a/20718367](https://stackoverflow.com/a/20718367)）
- 避免使用NULL，给个默认值。详见：[MySQL，一千个不用 Null 的理由](https://www.itcodemonkey.com/article/1405.html)、[MySQL中NULL对索引的影响](https://www.jianshu.com/p/3cae3e364946)
### 1.1.2 索引
1. 一些概念
- 聚集索引：InnoDB中叶子节点中存有数据的索引
- 二级索引：InnoDB中出聚集索引之外的索引
- 联合索引、复合索引：多列索引。多列索引(a,b,c)中，按a、a,b、a,b,c来筛选都是有效的，没有a的比如b、c、b,c都是无效的。
- [前缀索引](http://www.cnblogs.com/studyzy/p/4310653.html)：对文本的前几个字符建立索引：`KEY(column_name(prefix_length))`
- [全文索引](http://mysql.taobao.org/monthly/2015/10/01/)：是一种通过建立倒排索引，快速匹配文档的方式：`fulltext(b)`，[5.6之后可以支持MyISAM和InnoDB，5.6之前只支持MyISAM](https://dev.mysql.com/doc/refman/5.5/en/fulltext-restrictions.html)。
- [覆盖索引](https://yq.aliyun.com/articles/62419): 指一个查询语句的执行只需要从辅助索引中就可以得到查询记录，而不需要查询聚集索引中的记录
2. 索引优化
- 考虑在`WHERE`和`ORDER BY`命令上涉及的列建立索引
- [null列是可以用到索引的，不管是单列索引还是联合索引，但仅限于is null，is not null是不走索引的](https://www.jianshu.com/p/3cae3e364946)
- 值分布很稀少的字段不适合建索引，例如"性别"这种只有两三个值的字段
- 字符字段只建[前缀索引](http://www.cnblogs.com/studyzy/p/4310653.html)
- 字符字段最好不要做主键
  - 聚集索引占用空间越小，在二级索引中叶子节点占用空间就越小，就能cache更多索引
  - 随机长度字符串主键会产生碎片
- 不用外键，由程序保证约束
- 使用多列索引时主意顺序和查询条件保持一致，同时删除不必要的单列索引
### 1.1.3 查询SQL

- 可通过开启慢查询日志来找出较慢的SQL
- 不做列运算：SELECT id WHERE age + 1 = 10，**任何对列的操作都将导致表扫描**，它包括数据库教程函数、计算表达式等等，**查询时要尽可能将操作移至等号右边**
- sql语句尽可能简单：一条sql只能在一个cpu运算；大语句拆小语句，减少锁时间；一条大sql可以堵死整个库
- **OR改写成IN：OR的效率是n级别，IN的效率是log(n)级别**，in的个数建议控制在200以内
- 不用函数和触发器，在应用程序实现
- 避免%xxx式查询
- 少用JOIN
- 不用`SELECT *`
- 使用同类型进行比较，比如用'123'和'123'比，123和123比
- **尽量避免在WHERE子句中使用!=或<>操作符，否则将引擎放弃使用索引而进行全表扫描**
- **对于连续数值，使用BETWEEN不用IN**：SELECT id FROM t WHERE num BETWEEN 1 AND 5
- 列表数据不要拿全表，要使用LIMIT来分页，每页数量也不要太大
### 1.1.4 引擎
见 [01 InnoDB和MyISAM.md](01 InnoDB和MyISAM.md)
## 1.2 读写分离

从库读主库写，一般不要采用双主或多主引入很多复杂性.
## 1.4 表分区
1. **对用户来说，分区表是一个独立的逻辑表**，但是底层由多个物理子表组成,**MySQL实现分区的方式也意味着索引也是按照分区的子表定义，没有全局索引**.
2. **用户的SQL语句是需要针对分区表做优化，SQL条件中要带上分区条件的列，从而使查询定位到少量的分区上，否则就会扫描全部分区**，可以通过EXPLAIN PARTITIONS来查看某条SQL语句会落在那些分区上，从而进行SQL优化
3. 分区的好处
- 可以让单表存储更多的数据
- 可以通过清除整个分区批量删除大量数据，也可以增加新的分区来支持新插入的数据
- 部分查询能够从查询条件确定只落在少数分区上，速度会很快
- 分区表的数据还可以分布在不同的物理设备上，从而搞笑利用多个硬件设备
- 可以使用分区表赖避免某些特殊瓶颈，例如InnoDB单个索引的互斥访问、ext3文件系统的inode锁竞争
- 可以备份和恢复单个分区
4. 分区的缺点
- 一个表最多只能有1024个分区
- 如果分区字段中有主键或者唯一索引的列，那么所有主键列和唯一索引列都必须包含进来
- 分区表无法使用外键约束
- NULL值会使分区过滤无效
- 所有分区必须使用相同的存储引擎
5. 分区适合的场景有
- 最适合的场景数据的时间序列性比较强，则可以按时间来分区，如下所示：
```sql
CREATE TABLE members (
    firstname VARCHAR(25) NOT NULL,
    lastname VARCHAR(25) NOT NULL,
    username VARCHAR(16) NOT NULL,
    email VARCHAR(35),
    joined DATE NOT NULL
)
PARTITION BY RANGE( YEAR(joined) ) (
    PARTITION p0 VALUES LESS THAN (1960),
    PARTITION p1 VALUES LESS THAN (1970),
    PARTITION p2 VALUES LESS THAN (1980),
    PARTITION p3 VALUES LESS THAN (1990),
    PARTITION p4 VALUES LESS THAN MAXVALUE
);
```
查询时加上时间范围条件效率会非常高，同时对于不需要的历史数据能很容的批量删除。
- 如果数据有明显的热点，而且除了这部分数据，其他数据很少被访问到，那么可以将热点数据单独放在一个分区，让这个分区的数据能够有机会都缓存在内存中，查询时只访问一个很小的分区表，能够有效使用索引和缓存
## 1.5 垂直拆分

1. 垂直分库是根据数据库里面的数据表的相关性进行拆分，比如：一个数据库里面既存在用户数据，又存在订单数据，那么垂直拆分可以把用户数据放到用户库、把订单数据放到订单库。垂直分表是对数据表进行垂直拆分的一种方式，常见的是把一个多字段的大表按常用字段和非常用字段进行拆分，每个表里面的数据记录数一般情况下是相同的，只是字段不一样，使用主键关联
2. 优点
- 可以使得行数据变小，一个数据块(Block)就能存放更多的数据，在查询时就会减少I/O次数(每次查询时读取的Block 就少)
- 可以达到最大化利用Cache的目的，具体在垂直拆分的时候可以将不常变的字段放一起，将经常改变的放一起
3. 缺点
- 会引起表连接JOIN操作（增加CPU开销）可以通过在业务服务器上进行join来减少数据库压力
- 依然存在单表数据量过大的问题（需要水平拆分）
- 事务处理复杂 
## 1.6 水平拆分
1. 水平拆分是通过某种策略**将数据分片来存储**（比如按id hash），分库内分表和分库两部分，每片数据会分散到不同的MySQL表或库，达到分布式的效果，能够支持非常大的数据量。前面的表分区本质上也是一种特殊的库内分表
2. 水平拆分的优点
- 不存在单库大数据和高并发的性能瓶颈
- 应用端改造较少
- 提高了系统的稳定性和负载能力
3. 水平拆分缺点
- 分片事务一致性难以解决
- 跨节点Join性能差，逻辑复杂
- 数据多次扩展难度跟维护量极大
4. 分片原则
- 能不分就不分
- 分片数量尽量少，分片尽量均匀分布在多个数据结点上，因为一个查询SQL跨分片越多，则总体性能越差，虽然要好于所有数据在一个分片的结果，只在必要的时候进行扩容，增加分片数量
- 尽量不要在一个事务中的SQL跨越多个分片，分布式事务一直是个不好处理的问题

# 参考资料
- [MySQL大表优化方案](https://segmentfault.com/a/1190000006158186)
- [MySQL · 引擎特性 · InnoDB 全文索引简介](http://mysql.taobao.org/monthly/2015/10/01/)
- [https://dev.mysql.com/doc/refman/8.0/en/integer-types.html](https://dev.mysql.com/doc/refman/8.0/en/integer-types.html)
- [https://dev.mysql.com/doc/refman/8.0/en/enum.html](https://dev.mysql.com/doc/refman/8.0/en/enum.html)
