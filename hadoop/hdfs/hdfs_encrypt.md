# setup

```shell
docker run -it registry.in.wezhuiyi.com/algorithm-platform/hadoop-docker:2.7.1 /etc/bootstrap.sh -bash
# $HADOOP_PREFIX
echo $HADOOP_PREFIX
/usr/local/hadoop
cd $HADOOP_PREFIX
# 测试
bin/hadoop jar share/hadoop/mapreduce/hadoop-mapreduce-examples-2.7.1.jar grep input output 'dfs[a-z.]+'
bin/hdfs dfs -ls output
bin/hdfs dfs -ls /user/root/input
ls etc/hadoop/
cat etc/hadoop/core-site.xml
cat etc/hadoop/hdfs-site.xml
bin/hadoop dfs -ls hdfs://10.126.1.3:9000/user/root/input
```

> /
> Starting sshd:                                             [  OK  ]
> 20/02/16 20:37:35 WARN util.NativeCodeLoader: Unable to load native-hadoop library for your platform... using builtin-java classes where applicable
> Starting namenodes on [c8a95ec17ed9]
> c8a95ec17ed9: starting namenode, logging to /usr/local/hadoop/logs/hadoop-root-namenode-c8a95ec17ed9.out
> localhost: starting datanode, logging to /usr/local/hadoop/logs/hadoop-root-datanode-c8a95ec17ed9.out
> Starting secondary namenodes [0.0.0.0]
> 0.0.0.0: starting secondarynamenode, logging to /usr/local/hadoop/logs/hadoop-root-secondarynamenode-c8a95ec17ed9.out
> 20/02/16 20:37:50 WARN util.NativeCodeLoader: Unable to load native-hadoop library for your platform... using builtin-java classes where applicable
> starting yarn daemons
> starting resourcemanager, logging to /usr/local/hadoop/logs/yarn--resourcemanager-c8a95ec17ed9.out
> localhost: starting nodemanager, logging to /usr/local/hadoop/logs/yarn-root-nodemanager-c8a95ec17ed9.out#

# 正式

```shell
docker run -p 9000:9000 -it registry.in.wezhuiyi.com/algorithm-platform/hadoop-docker:2.7.1 /etc/bootstrap.sh -bash
```

```shell
hadoop dfs -ls hdfs://10.4.1.71:9000/
```

