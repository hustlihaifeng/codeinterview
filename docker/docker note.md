# 环境搭建
1. 从[http://mirrors.aliyun.com/docker-toolbox/windows/docker-toolbox/](http://mirrors.aliyun.com/docker-toolbox/windows/docker-toolbox/)下载 docker toolbox来在windows上安装。详见[http://www.runoob.com/docker/windows-docker-install.html](http://www.runoob.com/docker/windows-docker-install.html)
2. 安装完后，quick start，执行`docker run hello-world`来跑hello world。
3. 镜像加速：打开或新建：`/etc/docker/daemon.json` linux 或`%programdata%\docker\config\daemon.json` Windows来配置Daemon：
```json
{
  "registry-mirrors": ["http://hub-mirror.c.163.com"]
}
```
# docker使用
1. `docker run ubuntu:15.10 /bin/echo "Hello world"` Docker 以 ubuntu15.10 镜像创建一个新容器，然后在容器里执行 `/bin/echo "Hello world"`，然后输出结果
- `docker run`: 运行一个容器
> run         Run a command in a new container

- `ubuntu:15.10`：指定要运行的镜像，Docker首先从本地主机上查找镜像是否存在，如果不存在，Docker 就会从镜像仓库 Docker Hub 下载公共镜像。ubuntu是指仓库源，15.10是版本，如果没有指定版本，则是lastest版本。
- 输出结果：可以通过` -i -t`，让docker可以交互。
	- `-i`:允许你对容器内的标准输入 (STDIN) 进行交互
	> -i, --interactive                    Keep STDIN open even if not attached
	- `-t`:在新容器内指定一个伪终端或终端
	> -t:在新容器内指定一个伪终端或终端
- 后台运行：`docker run -d`：
```shell
$ docker run -d ubuntu:15.10 /bin/sh -c "while true; do echo hello world; sleep 1; done"
2b1b7a428627c51ab8810d541d759f072b4fc75487eed05812646b8534a2fe63
```
> $ docker run -d ubuntu:15.10 /bin/sh -c "while true; do echo hello world; sleep 1; done"
2b1b7a428627c51ab8810d541d759f072b4fc75487eed05812646b8534a2fe63

- 查看哪些容器在运行:`docker ps`
- 输出某个容器的日志：`docker logs 0cdbd3a4d94d`、`docker logs friendly_wright`
	- `docker logs -f 0cdbd3a4d94d`循环查看某个容器的输出
	> -f, --follow         Follow log output
- 停止容器：`docker stop`后面接容器id或者名字。
- 重启容器：`docker restart`
- 启动停止了的容器：`docker start`
- 查看容器内运行的进程：`docker top $id`
- 删除不需要的容器：`docker rm`必须是停止装填，不然会报错。

2. docker中起服务并映射到主机端口：`-P` `-p port:port`
- `docker run -d -P training/webapp python app.py` 映射的主机的任意端口，Flask默认端口是5000
>   -P, --publish-all                    Publish all exposed ports to random ports

- `docker run -d -p 5000:5000 training/webapp python app.py` 将docker的5000端口映射到主机的5000端口
> -p, --publish list                   Publish a container's port(s) to the host

- `docker port` 后面接id或者容器名可以查看某个特定容器的端口映射。

# docker镜像管理
1. 查看本地主机镜像列表：`docker images`
2. 主动下载某个镜像：`docker pull ubuntu:13.10`。 当本地没有某个镜像时，docker会自动下载这个镜像。
3. 查找镜像：`docker search httpd`， 也可以在`https://hub.docker.com/`上搜索镜像。
4. 创建新镜像：
- 对已有镜像进行修改后提交：`docker commit -m="has update" -a="runoob" e218edb10161 runoob/ubuntu:v2` 
  - -m:提交的描述信息
  - -a:指定镜像作者
  - e218edb10161：容器ID
  - runoob/ubuntu:v2:指定要创建的目标镜像名
  > docker commit -m="test" -a="haifeng" 5ae84f67f428 training/webapp:v2
  > sha256:d438c7edc266200014a264d477a8211df6221d3331fe28293f93decb9e49be9a

- 写Dockerfile 然后运行`docker build`从0开始构建镜像。如：
```shell
FROM    centos:6.7
MAINTAINER      Fisher "fisher@sudops.com"

RUN     /bin/echo 'root:123456' |chpasswd
RUN     useradd runoob
RUN     /bin/echo 'runoob:123456' |chpasswd
RUN     /bin/echo -e "LANG=\"en_US.UTF-8\"" >/etc/default/local
EXPOSE  22
EXPOSE  80
CMD     /usr/sbin/sshd -D
```
每一个指令都会在镜像上创建一个新的层，每一个指令的前缀都必须是大写的。第一条FROM，指定使用哪个镜像源。RUN 指令告诉docker 在镜像内执行命令，安装了什么。。。
	- `docker build -t runoob/centos:6.7 .`
	- -t ：指定要创建的目标镜像名
	- . ：Dockerfile 文件所在目录，可以指定Dockerfile 的绝对路径

5. 为镜像打标签：`docker tag 860c279d2fec runoob/centos:dev`为image id为860c279d2fec的用户创建标签`runoob/centos:dev`

# docker端口映射

如下的语句：

```shell
docker run -d -P training/webapp python app.py # 容器内部端口随机映射到主机的高端口
docker run -d -p 5000:5000 training/webapp python app.py # 容器内部端口绑定到指定的主机端口
docker run -d -p 127.0.0.1:5001:5000 training/webapp python app.py # 指定容器绑定的网络地址，比如绑定 127.0.0.1
docker run -d -p 127.0.0.1:5000:5000/udp training/webapp python app.py # 默认都是绑定 tcp 端口，如果要绑定 UDP 端口，可以在端口后面加上 /udp
docker run -d -P --name runoob training/webapp python app.py # 使用 --name 标识来命名容器
docker ps
```

# docker搭建私有仓库

```shell
sudo docker pull docker:17.09.0-ce
sudo docker run -i -t centos:6.7

alias docker='sudo /usr/bin/docker'
sudo docker pull registry
sudo docker run -d -p 5000:5000 registry
sudo docker exec -it 75bf70ff5e97 sh
find / -name registry # 找到docker/registry的包含目录：/etc/docker/registry
docker run -d -p 5000:5000 -v /home/lhf/docker_registry:/var/lib/registry registry
docker exec -it 6737b57050d1 sh # 登录docker容器
ifconfig # 查看拉起的docker容器的ip，得到172.17.0.2

# 测试
docker pull busybox
docker tag busybox 172.17.0.2:5000/busybox
docker push 172.17.0.2:5000/busybox
# 失败报错：Get https://172.17.0.2:5000/v2/: http: server gave HTTP response to HTTPS client，因为与docker registry交互默认使用的是https，然而此处搭建的私有仓库只提供http服务，所以当与私有仓库交互时就会报上面的错误。
sudo vim  /usr/lib/systemd/system/docker.service # 在ExecStart=/usr/bin/dockerd 后加 --insecure-registry 172.17.0.2:5000，来在启动docker server时增加启动参数为默认使用http访问。并将StartLimitInterval后面的值改为一个小值，比如1s，否则重启docker服务的时候，会提示重启太过频繁而失败，然后我们一直重试，一直失败！！！！还有一个报错，说什么配置重复，将重复的配置删除即可，我的是/usr/bin/dockerd跟后面的配置重复。
# 出错时，将原来的相关进程干掉，然后手动执行命令，看报错。或者调整参数（不要参数），看报错。来逐步定位问题。
sudo systemctl daemon-reload
sudo systemctl restart docker
docker push 172.17.0.2:5000/busybox
docker tag  hello-world 172.17.0.2:5000/hello-world
docker push 172.17.0.2:5000/hello-world
sudo docker pull 172.17.0.2:5000/busybox

curl -XGET http://172.17.0.2:5000/v2/_catalog # 可以得到 {"repositories":["busybox","hello-world"]},得到远端的景象列表
curl -XGET http://172.17.0.2:5000/v2/busybox/tags/list

```

# docker删除镜像

```shell
docker ps -a # 查看所有container
sudo docker stop 40cc9f5f2aa4 # 停止正在运行的container
sudo docker rm 40cc9f5f2aa4 # 删除该image的container
sudo docker images # 找到该image的imageid
sudo docker rmi 9f1de3c6ad53 # 删除该image
```

# docker进入容器的集中方式

- [Docker容器进入的4种方式](https://www.cnblogs.com/xhyan/p/6593075.html)
- `sudo docker exec -it 75bf70ff5e97 sh`

# docker退出却不关闭容器

- [Ctrl+P+Q](https://www.jianshu.com/p/b1ce248d2a42)

# docker重命名

```shell
sudo docker rename quizzical_heisenberg docker_repo
```

# 删除关闭了的docker容器

```shell
docker ps -a |grep Exited|awk '{print $1}'|xargs sudo /usr/bin/docker rm
```



# 参考资料

- [centos7 Docker私有仓库搭建及删除镜像](https://www.cnblogs.com/Tempted/p/7768694.html)

