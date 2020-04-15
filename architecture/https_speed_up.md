[TOC]



# 1. https性能优化

## 1.1 一直到今天我们还有66%的网站不支持HTTPS呢

1. 慢

> HTTPS未经任何优化的情况下要比HTTP慢几百毫秒以上，特别在移动端可能要慢500毫秒以上

2. 贵

> 特别在计算性能和服务器成本方面。HTTPS为什么会增加服务器的成本？相信大家也都清楚HTTPS要额外计算，要频繁地做加密和解密操作，几乎每一个字节都需要做加解密，这就产生了服务器成本

## 1.2 基于nginx的https性能优化

1. HTTP/2：

   1. 每个服务器只用一个连接，节省多次建立连接的时间，在TLS上效果尤为明显；
   2. 加速 TLS 交付，HTTP/2 只耗时一次 TLS 握手，通过一个连接上的多路利用实现最佳性能

2. TLS 1.3：

   1. **握手时间：**同等情况下，TLSv1.3 比 TLSv1.2 少一个 RTT
   2. **应用数据：**在会话复用场景下，支持 0-RTT 发送应用数据
   3. **会话复用机制：**弃用了 Session ID 方式的会话复用，采用 PSK 机制的会话复用

3. ### Brotli：

   1. Brotli 是由 Google 于 2015 年 9 月推出的无损压缩算法，它通过用变种的 LZ77 算法，Huffman 编码和二阶文本建模进行数据压缩，是一种压缩比很高的压缩方法。
   2. 针对常见的 Web 资源内容，Brotli 的性能要比 Gzip 好 17-25%；
   3. Brotli 的支持必须依赖 HTTPS，不过换句话说就是只有在 HTTPS 下才能实现 Brotli。

4. ### ECC 证书：

   1. 由于 256 位 ECC Key 在安全性上等同于 3072 位 RSA Key，加上 ECC 运算速度更快，ECDHE 密钥交换 + ECDSA 数字签名无疑是最好的选择。**由于同等安全条件下，ECC 算法所需的 Key 更短，所以 ECC 证书文件体积比 RSA 证书要小一些。**

5. 版本要求

> HTTP/2 要求 Nginx 1.9.5+，，OpenSSL 1.0.2+
>
> TLS 1.3 要求 Nginx 1.13+，OpenSSL 1.1.1+
>
> Brotli 要求 HTTPS，并在 Nginx 中添加扩展支持
>
> ECC 双证书 要求 Nginx 1.11+
>
> 这里 Nginx，我个人推荐 1.15+，因为 1.14 虽然已经能支持TLS1.3了，但是一些 TLS1.3 的进阶特性还只在 1.15+ 中提供。

6. 配置和测试详见：<https://yq.aliyun.com/articles/674714>

7. <https://www.jianshu.com/p/4a84c8388183> 有对https性能优化的一部分总结。

   1. 选择合适的算法

      - 证书校验：ECDSA>RSA
      - 非对称密钥交换：ECDHE_RSA>RSA
      - 对称加解密：2010年之后的Intel CPU支持AES-NI加密，可提升加解密速度

   2. 会话复用

      - 使用Session ID或者Session Ticket实现会话复用，减少握手次数。

   3. 将耗时的算法分离出去

      - 服务器安装专用的SSL/TLS硬件加速卡或者将消耗CPU大的算法转移到其它机器，降低CPU消耗。

   4. TLS False Start

      - **启用TLS False Start可以节省一个RTT时间**，但是客户端和服务端都需要支持NPN/ALPN，需要采用支持前向保密的密码套件，**即使用ECDHE进行密钥交换**。

   5. SPDY/HTTP2

      - SPDY/HTTP2通过多路复用技术减少了握手次数。



# 参考资料

1. [【协议】HTTPS性能优化实践](<https://zhuanlan.zhihu.com/p/25290538>) 主要讲腾讯网关如何从协议实现层优化。 [百度 HTTPS 性能优化经验](<https://www.infoq.cn/article/soKW3Lm9hoU4yUh2G81h>) 也有一些协议实现优化
2. [基于 Nginx 的 HTTPS 性能优化实践](<https://yq.aliyun.com/articles/674714>) 主要是基于软件、算法 类型和版本选型的优化，更具有可实施性。