# 查找hdfs配置文件

```shell
$ which hadoop
~/bigdata/hadoop-2.7.2/bin/hadoop
$ cd bigdata/hadoop-2.7.2/
$ bin/hadoop jar share/hadoop/mapreduce/hadoop-mapreduce-examples-2.7.2.jar  wordcount ./wc.input output2  # 用mapreduce跑wordcount,没开启yarn的话，会导致连接不上ResourceManager。详见https://stackoverflow.com/questions/20586920/hadoop-connecting-to-resourcemanager-failed
$  find .|grep hdfs-site.xml
./etc/hadoop/hdfs-site.xml
./share/hadoop/hdfs/templates/hdfs-site.xml
$ ls ./share/hadoop/hdfs/templates/
hdfs-site.xml
$ ls ./etc/hadoop/
capacity-scheduler.xml  hadoop-metrics2.properties  httpfs-signature.secret  log4j.properties            slaves
configuration.xsl       hadoop-metrics.properties   httpfs-site.xml          mapred-env.cmd              ssl-client.xml.example
container-executor.cfg  hadoop-policy.xml           kms-acls.xml             mapred-env.sh               ssl-server.xml.example
core-site.xml           hdfs-site.xml               kms-env.sh               mapred-queues.xml.template  yarn-env.cmd
hadoop-env.cmd          httpfs-env.sh               kms-log4j.properties     mapred-site.xml             yarn-env.sh
hadoop-env.sh           httpfs-log4j.properties     kms-site.xml             mapred-site.xml.template    yarn-site.xml # 可以看到，./etc/hadoop/hdfs-site.xml，才是真正的hdfs配置文件

# 在./etc/hadoop/hdfs-site.xml 里面可以找到dfs.namenode.name.dir、dfs.namenode.data.dir、dfs.http.address这些关键配置项。
```

# 关键配置项

- hdfs-site.xml

```shell
dfs.namenode.name.dir=file:/data1/unown/bigdata/hadoop-2.7.2/tmp/dfs/name
dfs.namenode.data.dir=file:/data1/unown/bigdata/hadoop-2.7.2/tmp/dfs/data
dfs.http.address=10.7.0.5:50070
```

- core-site.xml

```shell
fs.defaultFS=hdfs://10.7.0.5:9001
```

# WebHDFS

```shell
curl -i "http://10.7.0.5:50070/webhdfs/v1/tmp/?op=LISTSTATUS"
curl -i -L "http://10.7.0.5:50070/webhdfs/v1//tmp/test/wc.input?op=OPEN
```

## 拉取

1. 详见<https://stackoverflow.com/questions/16865162/is-there-any-way-to-download-a-hdfs-file-using-webhdfs-rest-api>

- `curl -L "http://10.7.0.5:50070/webhdfs/v1/tmp/test/wc.input?op=OPEN" -o a`会将文件下载到a里面
	- 也可以` curl -L "http://10.7.0.5:50070/webhdfs/v1/tmp/test/wc.input?op=OPEN" -O`来将文件输出到与远端同名的文件中，这里却是`wc.input?op=OPEN`，感觉也不好。
	- curl -i -L "http://localhost:50075/webhdfs/v1/demofile.txt?op=OPEN" -o ~/demofile.txt`（多了一个`-i`）会先输出头信息，然后后面是文件内容。头信息如：
```shell
HTTP/1.1 307 TEMPORARY_REDIRECT
Cache-Control: no-cache
Expires: Sun, 07 Apr 2019 02:39:14 GMT
Date: Sun, 07 Apr 2019 02:39:14 GMT
Pragma: no-cache
Expires: Sun, 07 Apr 2019 02:39:14 GMT
Date: Sun, 07 Apr 2019 02:39:14 GMT
Pragma: no-cache
Content-Type: application/octet-stream
Location: http://cpm10705:50075/webhdfs/v1/tmp/test/wc.input?op=OPEN&namenoderpcaddress=10.7.0.5:9001&offset=0
Content-Length: 0
Server: Jetty(6.1.26)

HTTP/1.1 200 OK
Access-Control-Allow-Methods: GET
Access-Control-Allow-Origin: *
Content-Type: application/octet-stream
Connection: close
Content-Length: 71
```

- 直接`curl -i "http://10.7.0.5:50070/webhdfs/v1/tmp/test/wc.input?op=OPEN"`会得到实际datanode的url，上面那个stackoverflow的连接说然后可以直接wget得到的datanode连接来得到文件，但是我试了不行。

```shell
HTTP/1.1 307 TEMPORARY_REDIRECT
Cache-Control: no-cache
Expires: Sun, 07 Apr 2019 02:41:58 GMT
Date: Sun, 07 Apr 2019 02:41:58 GMT
Pragma: no-cache
Expires: Sun, 07 Apr 2019 02:41:58 GMT
Date: Sun, 07 Apr 2019 02:41:58 GMT
Pragma: no-cache
Content-Type: application/octet-stream
Location: http://cpm10705:50075/webhdfs/v1/tmp/test/wc.input?op=OPEN&namenoderpcaddress=10.7.0.5:9001&offset=0
Content-Length: 0
Server: Jetty(6.1.26)
```

