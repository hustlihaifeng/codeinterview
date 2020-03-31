# 搭建CDH

1. 

```shell
docker pull cloudera/quickstart:latest
docker run --hostname=quickstart.cloudera --privileged=true -t -i cloudera/quickstart:latest /usr/bin/docker-quickstart
# docker run --name cdh --hostname=quickstart.cloudera --privileged=true -t -i -p 8020:8020 -p 8022:8022 -p 7180:7180 -p 21050:21050 -p 50070:50070 -p 50075:50075 -p 50010:50010 -p 50020:50020 -p 8890:8890 -p 60010:60010 -p 10002:10002 -p 25010:25010 -p 25020:25020 -p 18088:18088 -p 8088:8088 -p 19888:19888 -p 7187:7187 -p 11000:11000 cloudera/quickstart /bin/bash -c '/usr/bin/docker-quickstart && /home/cloudera/cloudera-manager --express && service ntpd start'

docker run --name cdh --hostname=quickstart.cloudera --privileged=true -d --restart=always -p 8020:8020 -p 8022:8022 -p 7180:7180 -p 21050:21050 -p 50070:50070 -p 50075:50075 -p 50010:50010 -p 50020:50020 -p 8890:8890 -p 60010:60010 -p 10002:10002 -p 25010:25010 -p 25020:25020 -p 18088:18088 -p 8088:8088 -p 19888:19888 -p 7187:7187 -p 11000:11000 cloudera/quickstart /bin/bash -c '/usr/bin/docker-quickstart && /home/cloudera/cloudera-manager --express --force && service ntpd start'

docker run --name cdh --hostname=quickstart.cloudera --privileged=true -d --restart=always -p 8020:8020 -p 8022:8022 -p 7180:7180 -p 21050:21050 -p 50070:50070 -p 50075:50075 -p 50010:50010 -p 50020:50020 -p 8890:8890 -p 60010:60010 -p 10002:10002 -p 25010:25010 -p 25020:25020 -p 18088:18088 -p 8088:8088 -p 19888:19888 -p 7187:7187 -p 11000:11000 cloudera/quickstart /bin/bash -c '/usr/bin/docker-quickstart'

docker run --name cdh --hostname=quickstart.cloudera --privileged=true -d -p 8020:8020 -p 8022:8022 -p 7180:7180 -p 21050:21050 -p 50070:50070 -p 50075:50075 -p 50010:50010 -p 50020:50020 -p 8890:8890 -p 60010:60010 -p 10002:10002 -p 25010:25010 -p 25020:25020 -p 18088:18088 -p 8088:8088 -p 19888:19888 -p 7187:7187 -p 11000:11000 cloudera/quickstart /bin/bash -c '/usr/bin/docker-quickstart'

docker logs -f cdh
成
```

2. 找hdfs配置文件：

```shell
# which hadoop
/usr/bin/hadoop
# find /|grep core-site.xml 
/usr/lib/hadoop-0.20-mapreduce/example-confs/conf.secure/core-site.xml
/usr/lib/hadoop-0.20-mapreduce/example-confs/conf.pseudo/core-site.xml
/etc/impala/conf.dist/core-site.xml
/etc/hadoop-kms/conf.dist/core-site.xml
/etc/oozie/conf.dist/hadoop-conf/core-site.xml
/etc/hadoop/conf.impala/core-site.xml
/etc/hadoop/conf.pseudo/core-site.xml
/etc/hadoop/conf.empty/core-site.xml
# cd ./lib/hadoop-0.20-mapreduce/example-confs/
# grep -nr -A 1 fs.default.name *
conf.pseudo/core-site.xml:6:    <name>fs.default.name</name>
conf.pseudo/core-site.xml-7-    <value>hdfs://localhost:8020</value>
--
conf.secure/core-site.xml:31:    <name>fs.default.name</name>
conf.secure/core-site.xml-32-    <value>hdfs://${local.namenode}:8020</value>
# hadoop fs -ls hdfs://localhost:8020/ 
Found 5 items
drwxrwxrwx   - hdfs  supergroup          0 2016-04-06 02:26 hdfs://localhost:8020/benchmarks
drwxr-xr-x   - hbase supergroup          0 2019-11-13 08:08 hdfs://localhost:8020/hbase
drwxrwxrwt   - hdfs  supergroup          0 2019-11-13 08:08 hdfs://localhost:8020/tmp
drwxr-xr-x   - hdfs  supergroup          0 2016-04-06 02:27 hdfs://localhost:8020/user
drwxr-xr-x   - hdfs  supergroup          0 2016-04-06 02:27 hdfs://localhost:8020/var
# 在87上
$ hadoop fs -ls hdfs://172.16.30.17:8020/
Found 5 items
drwxrwxrwx   - hdfs  supergroup          0 2016-04-06 10:26 hdfs://172.16.30.17:8020/benchmarks
drwxr-xr-x   - hbase supergroup          0 2019-11-13 16:08 hdfs://172.16.30.17:8020/hbase
drwxrwxrwt   - hdfs  supergroup          0 2019-11-13 16:08 hdfs://172.16.30.17:8020/tmp
drwxr-xr-x   - hdfs  supergroup          0 2016-04-06 10:27 hdfs://172.16.30.17:8020/user
drwxr-xr-x   - hdfs  supergroup          0 2016-04-06 10:27 hdfs://172.16.30.17:8020/var
```

3. 访问管理页面[http://172.16.30.17:7180](http://172.16.30.17:7180/)



# 搭建hadoop

1. 参考：<https://hub.docker.com/r/sequenceiq/hadoop-docker>

```shell
docker pull sequenceiq/hadoop-docker:2.7.0
# 172
docker run -d -p 9000:9000 --name=hdfs2.7.0 --restart=always sequenceiq/hadoop-docker:2.7.0 /etc/bootstrap.sh
hadoop fs -ls hdfs://172.16.30.17:9000/
# 87
docker run -d -p 9001:9000 --name=hdfs2.7.0 --restart=always sequenceiq/hadoop-docker:2.7.0 /etc/bootstrap.sh
hadoop fs -ls hdfs://10.4.1.87:9001/
```



# 参考资料

1. <https://docs.cloudera.com/documentation/enterprise/5-6-x/topics/quickstart_docker_container.html>
2. <https://hub.docker.com/r/sequenceiq/hadoop-docker>