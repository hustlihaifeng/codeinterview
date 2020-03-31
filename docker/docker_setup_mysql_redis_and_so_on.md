# 1. 更新私有仓库

   - `vim /etc/docker/daemon.json`

   ```shell
   {
           "registry-mirrors": ["https://dq5xnrwz.mirror.aliyuncs.com","https://registry.docker-cn.com","http://hub-mirror.c.163.com","https://mirror.ccs.tencentyum.com"]
   }
   ```

   - 参考
      - <https://www.jianshu.com/p/1a4025c5f186>
      - <https://blog.csdn.net/qq_27575627/article/details/73470120>
      - <https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors>

# 2. 拉起redis服务

   ```shell
   docker pull redis:3.2
   mkdir -p /data/redis/data
   docker run -p 6379:6379 -v /data/redis/data:/data  -d --name=redis --restart=always redis:3.2 redis-server --appendonly yes
   docker exec -it redis redis-cli
   
   docker start redis
   
   systemctl daemon-reload
   systemctl restart docker
   
   docker update --restart=always redis
   ```

# 3. 拉取并搭建mysql的docker服务

- <https://blog.csdn.net/woniu211111/article/details/80968154>

```shell
mkdir -p /data/mysql/conf
mkdir -p /data/mysql/data/3306
vim /data/mysql/conf/my.cnf.3306
cat /data/mysql/conf/my.cnf.3306
[mysqld]
user=mysql
character-set-server=utf8
default_authentication_plugin=mysql_native_password
[client]
default-character-set=utf8
[mysql]
default-character-set=utf8

docker run -p 3306:3306 --privileged=true -v /data/mysql/conf/my.cnf.3306:/etc/mysql/my.cnf -v /data/mysql/data/3306:/var/lib/mysql --name=mysql  --restart=always -e MYSQL_ROOT_PASSWORD=root  -d mysql:5.7.20
docker exec -it mysql mysql -uroot -proot
```

# 4. 搭建registry

```shell
docker pull registry
mkdir -p /data/docker_registry
docker run -p 5000:5000 -v /home/lhf/docker_registry:/var/lib/registry --name=registry  --restart=always -d registry
```

# 5. python：3.4

```shell
docker run --name=python3 --restart=always -d python:3.4 sh
```

