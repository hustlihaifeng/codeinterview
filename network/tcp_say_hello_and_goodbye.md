[TOC]





# 参考资料

1. [20-1-tcp连接——初始化序列号(ISN)](<https://blog.csdn.net/qq_35733751/article/details/80552037>)

   > 在前面学习tcp连接三次握手的时候，客户端和服务端在建立tcp连接时，双方都会发送SYN报文并初始化序号（英文为：Initial Sequence Number，简称ISN）
   >
   >   换句话说，只要A发送了一个tcp报文段，且这个tcp报文段的四元组和序号，和之前的tcp连接（四元组和序号）相同的话，就会被B确认。这其实反映了tcp的一些缺点，如果被一些恶意攻击者加以利用tcp的这种缺点：选择合适的序号，ip地址和端口的话，就能伪造出一个tcp报文段，从而打断正常的tcp连接。但是初始化序号的方式（通过算法来随机生成序号）就会使序号难以猜出，也就不容易利用这种缺点来进行一些恶意攻击行为。
   >
   >   通过上面所述我们知道，如果A和B之间发送数据每次都使用相同序号的话可能会引发一系列的问题，但是使用不同序号的话，那么B在接收到这个序号为1的tcp报文时，发现这个tcp报文的序号不在新tcp连接的接收范围内时会把这个tcp报文丢弃掉，也就避免了数据乱序的问题。
   >
   >   因此我们可以明白，客户端和服务端双方在建立tcp连接并初始化序列号，那么上面所说的这些情况从一开始就可以避免。另外，tcp在初始化序列号的过程也是比较复杂的，一般来说，这个序号的范围是`0~ 2^31−1`之间，而且序号的生成也是随机的，通常是一个很大的数值，也就是说每个tcp连接使用的序号也是不一样的。

   

2. [两张动图-彻底明白TCP的三次握手与四次挥手](<https://blog.csdn.net/qzcsu/article/details/72861891>)  动图将三次握手和四次挥手

3. [TCP 为什么是三次握手，而不是两次或四次？](<https://www.zhihu.com/question/24853633>)

> TCP连接的一方A，由操作系统动态随机选取一个**32位长的序列号（Initial Sequence Number），**假设A的初始序列号为1000，以该序列号为原点，对自己将要发送的每个字节的数据进行编号，1001，1002，1003…，并把自己的初始序列号ISN告诉B，让**B有一个思想准备**，什么样编号的数据是合法的，什么编号是非法的，比如编号900就是非法的，同时B还可以对A每一个编号的字节数据进行确认。如果A收到B确认编号为2001，则意味着字节编号为[1001-2000](tel:1001-2000)，共1000个字节已经安全到达。
>
> **一句话概括，TCP连接握手，握的是啥？**
>
> **通信双方数据原点的序列号！**

并分析了三次握手中各个包丢失的处理方法。

