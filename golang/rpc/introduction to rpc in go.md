# 为什么要RPC
1. 简单来说成熟的rpc库相对http容器，更多的是封装了“服务发现”，"负载均衡"，“熔断降级”一类面向服务的高级特性。可以这么理解，rpc框架是面向服务的更高级的封装。如果把一个http servlet容器上封装一层服务发现和函数代理调用，那它就已经可以做一个rpc框架了。
2. HTTP是一种应用层网络协议，RPC可以采用自定义协议，也可以通过HTTP协议来传输，thrift，grpc，xml-rpc，json-rpc都是通过HTTP传输的。HTTP既支持长连接，也支持短连接。


# 参考文献
- [既然有 HTTP 请求，为什么还要用 RPC 调用？](https://www.zhihu.com/question/41609070)
- [4.3 玩转RPC](https://chai2010.gitbooks.io/advanced-go-programming-book/content/ch4-rpc/ch4-03-netrpc-hack.html)
- [Go语言rpc方案调研](https://scguoi.github.io/DivisionByZero/2016/11/15/GO%E8%AF%AD%E8%A8%80RPC%E6%96%B9%E6%A1%88%E8%B0%83%E7%A0%94.html)
- [序列化和反序列化](https://www.infoq.cn/article/serialization-and-deserialization)