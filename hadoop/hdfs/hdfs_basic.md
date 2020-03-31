# hdfs架构与原理

## 架构

1. 分为NameNode和DataNodes，NameNode管理文件系统元数据，DataNodes存储实际的数据。

2. NameNode执行文件系统命名空间的操作，比如，打开，关闭和重命名地址跟路径。它应该部署在可靠的硬件上。

3. 这些DataNodes节点才是真正相应客户端读写请求的节点。数据节点可以被部署在硬件架子上，但是不必非要部署在非常可靠地机器上。他们还根据来自namenode的指令进行创建、删除和复制。他们复制自身给其他的数据节点，并继续处理，直到规定数量的副本被创建完毕。

   

## 原理



## 特点

1. hdfs设计的原则是更趋向于少量的大文件，而不是大量的小文件。
2. 当任何文件写入HDFS中时，它都会被打散分成小块儿存储，这就是所谓的Block. HDFS给block设置了默认大小为128MB，同时这个大小可以根据需求而增加。这些blocks以分布式的方式存储在HDFS系统中的不同节点上。这样的就为MapReduce提供了一种机制，可以在集群中平行的产生数据。对于每个block的多重复制是横跨集群在不同节点下的复制。这些复制都是重复的数据。在默认情况下，HDFS的默认因子为3。这样的设置为系统提供了容错性，可靠性和高可依赖性。总结一下，大文件在HDFS中被分解成n个小的blocks.每个block以分布式的方式实现在集群上的存储，同时每个block对于自身的复制是夸集群进行的。
3. 由于Hadoop通常的是运行在一个集群中，这就对机架数量有了要求。NameNode将不同的Block复制在不同的机架上用以提高系统的容错性。NameNode尝试着每个机架上都复制一个自己，这样就算一个机架不工作了，系统仍然可以高效的运行。
4. 大量在集群中的datanodes被HDFS存储在本地磁盘里。DataNodes定时给发送心跳信息给Namenode，来确保它是否在线。它同时也给不同的DataNode传递自己身数据，直到复制因子达到3的时候完成。
5. 高吞吐量的访问数据：当我们想执行一个任务或者一个动作的时候，这个任务被分解给不同的系统。所以所有的系统都会被平行的执行它们自己部分。这就使工作会在很短的时间内完成。利用这点，HDFS给予了非常棒的吞吐量。平行读取数据的方式，使我们降低了读取超大数据所需的时间。
6. Hadoop 分布式文件系统和Linux文件系统有诸多相似之处。所以我们几乎可以将用于本地文件系统的所有操作命令都用到HDFS文件系统操作中来，比如，创建一个路径，复制一个文件，更改权限等等。
7. 我们可以通过用浏览器访问[http://master-IP:50070](https://link.jianshu.com/?t=http%3A%2F%2Fmaster-ip%3A50070%2F)去访问文件系统。
8. 考虑到安全和授权的目的，NameNode给客户端提供token，这个token需要出示给DateNote进行认证，认证通过后，才可以读取文件。
9. 读操作先通过接口向NameNode发送一个请求，NameNode检查该客户端是否有足够的权限去访问这组数据，然后NameNode节点在返回文件储存路径、用于权限检查的token分享给客户端。当该客户端去数据节点读文件的时候，在检查token之后，数据节点允许客户端读特定的block. 一个客户端打开一个输入流开始从DataNode读取数据，然后，客户端直接从数据节点读取数据。如果在读取数据期间datanodes突然废了，这个客户端会继续访问Namenode, 这是NameNode会给出这个数据节点副本的所在位置。
10. 写文时客户端与NameNode进行交互，NameNode需要提供可供写入数据的DataNode节点的地址、用于权限检查的token。当客户端完成写入第一个block时，第一个数据节点会复制一样的block给另一个DataNode, 然后在这个数据节点完成接收block之后，它开始复制这些blocks给第三个数据节点。第三个数据节点发送通知给第二个数据节点，第二个数据节点在发送通知给第一个数据节点，第一个数据节点负责最后通知客户端，告知数据写入已经完成。不论复制因子是多少，客户端只发送一个数据副本给DataNode, DataNode完成后续所有任务的复制工作。所以，在Hadoop中写入文件并不十分消耗系统资源，因为它可以在多个数据点将blocks平行写入。

# hdfs基本操作

1. 如：

- `hadoop fs -ls /`
- `hadoop fs -lsr /user`
- `rm(r)、mkdir、put、get`

2. 详见：[Hadoop框架之HDFS的shell操作](https://www.cnblogs.com/hezhiyao/p/7627060.html)
3. hadoop dfs、hadoop fs、hdfs dfs的区别：

- Hadoop  fs：使用面最广，可以操作任何文件系统。
> FS relates to a generic file system which can point to any file systems like local, HDFS etc. 
So this can be used when you are dealing with different file systems such as Local FS, HFTP FS, S3 FS, and others

- hdfs    dfs：只能操作HDFS文件系统相关（包括与Local FS间的操作）
- hadoop  dfs:<已经Deprecated>, 不推荐(deprecated), 不支持(Discontinued)

## 常见操作：

```shell
$ hadoop fs -help # 查看帮助
$ hadoop fs -ls / # 注意是查看HDFS上的目录，而不是本地目录。
Found 3 items
drwxr-xr-x   - unown supergroup          0 2019-04-04 16:00 /data
drwxr-xr-x   - unown supergroup          0 2019-01-18 15:27 /historyforspark
drwxr-xr-x   - unown supergroup          0 2019-04-06 11:00 /tmp
$ ls /
bin   data   data1.lock   dev  home  lib64       media        mnt  proc    redis  run   src  sys  udisk02  var
boot  data1  dataservice  etc  lib   lost+found  MegaSAS.log  opt  recall  root   sbin  srv  tmp  usr
$ hadoop fs -ls -R /
$ hadoop fs -du -s /data
11950109056  /data
$ hadoop fs -du -sh /data
-du: Illegal option -sh
Usage: hadoop fs [generic options] -du [-s] [-h] <path> ...
$ hadoop fs -du -s -h /data
11.1 G  /data
$ hadoop fs -tail -10 /data/app.conf
-tail: Illegal option -10
Usage: hadoop fs [generic options] -tail [-f] <file>
$ hadoop fs -tail /data/app.conf
# dag service port
dag_service_port = 7777

# 镜像仓库节点
image_ip_node = 127.0.0.1:5000

# 数据库保留时长
backup_time  = 90
$ hdfs getconf -namenodes
cpm10705
$ hdfs getconf -secondaryNameNodes
0.0.0.0
$ hdfs dfs -mkdir /tmp/test
$ hdfs dfs -ls /tmp/
Found 3 items
drwx------   - unown supergroup          0 2019-04-06 11:00 /tmp/hadoop-yarn
drwxrwxrwt   - unown supergroup          0 2019-03-22 14:27 /tmp/logs
drwxr-xr-x   - unown supergroup          0 2019-04-06 12:31 /tmp/test
$ hdfs dfs -put wc.input /tmp/test/wc.input
$ hdfs dfs -ls /tmp/test/
Found 1 items
-rw-r--r--   1 unown supergroup         71 2019-04-06 12:33 /tmp/test/wc.input
$ hdfs dfsadmin -report # 报告HDFS的基本统计信息
Configured Capacity: 1968321044480 (1.79 TB)
Present Capacity: 775970242560 (722.68 GB)
DFS Remaining: 763896565760 (711.43 GB)
DFS Used: 12073676800 (11.24 GB)
DFS Used%: 1.56%
Under replicated blocks: 2690
Blocks with corrupt replicas: 0
Missing blocks: 0
Missing blocks (with replication factor 1): 0

-------------------------------------------------
Live datanodes (1):

Name: 10.7.0.5:50010 (cpm10705)
Hostname: cpm10705
Decommission Status : Normal
Configured Capacity: 1968321044480 (1.79 TB)
DFS Used: 12073676800 (11.24 GB)
Non DFS Used: 1192350801920 (1.08 TB)
DFS Remaining: 763896565760 (711.43 GB)
DFS Used%: 0.61%
DFS Remaining%: 38.81%
Configured Cache Capacity: 0 (0 B)
Cache Used: 0 (0 B)
Cache Remaining: 0 (0 B)
Cache Used%: 100.00%
Cache Remaining%: 0.00%
Xceivers: 1
Last contact: Sat Apr 06 12:42:45 CST 2019

$ hdfs dfs -ls /tmp/test
Found 1 items
-rw-r--r--   1 unown supergroup         71 2019-04-06 12:33 /tmp/test/wc.input
$ hdfs dfs -rm  -r -f /tmp/test/wc.input
19/04/06 12:44:33 INFO fs.TrashPolicyDefault: Namenode trash configuration: Deletion interval = 0 minutes, Emptier interval = 0 minutes.
Deleted /tmp/test/wc.input
$ hdfs dfs -ls /tmp/test


```

### 推拉数据

```shell
$ hadoop fs -put -f ./conf/app.conf hdfs://10.4.1.131:9000/data/app.conf                         
$ hadoop fs -get  hdfs://10.4.1.131:9000/data/app.conf ./app.conf
```



## hadoop集群Restful接口的使用

1. 详见：

- [hadoop集群Restful接口的使用](https://blog.csdn.net/u011518678/article/details/48133901)
- [HTTPFS: 基于HTTP操作hadoop hdfs文件系统](https://my.oschina.net/cloudcoder/blog/277426)
- <http://hadoop.apache.org/docs/r2.6.0/hadoop-project-dist/hadoop-hdfs/WebHDFS.html>

# 参考资料

- [Hadoop HDFS 教程（一）介绍](https://www.jianshu.com/p/8969eb90a59d)

- [Hadoop框架之HDFS的shell操作](https://www.cnblogs.com/hezhiyao/p/7627060.html)

