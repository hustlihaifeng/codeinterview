# 目录

# cpu
## 使用情况查看
1. top： 主要查看cpu使用情况
top中各列信息：
|列名|含义|
|-|-|
|PID|进程的ID|
|USER|进程所有者|
|PR|进程的优先级别，越小越优先被执行|
|NI|nice值|
|VIRT|进程占用的虚拟内存|
|RES|进程占用的物理内存|
|SHR|进程使用的共享内存|
|S|进程的状态。S表示休眠，R表示正在运行，Z表示僵死状态，N表示该进程优先值为负数|
|%CPU|进程占用CPU的使用率|
|%MEM|进程使用的物理内存和总内存的百分比|
|TIME+|该进程启动后占用的总的CPU时间，即占用CPU使用时间的累加值。|
|COMMAND|进程启动命令名称|

# 内存
1. free
# 磁盘
1. iostat
2. iotop
# 网络
1. tcpdump: 抓包工具
2. iftop:可以查看当前与各个ip通信的速度和排名，类似于iotop，top
3. ifstat：
# 参考资料
- [Linux下查看版本、CPU、内存、磁盘、Swap、网络等资源的使用情况](https://blog.csdn.net/CSDN_duomaomao/article/details/77877108)