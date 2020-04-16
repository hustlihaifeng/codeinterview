[TOC]

# 1. 方案简介

## 1.1 服务热更新

1. 一架行驶在高速上的大卡车，行驶过程中突然遭遇爆胎，热更新则是要求在不停车的情况下将车胎修补好，且补胎过程中卡车需要保持正常行驶

## 1.2 nginx方案

1. nginx方案

> - 1）首先备份原有的Nginx二进制文件，并用新编译好的Nginx二进制文件替换旧的
> - 2）然后向master进程发送`USR2`信号.此时Nginx进程会启动一个新版本Nginx，该新版本Nginx进程会发起一个新的master进程与work进程.即此时会有两个Nginx实例在运行，一起处理新来的请求.(lhf：请求依然发送到老的nginx端口，新的nginx实例如何处理新的请求?答案是多个进程监听同一个ip端口，这样问题是每个进程都会受到并处理请求，nginx解决办法是fork后没有获取到锁的进程将从父进程中继承的socket连接去除掉，多个进程会轮流的去获取锁来成为当前工作线程。（lhf：简单实现就是在某个固定路径文件总写入当前pid，工作时判断与当前进程pid相等就工作，配合定时器啥的）)
> - 3）再向原master进程发送`WINCH`信号，它会逐渐关闭相关work进程，此时原master进程仍保持监听新请求但不会发送至其下work进程，而是交给新的work进程。（lhf:根据前面的锁轮询方案，那么此时应该是给原来的工作进程设置一个状态，然后这个状态下的进程不去获取锁）
> - 4）最后等到所有原work进程全部关闭，向原master进程发送`QUIT`信号，终止原master进程，至此，完成Nginx热升级.

![](nginx_architecture.jpg)

![](nginx_architecture_admin.jpeg)

- 多线程在多并发情况下，线程的**内存占用大，线程上下文切换造成CPU大量的开销**
- 与Memcached的经典多线程模型相比，Nginx是经典的多进程模型。 
- TODO：源码级别学习nginx，高负载就看这个了。
- 关于多个socket端口复用：

> **设置socket的SO_REUSEADDR选项，即可实现端口复用：
>
> ```c
> int opt = 1;
> // sockfd为需要端口复用的套接字
> setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, (const void *)&opt, sizeof(opt));
> ```
>
> 1、当有一个有相同本地地址和端口的socket1处于TIME_WAIT状态时，而你启动的程序的socket2要占用该地址和端口，你的程序就要用到该选项。
>
> 2、SO_REUSEADDR允许同一port上启动同一服务器的多个实例(多个进程)。但每个实例绑定的IP地址是不能相同的。在有多块网卡或用IP Alias技术的机器可以测试这种情况。
>
> 3、SO_REUSEADDR允许单个进程绑定相同的端口到多个socket上，但每个socket绑定的ip地址不同。这和2很相似，区别请看UNPv1。
>
> 4、SO_REUSEADDR允许完全相同的地址和端口的重复绑定。但这只用于UDP的多播，不用于TCP。
>
> 端口复用允许在一个应用程序可以把 n 个套接字绑在一个端口上而不出错。同时，这 n 个套接字发送信息都正常，没有问题。但是，这些套接字并不是所有都能读取信息，只有最后一个套接字会正常接收数据。
>
> 我感觉你这个例子，只是说明在设置so_reuseaddr下，使用不同的传输层协议情况下对相同IP和PORT的复用。 了解一下so_reuseport，感觉这个设置才是真正的端口的复用。不过，so_reuseaddr在不同的操作系统下，表现的行为还不太一样

