# 1. 背景及目标

X系统需要出一个oracle版本，X系统当前使用`beego`框架开发，sql层使用beego orm编写，数据库使用mysql。

# 2. 难点和问题

1. 按理说，使用orm如果orm框架做的完善的话，切换数据库只需要换一下数据库驱动和连接方式就可以了。但是难就难在beego orm库对oracle支持不完善：beego orm库本身生成的sql语句是mysql式的，到oracle数据库中直接执行不了（mysql转移符号使用\`, 而oracle中使用",orm中生成的每条sql语句中都有转移符号，由此说明beego orm对oracle的支持是非常不完善的，似乎都没有测试过，任何一个简单的sql都跑不通过）
2. beego orm对oracle的其他支持库网上只能找到<https://github.com/satng/beego-orm-by-oracle>，遗憾的是，截至笔者查找时，这还是一个由一个人开发、两颗星、一个fork的测试库，项目中提供的`main.go`连编译都过不了。
3. 那么如果弃用orm，改用裸的sql来兼容oracle呢？改造成本高。

# 3. 解决问题的办法

1. 最后选择采用<https://github.com/satng/beego-orm-by-oracle>这个库，测试其兼容性，不兼容的部分进行改造。

## 3.1 首先是mysql和oracle两者sql语句的兼容性

1. 首先是数据库，oracle中没有mysql中对应的数据库概念，oracle中与mysql中数据库对应的是user，一个user可以建立多个表，多个user有不同表空间，通过user来实现mysql中的数据库隔离。
2. 其次是建表语句，建表语句与数据库特性息息相关，所以<https://github.com/satng/beego-orm-by-oracle>对oracle建表语句基本上算是零兼容，生成的建表语句不能执行。不过这点还好，建表是一次性的，我们可以手动建表。需要把MySQL中的列类型转化为oracle中的列类型，这里我们采用导出mysql表结构，对有差异的地方进行替换的方式：

```shell
mysqldump -utest_user  -h127.0.0.1 -P3306 -p --no-data test_db > test_db.sql
```



### 3.1.1 mysql和oracle建表语句差异

1. oracle不支持`drop table if exists`语句：`%s/DROP TABLE IF EXISTS/DROP TABLE/g`
2. oracle转移符号是":

```shell
%s/`/"/g
```

3. oracle中default语句要在not null语句前面：`%s/\( NOT NULL\)\( DEFAULT '.*'\)/\2\1/g`
4. oracle建表语句中不支持`AUTO_INCREMENT=xxx`的表注释：

```shell
%s/ AUTO_INCREMENT=[0-9]*//g
```

5. oracle总int是integer：`%s/int(11)/INTEGER/g`
6. oracle中没有InnoDB引擎：`%s/ ENGINE=InnoDB//g`
7. oracle建表语句不支持指定字符集：`%s/ DEFAULT CHARSET=utf8//g`
8. oracle12之前没有自增列(用sequence加存储过程可以实现)，oracle12后自增列和mysql语法不同：

```shell
%s/ NOT NULL AUTO_INCREMENT/ GENERATED BY DEFAULT AS IDENTITY/g
```

9. oracle中没有text类型：`%s/ longtext/ CLOB/g`
10. oracle中key不能指定名字：`%s/ KEY "\w\+"//g`
11. oracle中datetime类型是DATE：`%s/ datetime/ DATE/g`
12. oralce中bool类型和mysql不同：`%s/ tinyint(1)/ NUMBER(3,0)/g`
13. oracle中double类型和mysql中不同：`%s/ double/ FLOAT (24)/g`
14. orale中空字符串会被当做成NULL，所以尽量将字段设为NULL型，不然insert会报错：`%s/ NOT NULL/ NULL/g`

注：mysql和oracle列转换关系见 <https://docs.oracle.com/cd/E12151_01/doc.150/e12155/oracle_mysql_compared.htm#BABHHAJC>

### 3.2 orm库对oracle的支持程度

以<https://github.com/satng/beego-orm-by-oracle>为例：

1. 建表语句基本不支持：手动建表
2. 不支持like语句：修改orm库
3. insert获取的id不对：修改orm库，新增一个函数，开启事物，在事物中进行insert和`select order by id desc limit 1`语句。(其实这主要是oracle不支持last insert id特性导致的)
4. 列名中有`group`关键字时导致update语句报错：将列名中的group替换为如`group_col`
5. 时区设置不生效：提供了`orm.SetDataBaseTZ("default", time.Local)`这样一个sql语句来设置时区。

注：本文中对orm库的修改将会提交merge request到<https://github.com/satng/beego-orm-by-oracle>

# 4. 效果

1. 通过手动建表、对orm库进行改造的方式，实现beego orm顺利从mysql切换到oracle，并且没有将orm语句改为裸的sql语句，避免改造sql时的大量工作。其他基于beego orm且构建在mysql上的系统，要往oracle上迁移时，可以按照本文所介绍的方法进行。