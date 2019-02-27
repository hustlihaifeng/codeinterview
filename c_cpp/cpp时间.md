1. 与时间有关的函数头文件是[`<time.h>`](http://www.runoob.com/cprogramming/c-standard-library-time-h.html) `<ctime>`
2. 有四个与时间相关的类型
     `clock_t`、`time_t`、`size_t` 和 `tm`。类型 `clock_t`、`size_t` 和 `time_t` 能够把系统时间和日期表示为某种整数。结构类型 `tm` 把日期和时间以 `C` 结构的形式保存，`tm` 结构的定义如下：
```c
struct tm {
   int tm_sec;         /* 秒，范围从 0 到 59        */
   int tm_min;         /* 分，范围从 0 到 59        */
   int tm_hour;        /* 小时，范围从 0 到 23        */
   int tm_mday;        /* 一月中的第几天，范围从 1 到 31    */
   int tm_mon;         /* 月，范围从 0 到 11        */
   int tm_year;        /* 自 1900 年起的年数        */
   int tm_wday;        /* 一周中的第几天，范围从 0 到 6    */
   int tm_yday;        /* 一年中的第几天，范围从 0 到 365    */
   int tm_isdst;       /* 夏令时                */
};
```

3. 常用函数
- [`clock_t clock(void);`](http://www.runoob.com/cprogramming/c-function-clock.html) //用来计时，做性能对比

  - 该函数返回程序执行起（一般为程序的开头），处理器时钟数。
  - `typedef long clock_t;`
  - `#define CLOCKS_PER_SEC ((clock_t)1000)` 表示一秒钟会有多少个时钟计时单元，不能通过改变`CLOCKS_PER_SEC`来提高计时精度。
- [`time_t time(time_t *time);` ](http://www.runoob.com/cprogramming/c-function-time.html)//时间显示的基础
	- 该函数返回系统的当前日历时间，自 1970 年 1 月 1 日以来经过的秒数。
	- `typedef long time_t;`
	- 如果time不为空，则返回值也存储在time里面。
- [`char *ctime(const time_t *time);`](http://www.runoob.com/cprogramming/c-function-ctime.html) //配合time()用来打印当前时间
	- 该返回一个表示当地时间的字符串指针，字符串形式 `day month year hours:minutes:seconds year\n\0`。
	- 类似：[`char * asctime ( const struct tm * time );`](http://www.runoob.com/cprogramming/c-function-asctime.html)

- [`struct tm *localtime(const time_t *time);`](http://www.runoob.com/cprogramming/c-function-localtime.html) //配合time()用来获取具体的时分秒等值
	- 该函数返回一个指向表示本地时间的 tm 结构的指针。
	- 逆函数：	[`time_t mktime(struct tm *time);`](http://www.runoob.com/cprogramming/c-function-mktime.html)
	- 同类：	[`struct tm *gmtime(const time_t *time);`](http://www.runoob.com/cprogramming/c-function-gmtime.html)GMT格林尼治时间
	```cpp
	struct tm {
	  int tm_sec;   // 秒，正常范围从 0 到 59，但允许至 61
	  int tm_min;   // 分，范围从 0 到 59
	  int tm_hour;  // 小时，范围从 0 到 23
	  int tm_mday;  // 一月中的第几天，范围从 1 到 31
	  int tm_mon;   // 月，范围从 0 到 11
	  int tm_year;  // 自 1900 年起的年数
	  int tm_wday;  // 一周中的第几天，范围从 0 到 6，从星期日算起
	  int tm_yday;  // 一年中的第几天，范围从 0 到 365，从 1 月 1 日算起
	  int tm_isdst; // 夏令时
	}
	```
- [`double difftime ( time_t endtime, time_t starttime );`](http://www.runoob.com/cprogramming/c-function-difftime.html)
该函数返回endtime 和 starttime 之间相差的秒数。
- [`size_t strftime(char *str, size_t maxsize, const char *format, const struct tm *timeptr)`](http://www.runoob.com/cprogramming/c-function-strftime.html) 根据 format 中定义的格式化规则，格式化结构 timeptr 表示的时间，并把它存储在 str 中。格式符：
|说明符| 替换为 |实例|
|-|-|-|
|%a |缩写的星期几名称| Sun|
|%A |完整的星期几名称| Sunday|
|%b |缩写的月份名称| Mar|
|%B |完整的月份名称| March|
|%c |日期和时间表示法| Sun Aug 19 02:56:02 2012|
|%d |一月中的第几天（01-31）| 19|
|%H |24 小时格式的小时（00-23）| 14|
|%I |12 小时格式的小时（01-12）| 05|
|%j |一年中的第几天（001-366）| 231|
|%m |十进制数表示的月份（01-12）| 08|
|%M |分（00-59）| 55|
|%p |AM 或 PM 名称 |PM|
|%S |秒（00-61） |02|
|%U |一年中的第几周，以第一个星期日作为第一周的第一天（00-53） |33|
|%w |十进制数表示的星期几，星期日表示为 0（0-6） |4|
|%W |一年中的第几周，以第一个星期一作为第一周的第一天（00-53） |34|
|%x |日期表示法| 08/19/12|
|%X |时间表示法| 02:50:06|
|%y |年份，最后两个数字（00-99）| 01|
|%Y |年份 |2012|
|%Z |时区的名称或缩写 |CDT|
|%% |一个 % 符号 |%|
