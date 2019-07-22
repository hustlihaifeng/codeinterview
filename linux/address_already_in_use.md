# 问题

1. 拉起server的时候，报错：`listen tcp :4444: bind: address already in use`

# 原因

1. 4444端口被另一个`tcp`程序占用，我们需要找出这个程序，要么解除原程序，要么更换端口。

# 解决办法

1. 找出老程序

```shell
# lsof -i :4444
COMMAND     PID    USER   FD   TYPE    DEVICE SIZE/OFF NODE NAME
xxxx-serv 24098 xxxx    7u  IPv6 239782639      0t0  TCP *:4444 (LISTEN)
# pwdx 24098
24098: /path/to/xxxx-serv

```

一般这个`/path/to/xxxx-serv`是同一个程序的老进程残留，`kill -9 24098`删除掉即可。否则，需要更换新程序端口。