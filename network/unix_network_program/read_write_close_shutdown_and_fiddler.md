[TOC]

# read
1. `int read（socketfd，buffer，n）`

2. read返回0，有两种情况：
    1. 对端调用了`close（soketfd）`函数
    2. 对端调用了`close（fd，SHUT_WR)`,关闭对端的写连接，半关闭

3. 从2得知，read读对端写已经关闭的socket，会返回0。那么：
    1. read读取本端读已经关闭的socket会发生什么? TODO:
    2. write写本端写已经关闭的socket会发生什么？TODO:
    3. write写对端读已经关闭的socekt会发生什么?
        - 假设server和client 已经建立了连接，server调用了close, 发送FIN 段给client（其实不一定会发送FIN段，后面再说），此时server不能再通过socket发送和接收数据，此时client调用read，如果接收到FIN 段会返回0，但client此时还是可以write 给server的(因为还要回ack)，write调用只负责把数据交给TCP发送缓冲区就可以成功返回了，所以不会出错，**而已经关闭读的server收到数据后应答一个RST段，表示服务器已经不能接收数据，连接重置，client收到RST段后无法立刻通知应用层，只把这个状态保存在TCP协议层。如果client再次调用write发数据给server，由于TCP协议层已经处于RST状态了，因此不会将数据发出，而是发一个SIGPIPE信号给应用层，SIGPIPE信号的缺省处理动作是终止程序**。
        - 有时候代码中需要连续多次调用write，可能还来不及调用read得知对方已关闭了连接就被SIGPIPE信号终止掉了，这就需要在初始化时调用sigaction处理SIGPIPE信号，对于这个信号的处理我们通常忽略即可，signal(SIGPIPE, SIG_IGN); 如果SIGPIPE信号没有导致进程异常退出（捕捉信号/忽略信号），write返回-1并且errno为EPIPE（Broken pipe）。（非阻塞地write）

4. 什么时候发送RST包:
    1. RST标示复位、用来关闭异常的连接。发送RST包关闭连接时，不必等缓冲区的包都发出去，直接就丢弃缓冲区中的包，发送RST。而接收端收到RST包后，也不必发送ACK包来确认。
    2. 什么时候发送RST包： 
        1.  建立连接的SYN到达某端口，但是该端口上没有正在 监听的服务。
        2.  TCP收到了一个根本不存在的连接上的分节。
        3. 请求超时。 使用setsockopt的SO_RCVTIMEO选项设置recv的超时时间。接收数据超时时，会**发送**RST包。
        4. 正常关闭不会发送RST:
        >          1. 使用shutdown、close关闭套接字，发送的是FIN，不是RST。
        >          2. 套接字关闭前，使用sleep。对运行的程序Ctrl+C，会发送FIN，不是RST。
        >          3. 套接字关闭前，执行return、exit(0)、exit(1)，会发送FIN、不是RST。

5. 常用抓包工具：tcpdump、wireshark

    1. tcpdump常见用法：详见[tcpdump 示例教程](<https://colobu.com/2019/07/16/a-tcpdump-tutorial-with-examples/>)

        - ### 基于IP查找流量:`tcpdump host 1.1.1.1` 查看进出`1.1.1.1`的所有流量

        - ### 根据来源和目标进行筛选:

        ```shell
        tcpdump src 1.1.1.1 
        tcpdump dst 1.0.0.1
        ```

        - ### 显示特定端口的流量

        ```shell
        tcpdump port 3389 
        tcpdump src port 1025
        ```

        

    2. HTTP,HTTPS 用**Fiddler**：详见[Fiddler抓包简易教程](<https://www.jianshu.com/p/9e05a2522758>)

    3. TCP,UDP 用**wireshark**：详见[wireshark使用抓包详细图文教程](<https://blog.csdn.net/ch853199769/article/details/78753963>) 

# 参考资料
1. [socket 中read返回0的情况](https://www.cnblogs.com/kkshaq/p/4456179.html)
2. [linux网络编程之socket（十）：shutdown 与 close 函数 的区别](https://blog.csdn.net/Simba888888/article/details/9068059)
3. [TCP协议RST：RST介绍、什么时候发送RST包](https://blog.csdn.net/guowenyan001/article/details/11766929)