# 目录
- [什么是TCP粘包问题](#什么是TCP粘包问题)
- [为什么TCP会出现粘包问题](#为什么TCP会出现粘包问题)
- [为什么UDP没有粘包问题](#为什么UDP没有粘包问题)
- [TCP的粘包问题解决方案](#TCP的粘包问题解决方案)
	- [4. 封包解包](#4-封包解包)
		- [4.1 动态缓冲区暂存方式](#4-1-动态缓冲区暂存方式)
		- [4.2 利用底层的缓冲区来进行拆包](#4-2-利用底层的缓冲区来进行拆包)
- [参考文献](#参考文献)

# 什么是TCP粘包问题
1. TCP的粘包问题指接收端收到的包中包含两个或多个包的数据，形象的说就是两个包黏在一起，而且还有可能是两个部分包。
# 为什么TCP会出现粘包问题
1. **发送端**开启了Nagle算法:Nagle算法为了优化网络效率，会将多个小包合并发送。具体是当发送缓冲区没满时，Nagle算法会等待一小段时间，看看有没有其他包需要发送，如果有则将缓冲区填满再发送，除非有紧急数据或者超过等待时间。详见[TCP中的Nagle算法](https://blog.csdn.net/ce123_zhouwei/article/details/9050797)
2. **接收端**接收不及时造成的接收端粘包:TCP会把接收到的数据存在自己的缓冲区中,然后通知应用层取数据.当应用层由于某些原因不能及时的把TCP的数据取出来,就会造成TCP缓冲区中存放了几段数据。
# 为什么UDP没有粘包问题
1. 因为UDP是面向消息的协议，发送端不会把多个包合并发送，接收端也只能接受独立的消息。接收端会把所有接收到的消息都挂接到缓冲区的链式接受队列中，这样接收端应用程序一次recv只能从socket接收缓冲区中读出一个数据包（无论recv时指定了多大的缓冲区）。也就是说，UDP会保护消息边界。（UDP最大载荷为1472字节）
2. TCP是面向字节流的协议，把数据当做一串数据流。所以为了优化效率，会采用发送端合并的算法。因为面向流，所以接收端缓冲区是一个整体而非udp一样链式分割开。
3. TCP首部第四字节中的长度指首部长度，4位，4字节为单位，最多`(2^4-1)*4=60`字节（也叫数据偏移，指数据部分相对于TCP整个包中的字节偏移）。UDP首部中的长度是UDP整个包的字节长度，包括首部和数据两部分，理论范围为`8~2^16`，8是UDP首部的固定长度，即发送数据长度为0的消息也是可以的。**因为TCP首部中只存了首部长度，没有存总长度或者数据长度，所有粘包问题不能在现有基础上直接在接收端处理掉**。
# TCP的粘包问题解决方案
1. 把发送端的合并发送的算法关掉，任然有可能接收方粘包，且降低了网络传输效率。方法：
  - 使用`TCP_NODELAY`选项可以禁止`Negale`算法。默认情况下,发送数据采用Negale算法。
  - `TCP_NODELAY`和`TCP_CORK`都是禁用Nagle算法，只不过`TCP_NODELAY`完全关闭而`TCP_CORK`完全由自己决定发送时机。Linux文档上说两者不要同时设置。详见:[TCP选项：TCP_NODELAY和TCP_CORK（negle算法）](https://blog.csdn.net/caodongfang126/article/details/78029999)
2. 发送完后等待一小段时间，任然有可能接收方粘包。
3. 使用短连接，发送完数据就断开连接。详见：[https://blog.csdn.net/zhangxinrun/article/details/6721495](https://blog.csdn.net/zhangxinrun/article/details/6721495)
## 4. 封包解包
**在TCP包的数据部分，加上包长度结构体**。接收方读数据时，先读取包长度结构体长的数据，获取包长度，然后在读取指定长度的字节数据，来读取一个完整包。也即封包和拆包。**这部分拷贝自：[（经典）tcp粘包分析](https://blog.csdn.net/zhangxinrun/article/details/6721495)**,调整了下格式，方便查看。

### 4.1 动态缓冲区暂存方式
- 之所以说缓冲区是动态的是因为当需要缓冲的数据长度超出缓冲区的长度时会增大缓冲区长度.
- 大概过程描述如下:
```
    A,为每一个连接动态分配一个缓冲区,同时把此缓冲区和SOCKET关联,常用的是通过结构体关联.
    B,当接收到数据时首先把此段数据存放在缓冲区中.
    C,判断缓存区中的数据长度是否够一个包头的长度,如不够,则不进行拆包操作.
    D,根据包头数据解析出里面代表包体长度的变量.
    E,判断缓存区中除包头外的数据长度是否够一个包体的长度,如不够,则不进行拆包操作.
    F,取出整个数据包.这里的"取"的意思是不光从缓冲区中拷贝出数据包,而且要把此数据包从缓存区中删除掉.删除的办法就是把此包后面的数据移动到缓冲区的起始地址.
```
- 这种方法有两个缺点:
	- 1.为每个连接动态分配一个缓冲区增大了内存的使用.
	- 2.有三个地方需要拷贝数据:一个地方是把数据存放在缓冲区,一个地方是把完整的数据包从缓冲区取出来,一个地方是把数据包从缓冲区中删除.
- 改进办法: 即采用环形缓冲
	- 但是这种改进方法还是不能解决第一个缺点以及第一个数据拷贝,只能解决第三个地方的数据拷贝(这个地方是拷贝数据最多的地方).
	- 环形缓冲实现方案是定义两个指针,分别指向有效数据的头和尾.在存放数据和删除数据时只是进行头尾指针的移动.

### 4.2 利用底层的缓冲区来进行拆包

1. 由于TCP也维护了一个缓冲区,所以我们完全可以利用TCP的缓冲区来缓存我们的数据,这样一来就不需要为每一个连接分配一个缓冲区了.另一方面我们知道recv或者wsarecv都有一个参数,用来表示我们要接收多长长度的数据.利用这两个条件我们就可以对第一种方法进行优化.
2. 对于阻塞SOCKET来说,我们可以利用一个循环来接收包头长度的数据,然后解析出代表包体长度的那个变量,再用一个循环来接收包体长度的数据.相关代码如下:
```cpp
char PackageHead[1024];
char PackageContext[1024*20];
int len;
PACKAGE_HEAD *pPackageHead;
while( m_bClose == false )
{
	memset(PackageHead,0,sizeof(PACKAGE_HEAD));
	len =m_TcpSock.ReceiveSize((char*)PackageHead,sizeof(PACKAGE_HEAD));
	if( len == SOCKET_ERROR )
	{
		break;
	}
	if(len == 0)
	{
		break;
	}
	pPackageHead = (PACKAGE_HEAD *)PackageHead;
	memset(PackageContext,0,sizeof(PackageContext));
	if(pPackageHead->nDataLen>0)
	{
		len = m_TcpSock.ReceiveSize((char*)PackageContext,pPackageHead->nDataLen);
	}
}
```
`m_TcpSock`是一个封装了SOCKET的类的变量,其中的`ReceiveSize`用于接收一定长度的数据,直到接收了一定长度的数据或者网络出错才返回.
```cpp
int winSocket::ReceiveSize( char* strData, int iLen )
{
	if( strData == NULL )
	return ERR_BADPARAM;
	char *p = strData;
	int len = iLen;
	int ret = 0;
	int returnlen = 0;
	while( len > 0)
	{
		ret = recv( m_hSocket, p+(iLen-len), iLen-returnlen, 0 );
		if ( ret == SOCKET_ERROR || ret == 0 )
		{
			return ret;
		}
		len -= ret;
		returnlen += ret;
	}
	return returnlen;
}
```
# 参考文献
- [（经典）tcp粘包分析](https://blog.csdn.net/zhangxinrun/article/details/6721495)
- [TCP中的Nagle算法](https://blog.csdn.net/ce123_zhouwei/article/details/9050797)
- [TCP粘包，UDP不存在粘包问题](https://blog.csdn.net/hik_zxw/article/details/48398935)
- [TCP选项：TCP_NODELAY和TCP_CORK（negle算法）](https://blog.csdn.net/caodongfang126/article/details/78029999)
