# 1. 列修饰符
## 1.1 列名：`column`
1. 如：`orm:"column(name);"`
## 1.2 类型

#### a. string:varchar(255)

1. 通过`size`设置列长度

   ```go
   Title string `orm:"size(60)"
   ```

#### b. float32, float64:float,double

1. 通过digits / decimals来设置整数部分和小数部分长度

   ```go
   Money float64 `orm:"digits(12);decimals(4)"`
   ```

   总长度 12 小数点后 4 位 eg: `99999999.9999`

#### c. time.Time:datetime       int:timestamp

1. `on create current_timestamp`用`auto_now_add`来表示：第一次保存时才设置时间

2. `on update current_timestamp`用`auto_now`来表示：每次 model 保存时都会对时间自动更新。

   ```go
   Created time.Time `orm:"auto_now_add;type(datetime)"`
   Updated time.Time `orm:"auto_now;type(datetime)"`
   ```

3. **对于批量的 update 此设置是不生效的**

4. type:用type来设置time.Time对应的db类型为`date`或`datetime`

   ```go
   Created time.Time `orm:"auto_now_add;type(date)"`
   Created time.Time `orm:"auto_now_add;type(datetime)"`
   ```

#### d. 模型字段与数据库类型的对应

##### MySQL

| go                                          | mysql                           |
| ------------------------------------------- | ------------------------------- |
| int, int32 - 设置 auto 或者名称为 `Id` 时   | integer AUTO_INCREMENT          |
| int64 - 设置 auto 或者名称为 `Id` 时        | bigint AUTO_INCREMENT           |
| uint, uint32 - 设置 auto 或者名称为 `Id` 时 | integer unsigned AUTO_INCREMENT |
| uint64 - 设置 auto 或者名称为 `Id` 时       | bigint unsigned AUTO_INCREMENT  |
| bool                                        | bool                            |
| string - 默认为 size 255                    | varchar(size)                   |
| string - 设置 type(char) 时                 | char(size)                      |
| string - 设置 type(text) 时                 | longtext                        |
| time.Time - 设置 type 为 date 时            | date                            |
| time.Time                                   | **datetime**                    |
| byte                                        | tinyint unsigned                |
| rune                                        | integer                         |
| int                                         | integer                         |
| int8                                        | tinyint                         |
| int16                                       | smallint                        |
| int32                                       | integer                         |
| int64                                       | bigint                          |
| uint                                        | integer unsigned                |
| uint8                                       | tinyint unsigned                |
| uint16                                      | smallint unsigned               |
| uint32                                      | integer unsigned                |
| uint64                                      | bigint unsigned                 |
| float32                                     | double precision                |
| float64                                     | double precision                |
| float64 - 设置 digits, decimals 时          | numeric(digits, decimals)       |

## 1.3 索引
### 1.3.1 普通索引:`index`
### 1.3.2 主键：`pk`
### 1.3.2 自增列:`auto`

1. `auto`:当 Field 类型为 int, int32, int64, uint, uint32, uint64 时，可以用`auto`关键字设置字段为自增健。
2. 当模型定义里**没有主键**时，符合**上述类型**且名称为 **`Id`** 的 Field 将被视为**自增健**

### 1.3.4 唯一键:`unique`
### 1.3.5 多列索引:`TableIndex`、`TableUnique`
```go
type User struct {
    Id    int
    Name  string
    Email string
}

// 多字段索引
func (u *User) TableIndex() [][]string {
    return [][]string{
        []string{"Id", "Name"},
    }
}

// 多字段唯一键
func (u *User) TableUnique() [][]string {
    return [][]string{
        []string{"Name", "Email"},
    }
}
```
## 1.3.6 默认值:`default`
## 1.3.7 是否为空：`null`

1. 所有字段默认**NOT NULL**

## 1.3.8 注释:`description`

1. 如

   ```go
   type User struct {
       ...
       Status int `orm:"default(1)" description:(这是状态字段)`
       ...
   }
   ```

2. **注释中禁止包含引号,上面的注释不在双引号中**

### 1.3.9 其他

#### 分隔符

多个设置间使用 `;` 分隔，设置的值如果是多个，使用 `,` 分隔。

#### 忽略字段

设置 `-` 即可忽略 struct 中的字段

```go
type User struct {
...
    AnyField string `orm:"-"`
...
}
```

#### 默认名称映射

1. 除了开头的大写字母以外，遇到大写会增加 `_`，原名称中的下划线保留。

   ```
   DB_AuthUser -> d_b__auth_user
   ```

# 2. 表修饰符

## 2.1 表名:`TableName`

```go
type User struct {
    Id int
    Name string
}

