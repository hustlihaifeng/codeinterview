[TOC]

# docker实现原理

1. 使用Namespaces实现了系统环境的隔离，Namespaces允许一个进程以及它的子进程从共享的宿主机内核资源（网络栈、进程列表、挂载点等）里获得一个仅自己可见的隔离区域，让同一个Namespace下的所有进程感知彼此变化，对外界进程一无所知，仿佛运行在一个独占的操作系统中；
2. 使用CGroups限制这个环境的资源使用情况，比如一台16核32GB的机器上只让容器使用2核4GB。使用CGroups还可以为资源设置权重，计算使用量，操控任务（进程或线程）启停等；
3. 使用镜像管理功能，利用Docker的镜像分层、写时复制、内容寻址、联合挂载技术实现了一套完整的容器文件系统及运行环境，再结合镜像仓库，镜像可以快速下载和共享，方便在多环境部署。

# docker容器隔离的问题

1. Docker不像虚机虚拟化一个Guest OS，而是利用宿主机的资源，和宿主机共用一个内核，所以会存在下面问题：
   1. Docker是利用CGroups实现资源限制的，**只能限制资源消耗的最大值，而不能隔绝其他程序占用自己的资源**;
   2. Namespace的6项隔离看似完整，实际上依旧没有完全隔离Linux资源，比如/proc 、/sys 、/dev/sd*等目录未完全隔离，SELinux、time、syslog等所有现有Namespace之外的信息都未隔离。
      1. 在Docker容器中执行 top、free 等命令，会发现看到的资源使用情况都是宿主机的资源情况，而我们需要的是这个容器被限制了多少CPU，内存，当前容器内的进程使用了多少；
      2. 程序运行在容器里面，调用API获取系统内存、CPU，取到的是宿主机的资源大小；
      3. 对于多进程程序，一般都可以将worker数量设置成auto，自适应系统CPU核数，但在容器里面这么设置，取到的CPU核数是不正确的，例如Nginx，其他应用取到的可能也不正确，需要进行测试。
      4. 这些问题的本质都一样，在Linux环境，很多命令都是通过读取/proc 或者 /sys 目录下文件来计算资源使用情况。

# **从CGroups中读取**

1. Docker 在 1.8 版本以后会将分配给容器的CGroups信息挂载进容器内部，容器里面的程序可以通过解析CGroups信息获取到容器资源信息。

2. > **1、读取容器CPU核数**
   >
   > ```
   > # 这个值除以100000得到的就是容器核数
   > ~ # cat  /sys/fs/cgroup/cpu/cpu.cfs_quota_us 
   > 400000复制代码
   > ```
   >
   > **2、获取容器内存使用情况（USAGE / LIMIT）**
   >
   > ```
   > ~ # cat /sys/fs/cgroup/memory/memory.usage_in_bytes 
   > 4289953792
   > ~ # cat /sys/fs/cgroup/memory/memory.limit_in_bytes 
   > 4294967296复制代码
   > ```
   >
   > 将这两个值相除得到的就是内存使用百分比。
   >
   > **3、获取容器是否被设置了OOM，是否发生过OOM**
   >
   > ```
   > ~ # cat /sys/fs/cgroup/memory/memory.oom_control 
   > oom_kill_disable 0
   > under_oom 0
   > ~ #
   > ~ #复制代码
   > ```
   >
   > 这里需要解释一下：
   >
   > - oom_kill_disable默认为0，表示打开了oom killer，就是当内存超时会触发kill进程。可以在使用docker run时候指定disable oom，将此值设置为1，关闭oom killer；
   > - under_oom 这个值仅仅是用来看的，表示当前的CGroups的状态是不是已经oom了，如果是，这个值将显示为1。
   >
   > **4、获取容器磁盘I/O**
   >
   > ```
   > ~ # cat /sys/fs/cgroup/blkio/blkio.throttle.io_service_bytes
   > 253:16 Read 20015124480
   > 253:16 Write 24235769856
   > 253:16 Sync 0
   > 253:16 Async 44250894336
   > 253:16 Total 44250894336
   > Total 44250894336复制代码
   > ```
   >
   > **5、获取容器虚拟网卡入/出流量**
   >
   > ```
   > ~ # cat /sys/class/net/eth0/statistics/rx_bytes 
   > 10167967741
   > ~ # cat /sys/class/net/eth0/statistics/tx_bytes 
   > 15139291335
   > ~ #复制代码
   > ```
   >
   > 

# 参考资料

1. [Docker容器实现原理及容器隔离性踩坑介绍](<https://juejin.im/post/5d2d9083e51d45777b1a3e5d>)