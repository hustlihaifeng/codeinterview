
# 概览
netstat可用来查看socket套接字情况，包括TCP、UDP、unix domain socket。
# 常用筛选命令
- 筛选状态
	- `-a, --all`: 所有状态
	- `-l, --listening`: 正在监听的套接字
- 筛选套接字协议类型：
	- `-t`: TCP
	- `-u`: UDP
	- `-x`: unix domain socket
- 域名解析、用户名解析
	- `-n, --numeric`: 不解析
	- 默认进行解析
- PID/Program name：
	- `-p, --programs`: 打印
	- 默认不打印
- 用户名等其他信息：
	- `-e, --extend`: 打印
	- 默认不打印
# 端口状态详解
![](state%20of%20tcp.png)

- [netstat命令中的TCP SOCKET 状态](https://blog.csdn.net/konga/article/details/8265146)
- [TCP/IP详解--TCP连接中TIME_WAIT状态过多](https://blog.csdn.net/yusiguyuan/article/details/21445883)

# 参考资料
- [netstat 的10个基本用法](https://linux.cn/article-2434-1.html)
- [netstat命令中的TCP SOCKET 状态](https://blog.csdn.net/konga/article/details/8265146)
- [TCP/IP详解--TCP连接中TIME_WAIT状态过多](https://blog.csdn.net/yusiguyuan/article/details/21445883)