func (u *User) TableName() string {
    return "auth_user"
}
```

### 表名前缀:`RegisterModelWithPrefix`

```go
orm.RegisterModelWithPrefix("prefix_", new(User))
```

创建后的表名为 `prefix_auth_user`

## 2.2 引擎：`TableEngine`

```go
// 设置引擎为 INNODB
func (u *User) TableEngine() string {
    return "INNODB"
}
```

## 2.3 字符集 :TODO

# 3. 其他

## 3.1 自动建表

1. `force=1`：drop tables before create

```go
// 数据库别名
name := "default"

// drop table 后再建表
force := true

// 打印执行过程
verbose := true

// 遇到错误立即返回
err := orm.RunSyncdb(name, force, verbose)
if err != nil {
    fmt.Println(err)
}
```



1. **非force模式**：

- 自动建表功能在非 force 模式下，会自动创建新增加的字段、新增加的索引。
- 对于**改动过的旧字段，旧索引，需要用户自行进行处理**

2. RunSyncdb `func(name string, force bool, verbose bool) error`
   RunSyncdb run syncdb command line. name means table's alias name. default is "default". **force means run next sql if the current is error**(与上面注释的不一样啊). verbose means show all info when running command or not.

## 3.2 安装orm

```go
go get github.com/astaxie/beego/orm
```

## 3.3 初始化操作

```go
import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
    Id   int
    Name string `orm:"size(100)"`
}

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
    // set default database
    // 参数1        数据库的别名，用来在 ORM 中切换数据库使用
    // 参数2        driverName
    // 参数3        对应的链接字符串
    orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")

    // 参数4(可选)  设置最大空闲连接
    // 参数5(可选)  设置最大数据库连接 (go >= 1.2)
    maxIdle := 30
    maxConn := 30
    orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8", maxIdle, maxConn)

    // register model
    orm.RegisterModel(new(User))

    // create table,也可以没有这个，那么我们自己建好表就行
    orm.RunSyncdb("default", false, true)
}
```

## 3.4 时区

1. ORM 默认使用 time.Local 本地时区

- 作用于 ORM 自动创建的时间
- 从数据库中取回的时间转换成 ORM 本地时间

2. 如果需要的话，你也可以进行更改

```
// 设置为 UTC 时间
orm.DefaultTimeLoc = time.UTC
```

3. ORM 在进行 **RegisterDataBase** 的同时，会获取数据库使用的时区，然后在 time.Time 类型存取时做相应转换，以匹配时间系统，从而保证时间不会出错。
4. 使用 `go-sql-driver` 驱动时(使用MySQL一般会加上这个)，请注意参数设置: 从某一版本开始，驱动默认使用 UTC 时间，而非本地时间，所以请指定时区参数或者全部以 UTC 时间存取
   例如：`root:root@/orm_test?charset=utf8&loc=Asia%2FShanghai`
   参见 [loc](https://github.com/go-sql-driver/mysql#loc) / [parseTime](https://github.com/go-sql-driver/mysql#parsetime)

## 3.5 注册模型

如果使用 orm.QuerySeter 进行高级查询的话，这个是必须的。

反之，**如果只使用 Raw 查询和 map struct，是无需这一步的**。您可以去查看 [Raw SQL 查询](https://beego.me/docs/mvc/model/rawsql.md)

#### RegisterModel

将你定义的 Model 进行注册，最佳设计是有单独的 models.go 文件，在他的 init 函数中进行注册。

迷你版 models.go

```
package main

import "github.com/astaxie/beego/orm"

type User struct {
    Id   int
    Name string
}

