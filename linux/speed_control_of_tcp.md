
# 速度控制
1. 一个 TCP 发送方能向网路中发送流量的速率：
```
LastByteSent-LastByteAcked≤min{CongWin, RcvWindow}
```
- `CongWin`（拥塞窗口，congestion window），TCP连接的每一端都记录一个额外的变量 `CongWin`
-  `RcvWindow`: 指接受窗口，也即接收方还有多少缓存空间。
	- TCP报文段首部中，“窗口大小”字段通常用于告知对方自己的能够接受的数据量大小。也叫接收窗口，上面的`RcvWindow`可以看做对方的TCP首部中窗口大小。详见[TCP报文段首部中“窗口”字段](https://blog.csdn.net/bian_qing_quan11/article/details/72636675)
- `LastByteSent`: 对应于`tcp`首部的序号字段。TCP 连接中传送的数据流中的**每一个字节都编上一个序号**。**序号字段**的值则指的是本报文段所发送的数据的**第一个字节**在整个报文字节流中的**序号**。
- `LastByteAcked`: 对应于`tcp`首部的确认号字段。含义类似于序号。确认号会依据上次收到的对方的包里面的序号字段和数据流长度来计算：上次收包序号+上次收包数据字节长度（不包括首部）

# 滑动窗口代表的流量控制
TODO：没收到一个`ACK`窗口后移。
# 拥塞控制
TODO:慢启动、加性增、乘性减。

# 参考资料
- [TCP报文段首部中“窗口”字段](https://blog.csdn.net/bian_qing_quan11/article/details/72636675)
- [TCP的seq和ack号计算方法](https://blog.csdn.net/HappyRocking/article/details/78198776)