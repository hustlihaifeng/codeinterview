# oracle环境搭建

1. https://github.com/oracle/docker-images/blob/master/OracleDatabase/SingleInstance/README.md> 有从docker镜像安装的文档，需要先下载对应的zip文件：如`linuxx64_12201_database.zip`。 

   1. 建议从docker镜像安装，使用zip文件直接在centos服务器安装的话，需要处理256色、x-server等问题，反正我是没装成功过。
   2. 拉起oracle容器后，sys密码会在`docker logs`的第一行。
   3. 登录oralce需要安装oracle client，主要是下面这三个rpm包：

   ```shell
   oracle-instantclient11.2-sqlplus-11.2.0.4.0-1.x86_64.rpm
   oracle-instantclient11.2-devel-11.2.0.4.0-1.x86_64.rpm
   oracle-instantclient11.2-basic-11.2.0.4.0-1.x86_64.rpm
   ```

   4. oracle client装好后，使用sys登录:`rlwrap sqlplus sys/init_password@host:port/service_name as sysdba`,然后创建常见用户：

   ```sql
   CREATE USER test_user IDENTIFIED BY test_pass;
   GRANT CONNECT, RESOURCE, DBA TO test_user;
   ```

   然后就可以使用：`rlwrap sqlplus test_user/test_pass@host:port/service_name`来登录oracle了。

2. client 安装过程与1中database类似，关键命令是：`docker build -t oracle/instantclient:19 .`，其他详见上面的文档。

# oracle数据库连接方式

1. oracle client安装好后，有sqlplus这个工具来链接，类似于mysql-client。链接格式是：`sqlplus user/password@host:port/service_name`

   1. 服务名一般会在`$ORACLE_HOME/tnsnames.ora`里面配置，如果`$ORACLE_HOME/tnsnames.ora`不存在，在`$ORACLE_HOME`下面find一下`tnsnames.ora`即可。
   2. 在线查询服务名,下面三种都可以：

   ```sql
   select * from global_name;
   select sys_context('userenv','db_name') from dual;
   select ora_database_name from dual;
   ```

   3. 服务名没配好，或者错误的话，会报错：`ORA-12162: TNS:net service name is incorrectly specified`

2. sqlplus 默认不能退格，不能搜索历史记录，不能使用方向键，使用起来很不方便。安装一个`rlwrap`，可以解决这个问题。这样链接串变成：`rlwrap sqlplus user/password@host:port/service_name`.

3. sqlplus 默认显示的一行长度很短，select输出被被折行显示，看起来很乱。`set linesize 1000`可以将行长度调大。

4. 3中的设置并不能根本解决问题，只能缓解。加上sqlplus还有中文字符显示问题，当行最大长度限制（2499）。如果桌面环境能够访问oracle数据库的话。那么推荐使用plsql这个客户端来在桌面访问oralce数据库，没有单行长度限制，没有中文乱码。

# oracle依赖库

1. 使用sqlplus形式链接oracle需要装三个依赖包：

```shell
oracle-instantclient11.2-sqlplus-11.2.0.4.0-1.x86_64.rpm
oracle-instantclient11.2-devel-11.2.0.4.0-1.x86_64.rpm
oracle-instantclient11.2-basic-11.2.0.4.0-1.x86_64.rpm
```

2. 推荐将这三个依赖包打进基础软件包`rpms.tar.gz`,然后在镜像构建过程中安装这三个依赖：

- `install_rpms.sh`

```shell
#!/bin/bash

echo "安装rpm包"
tar -xf rpms.tar.gz
cd rpms
yum localinstall -y *.rpm
cd ..
rm -rf rpms
yum clean all
echo "rpm包安装完成"
```

- 在Dockerfile里面加上

```shell
ADD ./local_dir /target_dir/
RUN cd /target_dir/ && \
    /bin/bash install_rpms.sh
```

