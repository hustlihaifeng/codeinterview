[TOC]

# 1. RTT RTTs RTTd RTO

## 1.1 是什么

1. RTT——Round Trip Time: 采样值，非重传的包，在发送到收到ACK的时间

- RTTs（ Smoothed RTT）, : **`RTTs=  (1- α) * RTTs+ α * RTT`**, RFC 2988推荐的α值为1/8

- RTTd：RTTd是RTT的偏差的加权平均值，与RTTs和新的RTT样本之差有关。`RTTd=(1-B)*RTTd+B*|RTTs-RTT|`,当第一次测量时，RTTd值取为RTT样本值的一半. B的推荐值是1/4

- RTO：Retransmission TimeOut,超时重传时间. `RTO=RTTs+4*RTTd`

2. 为了解决重传后不知道ACK是初始包还是重传包的问题，Karn提出了一个算法：在计算加权平均RTTs时，只要报文段重传了，就不采用其往返时间样本。这样得出的加权平均RTTs和RTO就相对比较准确了。
- 但是，但是，要是出现这样的情况呢？？：报文段的时延突然增大了很多。因此在原来得出的重传时间内，不会收到确认报文段。于是就重传报文段。但根据Karn算法，不考虑重传的报文段的往返时间样本。这样：超时重传时间就无法更新。
- 因此：要对Karn算法进行修正：方法是：报文段每重传一次，就把超时冲传时间RTO增大一些。典型的做法是：取新的重传时间为2倍的旧的重传时间。当不再发生报文段的重传时，才根据上面给出公式计算超时重传时间



# 参考资料

1. [超时重传的时间计算](<https://blog.csdn.net/msdnwolaile/article/details/51227491>)
2. [TCP 的那些事儿（下）](<https://coolshell.cn/articles/11609.html>)