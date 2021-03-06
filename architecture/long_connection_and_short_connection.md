[TOC]

# 1. 长链接和短连接优缺点和应用场景

## 1.1 长链接优缺点

1. 长连接可以省去较多的TCP建立和关闭的操作，减少浪费，节约时间，但是一直连接对于客户端来说比较耗电。
2. 客户端与服务端之间的连接如果一直不关闭，随着客户端连接越来越多，server早晚有扛不住的时候，这时候server端需要采取一些策略，如关闭一些长时间没有读写事件发生的连接，这样可以避免一些恶意连接导致服务端服务受损；

## 1.2 短连接优缺点

1. 短连接对于服务器来说管理较为简单，存在的连接都是有用的连接，不需要额外的控制手段。
2. 一次TCP连接和断开需要7个来回，如果客户端请求频繁，将在TCP的建立和关闭操作上浪费大量时间和带宽。

## 1.3 长短链接应用场景

1. 长链接: 频繁请求资源场景，双向实时通信场景（如websocket），一般数据库连接。smtp,pop3,telnet这种就可以认为是长连接。一般的网络游戏应用都是长连接
2. 短连接：普通的web服务中请求一个网页；每当浏览器每遇到一个Web资源（html文件、JavaScript文件、图像文件、CSS文件），就会建立一个HTTP会话（http是无状态的；http1.1增加了持久链接支持，可以显示的指定keep-alive，如果浏览器或者服务器在其头信息加入了这行代码 Connection:keep-alive
   TCP连接在发送后将仍然保持打开状态，于是，浏览器可以继续通过相同的连接发送请求）。http 1.0一般就指短连接。


# 参考资料

1. [TCP长连接和短链接的区别及应用场景](https://blog.csdn.net/qq_16635171/article/details/104312443)
2. [网络连接中的长连接和短链接是什么意思?](<https://www.zhihu.com/question/22677800/answer/63806275>)