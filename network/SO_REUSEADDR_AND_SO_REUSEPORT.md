# `SO_REUSEADDR`和`SO_REUSEPORT`的作用

1. BSD socket实现是其他socket实现的鼻祖，由于fork BSD实现的不同，不同系统对`SO_REUSEADDR`和`SO_REUSEPORT`的支持不同。

2. 一个`TCP/UDP` 连接由：`{<protocol>, <src addr>, <src port>, <dest addr>, <dest port>}`这五元组标识，如果完全相同的话，回报错`EADDRINUSE`。

3. `SO_REUSEADDR`:影响TCP连接的`TIME_WAIT`收尾阶段的链接，如果设置新的链接设置了`SO_REUSEADDR`,那么可以重用处于`TIME_WAIT`阶段的链接（五元组标识）。

4. `SO_REUSEPORT`:如果涉及的所有连接都设置了`SO_REUSEPORT`，那么他们可以共用同一对`<src addr>, <src port>`。

5. `SO_REUSEPORT`和`SO_REUSEADDR`相互不影响，`SO_REUSEADDR`比`SO_REUSEPORT`更早实现，有更多的支持。

# 系统支持情况

1. `macOS`基本上与`BSD`相同
2. `Linux < 3.9`：只有`SO_REUSEADDR`存在, 行为基本与BSD相同，除了一下两点：

- 在client端，`Linux < 3.9`的`SO_REUSEADDR`表现地和`SO_REUSEPORT`一样。因为3.9之前没有`SO_REUSEPORT`，

  > it is important to be able to bind multiple sockets to exactly to the same UDP socket address for various protocols。

- 在server端

  > As long as a listening (server) TCP socket is bound to a specific port, the `SO_REUSEADDR` option is entirely ignored for all sockets targeting that port. Binding a second socket to the same port is only possible if it was also possible in BSD without having `SO_REUSEADDR` set. 如不同的本地ip。

3. `Linux >= 3.9`： 添加了`SO_REUSEPORT`，含义与`BSD`相同。不同点是：

- 共用`<src addr>, <src port>`的进程必须属于同一用户：

> **All sockets that want to share the same address and port combination must belong to processes that share the same effective user ID!**

> Additionally the kernel performs some "special magic" for `SO_REUSEPORT` sockets that isn't found in other operating systems: For UDP sockets, it tries to distribute datagrams evenly, for TCP listening sockets, it tries to distribute incoming connect requests (those accepted by calling `accept()`) evenly across all the sockets that share the same address and port combination. **Thus an application can easily open the same port in multiple child processes and then use `SO_REUSEPORT` to get a very inexpensive load balancing.**

4. android与Linux相同。
5. windows：只有`SO_REUSEADDR`,设置了`SO_REUSEADDR`后的表现与BSD中同时设置`SO_REUSEPORT`、`SO_REUSEADDR`一样。

> A socket with `SO_REUSEADDR` can always bind to exactly the same source address and port as an already bound socket, **even if the other socket did not have this option set when it was bound**.

这样有安全问题，微软添加了`SO_EXCLUSIVEADDRUSE`来解决这个问题，

> Setting `SO_EXCLUSIVEADDRUSE` on a socket makes sure that if the binding succeeds, the combination of source address and port is owned exclusively by this socket and no other socket can bind to them, not even if it has `SO_REUSEADDR` set.


# 参考资料

- <https://stackoverflow.com/a/14388707>