func init(){
    orm.RegisterModel(new(User))
}
```

RegisterModel 也可以同时注册多个 model

```
orm.RegisterModel(new(User), new(Profile), new(Post))
```

详细的 struct 定义请查看文档 [模型定义](https://beego.me/docs/mvc/model/models.md)

#### RegisterModelWithPrefix

使用表名前缀

```
orm.RegisterModelWithPrefix("prefix_", new(User))
```

创建后的表名为 prefix_user

#### QueryTable

传入表名，或者 Model 对象，返回一个 [QuerySeter](https://beego.me/docs/mvc/model/query.md)

```
o := orm.NewOrm()
var qs orm.QuerySeter
qs = o.QueryTable("user")
// 如果表没有定义过，会立刻 panic
```

## 3.6 GetDB

从已注册的数据库返回 *sql.DB 对象，默认返回别名为 default 的数据库。

```
db, err := orm.GetDB()
if err != nil {
    fmt.Println("get default DataBase")
}

db, err := orm.GetDB("alias")
if err != nil {
    fmt.Println("get alias DataBase")
}
```

## 3.7 调试模式打印查询语句

简单的设置 Debug 为 true 打印查询的语句

可能存在性能问题，不建议使用在生产模式

```
func main() {
    orm.Debug = true
...
```

## 3.8 ORM 接口使用

使用 ORM 必然接触的 Ormer 接口，我们来熟悉一下

```
var o orm.Ormer
o = orm.NewOrm() // 创建一个 Ormer
// NewOrm 的同时会执行 orm.BootStrap (整个 app 只执行一次)，用以验证模型之间的定义并缓存。
```

切换数据库，或者，进行事务处理，都会作用于这个 Ormer 对象，以及其进行的任何查询。

所以：需要 **切换数据库**（Using） 和 **事务处理**（Begin、Rollback、Commit） 的话，不要使用全局保存的 Ormer 对象。

- type Ormer interface {
  - [Read(interface{}, …string) error](https://beego.me/docs/mvc/model/object.md#read)
  - [ReadOrCreate(interface{}, string, …string) (bool, int64, error)](https://beego.me/docs/mvc/model/object.md#readorcreate)
  - [Insert(interface{}) (int64, error)](https://beego.me/docs/mvc/model/object.md#insert)
  - [InsertMulti(int, interface{}) (int64, error)](https://beego.me/docs/mvc/model/object.md#insertmulti)
  - [Update(interface{}, …string) (int64, error)](https://beego.me/docs/mvc/model/object.md#update)
  - [Delete(interface{}) (int64, error)](https://beego.me/docs/mvc/model/object.md#delete)
  - [LoadRelated(interface{}, string, …interface{}) (int64, error)](https://beego.me/docs/mvc/model/query.md#%E8%BD%BD%E5%85%A5%E5%85%B3%E7%B3%BB%E5%AD%97%E6%AE%B5)
  - [QueryM2M(interface{}, string) QueryM2Mer](https://beego.me/docs/mvc/model/query.md#%E5%A4%9A%E5%AF%B9%E5%A4%9A%E5%85%B3%E7%B3%BB%E6%93%8D%E4%BD%9C)
  - [QueryTable(interface{}) QuerySeter](https://beego.me/docs/mvc/model/orm.md#querytable)
  - [Using(string) error](https://beego.me/docs/mvc/model/orm.md#using)
  - [Begin() error](https://beego.me/docs/mvc/model/transaction.md)
  - [Commit() error](https://beego.me/docs/mvc/model/transaction.md)
  - [Rollback() error](https://beego.me/docs/mvc/model/transaction.md)
  - [Raw(string, …interface{}) RawSeter](https://beego.me/docs/mvc/model/orm.md#raw)
  - [Driver() Driver](https://beego.me/docs/mvc/model/orm.md#driver)
- }

#### Raw

使用 sql 语句直接进行操作

Raw 函数，返回一个 [RawSeter](https://beego.me/docs/mvc/model/rawsql.md) 用以对设置的 sql 语句和参数进行操作

```go
o := orm.NewOrm()
var r orm.RawSeter
r = o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "slene")
```