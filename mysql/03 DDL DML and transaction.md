# 目录
- [原理](#原理)
- [delete truncate drop的区别](#delete-truncate-drop的区别)
	- [1. 删除哪些东西](#1-删除哪些东西)
	- [2. 是否能回滚](#2-是否能回滚)
	- [3. 执行速度](#3-执行速度)
	- [4. 自增id变化](#4-自增id变化)
- [参考资料](#参考资料)
# 原理
1. 每个DDL执行前，都会在当前回话执行隐式commit操作. 见[6.6 Additional Notes on DDL and the Normal Transaction](https://dev.mysql.com/doc/internals/en/transactions-notes-on-ddl-and-normal-transaction.html)
> each DDL statement in MySQL begins with an implicit normal transaction commit (a call to end_active_trans()), and thus leaves nothing to modify.

2. DDL执行前，会先执行隐式commit，然后需要获取metadata lock。
- 测试得到的，两个回话`SELECT @@tx_isolation`都得到`REPEATABLE-READ`
```sql
1> begin;
2> begin;
1> insert into user(name) values(1); -- 可重复读隔离级别，此时会添加metadata lock（元数据锁）
2> select * from user; -- 没有发现1中添加的新记录
2> insert into user(name) values(2);
2> create table tmp select * from user; -- metadata lock锁等待
3> select * from user; -- 2中插入的记录可见，1中插入的记录不可见。
1> select * from user; -- 2中插入记录不可见，1中插入记录可见。
```
- 通过上面的例子可以发现，2中获取metadata lock等待，但是2中插入的记录，在3中已经可见。说明执行DDL是先执行隐式commit，然后才获取metadata lock.

# delete truncate drop的区别
## 1. 删除哪些东西 
- truncate 和 delete 只删除数据不删除表的结构(定义);
-  drop 将删除表数据和表定义，也会删除被依赖的约束(constrain)、触发器(trigger)、索引(index)。而依赖于该表的存储过程/函数将保留,但是变为 invalid 状态。
## 2. 是否能回滚
- delete 语句是数据库操作语言(dml)，这个操作会放到 rollback segement 中，事务提交之后才生效；

- drop、truncate都是**DDL语句(数据定义语言),执行前后会自动commit。执行完后不能rollback**。
	- 根据[https://www.quora.com/What-is-the-main-difference-between-Truncate-Delete-and-Drop-in-a-database](https://www.quora.com/What-is-the-main-difference-between-Truncate-Delete-and-Drop-in-a-database) 所说，truncate在执行过程中是可以回滚的，因为在rollback日志中也记录了页。但是truncate一般很快，可能没有进行终止操作的机会。
	- both DELETE and TRUNCATE operations can be COMMITTED AND ROLL-BACKED if provided **inside a Transaction**
> 6. It de-allocates Data Pages instead of Rows and records Data Pages instead of Rows in Transaction logs, thus is faster than DELETE.
> 7. While de-allocating Pages it locks Pages and not Rows, thus it requires less number of locks and few resources.
> 8. This is a DDL command as it resets IDENTITY columns, de-allocates Data Pages and empty them for use of other objects in the database.
> 9. Note: It is a misconception among some people that TRUNCATE cannot be roll-backed. But in reality both DELETE and TRUNCATE operations can be COMMITTED AND ROLL-BACKED if provided inside a Transaction. 
## 3. 执行速度
- 一般来说: drop> truncate > delete
- DELETE   语句每次删除一行，并在事务日志中为所删除的每行记录一项。
- TRUNCATE   TABLE   通过释放存储表数据所用的数据页来删除数据，并且只在事务日志中记录页的释放。
## 4. 自增id变化
- delete后，自增id不会清0，下次insert时自增id持续增加。
- truncate后，自增id清0，下次insert 从0开始。drop也一样。
# 参考资料
- [6.6 Additional Notes on DDL and the Normal Transaction](https://dev.mysql.com/doc/internals/en/transactions-notes-on-ddl-and-normal-transaction.html)
- [SQL truncate 、delete与drop区别](https://www.cnblogs.com/8765h/archive/2011/11/25/2374167.html)
- [https://www.quora.com/What-is-the-main-difference-between-Truncate-Delete-and-Drop-in-a-database](https://www.quora.com/What-is-the-main-difference-between-Truncate-Delete-and-Drop-in-a-database)