- [网络编程中的SO_REUSEADDR和SO_REUSEPORT参数详解](<https://zhuanlan.zhihu.com/p/35367402>) 这里面说`SO_REUSEADDR`主要用于复用`time_wait`装填的socket，而`SO_REUSEPORT`才是真正的复用port。



2. nginx配置更新和服务热升级,详见：[nginx启动、重启、重新加载配置文件和平滑升级](<https://blog.csdn.net/gnail_oug/article/details/52754491>)

   1. nginx配置更新：`nginx -s reload`  或者  `kill -HUP 主进程号`

      - 主进程好即master进程号

      ```shell
      [root@localhost sbin]# ps -ef|grep nginx
      root       9944      1  0 13:22 ?        00:00:00 nginx: master process ./nginx
      nobody     9949   9944  0 13:23 ?        00:00:00 nginx: worker process
      root       9960   9917  0 13:28 pts/1    00:00:00 grep nginx
      [root@songguoliang sbin]# kill -HUP 9944
      ```

   2. nginx平滑升级：

      1. 用新的nginx可执行程序替换旧的可执行程序，即下载新的nginx，重新编译到旧版本的安装路径中(重新编译之前可以备份旧的可执行文件)
      2. 给nginx主进程号发送USR2信号

      ```shell
      [root@localhost sbin]# ps -ef |grep nginx
      root       9944      1  0 13:22 ?        00:00:00 nginx: master process ./nginx
      nobody     9965   9944  0 13:29 ?        00:00:00 nginx: worker process
      root      10010   9917  0 13:42 pts/1    00:00:00 grep nginx
      [root@localhost sbin]# kill -USR2 9944
      [root@localhost sbin]# ps -ef |grep nginx
      root       9944      1  0 13:22 ?        00:00:00 nginx: master process ./nginx
      nobody     9965   9944  0 13:29 ?        00:00:00 nginx: worker process
      root      10012   9944  0 13:43 ?        00:00:00 nginx: master process ./nginx
      nobody    10013  10012  0 13:43 ?        00:00:00 nginx: worker process
      root      10015   9917  0 13:43 pts/1    00:00:00 grep nginx
      
      ```

       	3. 给旧的主进程发送WINCH信号，kill -WINCH 旧的主进程号

      ```shell
      [root@localhost sbin]# ps -ef |grep nginx
      root       9944      1  0 13:22 ?        00:00:00 nginx: master process ./nginx
      nobody     9965   9944  0 13:29 ?        00:00:00 nginx: worker process
      root      10012   9944  0 13:43 ?        00:00:00 nginx: master process ./nginx
      nobody    10013  10012  0 13:43 ?        00:00:00 nginx: worker process
      root      10092   9917  0 14:05 pts/1    00:00:00 grep nginx
      [root@localhost sbin]# kill -WINCH 9944
      [root@localhost sbin]# 
      [root@localhost sbin]# 
      [root@localhost sbin]# 
      [root@localhost sbin]# ps -ef |grep nginx
      root       9944      1  0 13:22 ?        00:00:00 nginx: master process ./nginx
      root      10012   9944  0 13:43 ?        00:00:00 nginx: master process ./nginx
      nobody    10013  10012  0 13:43 ?        00:00:00 nginx: worker process
      root      10094   9917  0 14:06 pts/1    00:00:00 grep nginx
      ```

      旧的主进程号收到WINCH信号后，将旧进程号管理的旧的工作进程优雅的关闭。即一段时间后旧的工作进程全部关闭，只有新的工作进程在处理请求连接。这时，依然可以恢复到旧的进程服务，因为旧的进程的监听socket还未停止。

      4. 给旧的主进程发送QUIT信号，使其关闭。

      ```shell
      [root@localhost sbin]# kill -QUIT 9944
      [root@localhost sbin]# ps -ef |grep nginx
      root      10012      1  0 13:43 ?        00:00:00 nginx: master process ./nginx
      nobody    10013  10012  0 13:43 ?        00:00:00 nginx: worker process
      root      10118   9917  0 14:16 pts/1    00:00:00 grep nginx
      ```

      给旧的主进程发送QUIT信号后，旧的主进程退出，并移除logs/nginx.pid.oldbin文件，nginx的升级完成。

## 1.3 golang实现类似nginx的服务热更新

### 1.3.1  流程

1. 备份老文件，替换新文件。向服务发送USR2信号。
2. 进程收到`USR2`信号后，fork子进程（启动新版本服务，并和老进程监听在同一端口），并将当前socket句柄等进程环境交给它。
3. 新进程开始监听socket请求。
4. 让老的进程停止接受新的请求。
5. 老进程停止工作。

### 1.3.2 实现

```go
    func (a *app) signalHandler(wg *sync.WaitGroup) {
        ch := make(chan os.Signal, 10)
        signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
        for {
            sig := <-ch
            switch sig {
            case syscall.SIGINT, syscall.SIGTERM:
                // 确保接收到INT/TERM信号时可以触发golang标准的进程终止行为
                signal.Stop(ch)
                a.term(wg)
                return
            case syscall.SIGUSR2:
                err := a.preStartProcess()
                if err != nil {
                    a.errors <- err
                }
                // 发起新进程
                if _, err := a.net.StartProcess(); err != nil {
                    a.errors <- err
                }
            }
        }
    }

// 复制当前进程socket连接，发起新进程
    execSpec := &syscall.ProcAttr{
    Env: os.Environ(),
    Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
        // 其他socket，由老进程处理，新的由新的socket处理，那么此时需要客户端重新连接socket才能使用新的服务。最好是在父进程对某个连接处理完某个阶段后，接下来由子进程进行处理。
        // 如果想做到客户端无感知，那么需要将老的socket一并传递给子进程，子进程进行接下来的处理。
    }
    fork, err := syscall.ForkExec(os.Args[0], os.Args, execSpec)
    ...
```

## 1.4 配置文件更新

1. 与本方案类似，子进程重新读取配置文件即可。

## 1.5 子进程pid变化，可能导致systemd & supervisor这类监控进程失败

1. 使用pidfile，每次进程重启更新一下pidfile，让进程管理者通过这个文件感知到main pid的变更。
2. 更通用的做法：**起一个master来管理服务进程**，每次热重启master拉起一个新的进程，把旧的kill掉。这时master的pid没有变化，对于进程管理者来说进程处于正常的状态。[一个简洁的实现](https://link.juejin.im/?target=https%3A%2F%2Fgithub.com%2Fkuangchanglang%2Fgraceful)

## 1.6 TODO：UDP服务的热更新呢？有什么不同么

# 成熟的开源实现

1. `github.com/facebookgo/grace/gracehttp`

```go
    func main() {
        app := gin.New()// 项目中时候的是gin框架
        router.Route(app)
        var server *http.Server
        server = &http.Server{
            Addr:    ":8080",
            Handler: app,
        }
        gracehttp.Serve(server)
    }
```

2. [4.2 grace 模块](<http://www.kancloud.cn:8080/hello123/beego/126136> )
   1. [beego高级编程---->grace模块热重启导致旧进程未处理完请求直接退出](<https://studygolang.com/articles/3233>)
3. TODO:实战测试

# 参考资料

1. [golang hotfix热更新详解](<https://mojotv.cn/2018/12/26/golang-hot-restart>)
2. [Golang服务器热重启、热升级、热更新(safe and graceful hot-restart/reload http server)详解](https://www.cnblogs.com/sunsky303/p/9778466.html)
3. [nginx架构模型分析](<https://juejin.im/post/5cdea826e51d4510b934dcb5>)