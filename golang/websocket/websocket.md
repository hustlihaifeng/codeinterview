# 1. 简单的回显服务

## 1.1 服务端

1. `http.Handle`将url和对应的处理函数联系起来
2. `http.ListenAndServe`在指定地址上进行监听
3. 处理函数对一个`websocket.Conn`对象进行处理，如回显服务就是`io.Copy(ws, ws)`。

## 1.2 客户端

1. `websocket.Dial`新建一个websocket连接，需要指定服务端地址，如：`ws://localhost:12345/echo`，本地地址`http://localhost/`。
2. 然后通过`ws.Write`、`ws.Read`来对websocket进行读写。

## 1.3 测试

1. <http://coolaf.com/tool/chattest> 上面ws://xxx.xxx.xxx.xxx:xx/v1/ws/test/1进行测试。
2. chrome有个websocket测试插件：<https://chrome.google.com/webstore/detail/smart-websocket-client/omalebghpgejjiaoknljcfmglgbpocdp?utm_source=chrome-app-launcher-info-dialog>

## 1.4 问题

1. 客户端发送close后，服务端没有监听和处理，导致接下来服务端发送Heartbeat时出现broken pipe。或许需要注册一个close函数，来进行清理。
2. 短时间进行断开，连接操作。客户端报错

```shell
发生错误: undefined
websocket连接已断开!!!
```

然后导致服务端写的时候发生broken pipe。

# 2. websocket官方文档

## 2.1 注意事项

- 详见<https://godoc.org/github.com/gorilla/websocket>、[websocket标准：rfc6455](https://godoc.org/github.com/gorilla/websocket)

1. 应用需要保证text message是utf-8编码的
2. close ping peng三种**control message**，WriteControl, WriteMessage or NextWriter都可以用来写control message。NextReader, ReadMessage or the message Read method时会调用control handler，默认的close和ping handler会等待一段时间来让所发送的消息到达对方（并不知道到是否确保）。应用应该读connection来确保control消息得到处理。
3. NextReader, ReadMessage or the message Read method在遇到close message时返回一个`*CloseError`，默认的CloseHandler 发送一个close message给对端（peer）。

> The default close and ping handlers can block these methods for a short time when the handler writes to the connection.

4. 默认的pong handler什么都不做，如果应用发送ping message，那么应用应该设置一个pong handler(比如给客户端发送心跳，超时没响应后踢掉该客户端)
5. **应用应该确保只有一个goroutine并发的来读或者写链接。Close 和WriteControl函数可以和其他所有函数一起并发的被调用。**
6. 如果CheckOrigin 函数返回false，则对端收到http status 403

> If the CheckOrigin field is nil, then the Upgrader uses a safe default: fail the handshake if the Origin request header is present and the Origin host is not equal to the Host request header.

7. 可以设置读写缓冲区大小,默认4096。缓冲区大小不限制消息大小上限

> The buffer sizes in bytes are specified by the ReadBufferSize and WriteBufferSize fields in the Dialer and Upgrader.



8. `func FormatCloseMessage(closeCode int, text string) []byte`可以用来构建close message
9. `func (c *Conn) Close() error`函数**直接关闭链接，并不写或者等待close message**。
10. 一个链接至多有一个open reader，只多一个open wriate

> There can be at most one open reader on a connection. NextReader discards the previous message if the application has not already consumed it.

> There can be at most one open writer on a connection. NextWriter closes the previous writer if the application has not already done so.

11. 链接读失败后，后面会一直返回读失败。也即一次读失败后就该终止链接。
12. 各种消息都可以通过writer写：`TextMessage, BinaryMessage, CloseMessage, PingMessage and PongMessage`
13. `SetReadDeadline ` `SetWriteDeadline `设置链接超时时间，一个超时之后**all future**读写会失败。

> After a read has timed out, the websocket connection state is corrupt and all future reads will return an error. A zero value for t means reads will not time out.

14. `func (c *Conn) WriteControl(messageType int, data []byte, deadline time.Time) error`用来写control消息

> WriteControl writes a control message with the given deadline. The allowed message types are CloseMessage, PingMessage and PongMessage.

15. 用`IsUnexpectedCloseError`来检测客户端异常退出，没有返回close，详见<https://godoc.org/github.com/gorilla/websocket#example-IsUnexpectedCloseError>：

```go
for {
    messageType, p, err := c.ReadMessage()
    if err != nil {
        if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
            log.Printf("error: %v, user-agent: %v", err, req.Header.Get("User-Agent"))
        }
        return
    }
    processMesage(messageType, p)
}
```



## 2.2 官方样例学习

1. 地址：<https://godoc.org/github.com/gorilla/websocket>

- 右上角有四个链接，分别是：
  - <https://godoc.org/github.com/gorilla/websocket#pkg-index>: 函数列表
  - <https://godoc.org/github.com/gorilla/websocket#pkg-examples>：只讲了一个例子<https://godoc.org/github.com/gorilla/websocket#pkg-examples>
  - <https://godoc.org/github.com/gorilla/websocket#pkg-files>：里面讲了关闭类型常量、消息类型常量、错误类型常量
  - <https://godoc.org/github.com/gorilla/websocket#pkg-subdirectories>：里面有几个官方样例

### 1. 聊天室

1. 代码见：<https://godoc.org/github.com/gorilla/websocket/examples/chat>

2. 通信模式：由于`websocket`的reader和writer并发能力都为1，所以需要一个单独的reader协程来读，一个单独的writer协程来写。为了防止并发问题，元数据的修改统一通过channel发送到manager。reader协程的数据通过manager的channel发送到manager，manager处理后将数据发送到writer的channel里面，然后channel发送到`websocket`对端。

3. 登录登出

   - 登录时：新建client对象，并拉起读写协程，然后将本client对象发送到master的register channel中。manager读取到后，将该client对象加入到元数据map中。

   - 登出时：reader协程得到CloseError，发送本client对象到manager的unregister channel后关闭本websocket链接。manager从unregister channel中收到该对象后，从元数据map中删除该对象，然后关闭该对象的writer channel。
     - 对端异常关闭：reader routine里面有设置ReadDeadline为当前时间加上pong消息等待时间，然后在writer线程中定期发送ping消息。正常情况下，定时收到pong消息，不断更新ReadDeadline，不会有问题。当对端异常关闭时，pong消息收不到，reader的读channel会检测到ReadDeadline超时，然后发生错误，进而触发reader 协程退出逻辑。**reader协程的退出逻辑就是登出的逻辑**。对端异常关闭时，还可能发送close消息发送失败，此时writer线程需要检测发送失败消息和写超时消息，任何一个都需要进入退出逻辑（关闭链接）。
     - 主动关闭链接：关闭writer channel，writer 协程检测到writer channel关闭事件，就结束。writer的关闭事件即主动关闭链接的操作：向对端发送close消息，然后关闭ping channel。这样不一会儿reader 协程就会检测到ReadDeadline，进而触发到读协程的关闭事件：登出。
     - 读写协程的关闭事件都需要关闭websocket链接，写协程多了关闭ping channel，读协程多了向manager的unregister channel发送本client对象。manster的unregister响应需要删除元数据并关闭writer channel。两个协程由路由响应函数拉起即可，不必放到manager中，来提高manager效率。