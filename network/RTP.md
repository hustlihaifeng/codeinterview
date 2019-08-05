1. # RTP概述
   1. 协议层级：应用程序通常在 `UDP` 上运行 `RTP` 以便使用其多路结点和校验服务；这两种协议都提供了传输层协议的功能。
   2. 可靠性： `RTP` 并不保证传送或防止无序传送，也不确定底层网络的可靠性。 `RTP` 实行有序传送， `RTP`中的序列号允许接收方重组发送方的包序列，同时序列号也能用于决定适当的包位置，例如：在视频解码中，就不需要顺序解码。RTP提供抖动补偿和数据无序到达检测的机制。由于IP网络的传输特性，数据的无序到达是很常见的。 RTP允许数据通过IP组播的方式传送到多个目的地。
   3. 组成：RTP标准定义了两个子协议，RTP和RTCP
      1. 数据传输协议RTP：用于实时传输数据。该协议提供的信息包括：时间戳（用于同步）、序列号（用于丢包和重排序检测）、以及负载格式（用于说明数据的编码格式）
      2. 控制协议`RTCP`：用于`QoS`反馈和同步媒体流。相对于`RTP`来说，`RTCP`所占的带宽非常小，通常只有5%
   4. RTP回话：每一个多媒体流会建立一个`RTP`会话，一个会话包含带有`RTP`和`RTCP`端口号的`IP`地址。形成会话的端口由其他协议（例如`RTSP`和SIP）来协商。RTP和RTCP使用UDP端口1024 - 65535。
   # 和RTCP RTSP RTMP等的关系

   1. 主流RTP实现构建在 [User Datagram Protocol](https://en.wikipedia.org/wiki/User_Datagram_Protocol) (UDP)协议上.[[3\]](https://en.wikipedia.org/wiki/Real-time_Transport_Protocol#cite_note-Perkins_46-3) 其他专门为多媒体回话构建的传输层协议包括 [SCTP](https://en.wikipedia.org/wiki/SCTP)[[5\]](https://en.wikipedia.org/wiki/Real-time_Transport_Protocol#cite_note-5) and [DCCP](https://en.wikipedia.org/wiki/DCCP),[[6\]](https://en.wikipedia.org/wiki/Real-time_Transport_Protocol#cite_note-6) 但是到2102年他们还没有被广泛使用.[[7\]](https://en.wikipedia.org/wiki/Real-time_Transport_Protocol#cite_note-7)
   2. RTSP：Real Time Streaming Protocol

   # 协议头

   1. `rtp`包在应用层创建，被传到传输层去传输。`RTP`包头格式如下：
   2. rtp协议头格式如下：

   ![](D:\code\codeinterview\src\network\RTP_header.png)

   - 最小12 bytes，After the header, optional header extensions may be present.
   - **Version**: (2 bits) Indicates the version of the protocol. Current version is 2
   - **P (Padding)**: (1 bit)表名RTP包结尾是否有填充字节。如果有的话，最后一个字节表示填充字节长度（包括自身)

   # 协议工作方式

   1. RTP 使用偶数端口号接收发送数据，相应的RTCP则使用相邻的下一位奇数端口号
