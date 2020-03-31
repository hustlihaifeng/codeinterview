# 1. 找镜像

```shell
locate cdh|grep -vE "docker"
cd /data/yiplatform/bigdata/cdh-5.15/
docker import ./cdh.tar.gz cdh:5.15
docker images|grep cdh
```

# 2. 起cdh容器

- https://www.cnblogs.com/yinzhengjie/articles/10413793.html

- https://blog.csdn.net/Post_Yuan/article/details/81235436?utm_source=blogxgwz2

```shell
docker run  -d -it -p 7180:7180 -p 7000:8020  --privileged  --hostname cdh_5  --name cdh_5 cdh:5.15 /usr/sbin/init
```

注意(主机名不能带 下划线)

# 3. 构建镜像
## 1. 安装步骤
1. 启动基础镜像

```
docker run -it -p 7180:7180 -p 7000:9000 --privileged   --hostname cdh  --name cdh_5  centos:7 /usr/sbin/init
```

2. 安装 httpd
```shell
yum -y install httpd
cp /etc/httpd/conf/httpd.conf /etc/httpd/conf/httpd.conf.`date +%F` 
grep ".tgz" /etc/httpd/conf/httpd.conf | grep -v '    #'
sed -i s'#.tgz#.tgz .parcel#' /etc/httpd/conf/httpd.conf
grep ".tgz" /etc/httpd/conf/httpd.conf | grep -v '    #'
systemctl start httpd 
```

3. 安装java
```shell
yum -y install java-1.8.0-openjdk.x86_64
```

4. 上传java连接mysql的jar包
```shell
docker cp mysql-connector-java-5.1.34-bin.jar xxx:/usr/share/java
mv mysql-connector-java-5.1.34-bin.jar mysql-connector-java.jar
```

5. 制作本地镜像
```shell
yum -y install yum-utils createrepo yum-plugin-priorities
createrepo .
```

6. 安装 cloudera
```shell
yum -y install cloudera-manager-daemons cloudera-manager-server
```

7. 初始化数据库配置
```shell
mysql -u root --password='root123' -e 'create database hive default character set utf8;' 
mysql -u root --password='root123' -e "CREATE USER 'hive'@'%' IDENTIFIED BY 'root123'" 
mysql -u root --password='root123' -e "GRANT ALL PRIVILEGES ON hive. * TO 'hive'@'%'" 
mysql -u root --password='root123' -e "create user 'amon'@'%' identified by 'root123'"
mysql -u root --password='root123' -e 'create database amon default character set utf8'
mysql -u root --password='root123' -e "grant all privileges on amon.* to 'amon'@'%'"
mysql -u root --password='root123' -e "create user 'rman'@'%' identified by 'root123'"
mysql -u root --password='root123' -e 'create database rman default character set utf8' 
mysql -u root --password='root123' -e "grant all privileges on rman.* to 'rman'@'%'"
mysql -u root --password='root123' -e "create user 'cm'@'%' identified by 'root123'"
mysql -u root --password='root123' -e 'create database cm default character set utf8' 
mysql -u root --password='root123' -e "grant all privileges on cm.* to 'cm'@'%'" 

/usr/share/cmf/schema/scm_prepare_database.sh -h 10.4.1.87 mysql cm cm root123
```

8. 安装 manager
```shell
yum -y install cloudera-manager-daemons cloudera-manager-server
```

9. 启动manager
```shell
systemctl start cloudera-scm-server 
```

10. 需要开通sshd服务
```shell
yum install openssh.x86_64
yum install openssh-server.x86_64
yum install openssh-clients.x86_64
yum -y install net-tools
```



# 实践

```shell
docker run  -d -it -p 7181:7180 -p 8021:8020  --privileged  --hostname cdh_5  --name cdh_haifeng cdh:5.15 /usr/sbin/init
```

