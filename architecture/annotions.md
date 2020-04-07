[TOC]


# 1. 性能指标

## 1. QPS

1. **Query Per Second**,每秒请求数，就是说服务器在一秒的时间内处理了多少个请求。
2. QPS = 总请求数 / ( 进程总数 * 请求时间 )
3. 峰值QPS估算，每天80%的访问集中在20%的时间里，这20%时间叫做峰值时间公式：( 总PV数 * 80% ) / ( 每天秒数 * 20% ) = 峰值时间每秒请求数(QPS)

## 2. TPS

1. **Transactions Per Second**,每秒事务数

## 3. RT

1. **响应时间**, 指系统对请求作出响应的时间，一般取平均响应时间。可以通过Nginx、Apache之类的Web Server得到



# 2. 用户指标
## 1. PV

1. **Page View**,页面访问量，即页面浏览量或点击量，用户每次刷新即被计算一次。可以统计服务一天的访问日志得到。

## 2. UV

1. **Unique Visitor**,独立访客，统计1天内访问某站点的用户数。可以统计服务一天的访问日志并根据用户的唯一标识去重得到

## 3. DAU/MAU

1. **Daily Active User**/**Month Active User**, DAU通常统计一日（统计日）之内，登录或使用了某个产品的用户数（去除重复登录的用户）.



# 3. 常用软件性能指标

## 3.1 MySQL

1. [MySQL数据库三个关键性能指标--TPS\QPS\IOPS](https://database.51cto.com/art/201910/604214.htm) 查询方式:

   - TPS

   ```sql
   use information_schema;
   select VARIABLE_VALUE into @num_com from GLOBAL_STATUS where VARIABLE_NAME ='COM_COMMIT';
   select VARIABLE_VALUE into @num_roll from GLOBAL_STATUS where VARIABLE_NAME ='COM_ROLLBACK';
   select VARIABLE_VALUE into @uptime from GLOBAL_STATUS where VARIABLE_NAME ='UPTIME';
   select (@num_com+@num_roll)/@uptime;
   ```

   - QPS

   ```sql
   use information_schema;
   select VARIABLE_VALUE into @num_queries from GLOBAL_STATUS where VARIABLE_NAME ='QUESTIONS';
   select VARIABLE_VALUE into @uptime from GLOBAL_STATUS where VARIABLE_NAME ='UPTIME';
   select @num_queries/@uptime;
   ```

# 参考资料

1. [QPS、RT、PV、UV、SLA、DAU 介绍](<https://www.jianshu.com/p/ef44f5c11115>)
2. [PV、UV、QPS、并发数、TPS概念以及计算方式](<https://blog.csdn.net/boonya/article/details/104743027>)