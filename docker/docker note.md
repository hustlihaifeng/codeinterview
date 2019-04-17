# 环境搭建
1. 从[http://mirrors.aliyun.com/docker-toolbox/windows/docker-toolbox/](http://mirrors.aliyun.com/docker-toolbox/windows/docker-toolbox/)下载 docker toolbox来在windows上安装。详见[http://www.runoob.com/docker/windows-docker-install.html](http://www.runoob.com/docker/windows-docker-install.html)
2. 安装完后，quick start，执行`docker run haifeng_test`来跑hello world。
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
ifconfig # 查看拉起的docker容器的ip，得到127.0.0.1

# 测试
docker pull busybox
docker tag busybox 127.0.0.1:5000/busybox
docker push 127.0.0.1:5000/busybox
# 失败报错：Get https://127.0.0.1:5000/v2/: http: server gave HTTP response to HTTPS client，因为与docker registry交互默认使用的是https，然而此处搭建的私有仓库只提供http服务，所以当与私有仓库交互时就会报上面的错误。配置成127.0.0.1不会报错
sudo vim  /usr/lib/systemd/system/docker.service # 在ExecStart=/usr/bin/dockerd 后加 --insecure-registry 127.0.0.1:5000，来在启动docker server时增加启动参数为默认使用http访问。并将StartLimitInterval后面的值改为一个小值，比如1s，否则重启docker服务的时候，会提示重启太过频繁而失败，然后我们一直重试，一直失败！！！！还有一个报错，说什么配置重复，将重复的配置删除即可，我的是/usr/bin/dockerd跟后面的配置重复。
# 出错时，将原来的相关进程干掉，然后手动执行命令，看报错。或者调整参数（不要参数），看报错。来逐步定位问题。
sudo systemctl daemon-reload
sudo systemctl restart docker
docker push 127.0.0.1:5000/busybox
docker tag  haifeng_test 127.0.0.1:5000/haifeng_test
docker push 127.0.0.1:5000/haifeng_test
sudo docker pull 127.0.0.1:5000/busybox

curl -XGET http://127.0.0.1:5000/v2/_catalog # 可以得到 {"repositories":["busybox","haifeng_test"]},得到远端的景象列表
curl -XGET http://127.0.0.1:5000/v2/busybox/tags/list

```

# docker删除镜像

```shell
docker ps -a # 查看所有container
sudo docker stop 40cc9f5f2aa4 # 停止正在运行的container
sudo docker rm 40cc9f5f2aa4 # 删除该image的container
sudo docker images # 找到该image的imageid
sudo docker rmi 9f1de3c6ad53 # 删除该image
```

# docker进入容器的几种方式

- [Docker容器进入的4种方式](https://www.cnblogs.com/xhyan/p/6593075.html)
- `sudo docker exec -it 75bf70ff5e97 sh`

# docker退出却不关闭容器

- [Ctrl+P+Q](https://www.jianshu.com/p/b1ce248d2a42)

# docker重命名

```shell
sudo docker rename quizzical_heisenberg docker_repo
```

# docker打包镜像文件

- 详见[https://github.com/Juniper/contrail-docker/wiki/Save-docker-images-to-a-tar.gz-file](https://github.com/Juniper/contrail-docker/wiki/Save-docker-images-to-a-tar.gz-file)

```shell
docker save 127.0.0.1:5000/contrail-analyticsdb-u14.04:4.0.0.0-3016 | gzip -c > contrail-analyticsdb-u14.04-4.0.0.0-3016.tar.gz

docker images |grep ubuntu
docker save 0b1edfbffd27 | gzip -c > ubuntu-16.04.tar.gz
hdfs dfs -put ./ubuntu-16.04.tar.gz /tmp/test/ubuntu-16.04.tar.gz

docker save haifeng_test:test |gzip -c > haifeng_test-test.tar.gz
```

# docker从打包文件导入镜像

```shell
$ docker load -i /data1/download/image-7-1554608667
Loaded image ID: sha256:0b1edfbffd27c935a666e233a0042ed634205f6f754dbe20769a60369c614f85
$ docker load -i /data1/download/image-7-1554608667|grep -Eo 'sha256:\w*'
sha256:0b1edfbffd27c935a666e233a0042ed634205f6f754dbe20769a60369c614f85
$ docker images --no-trunc|grep sha256:0b1edfbffd27c935a666e233a0042ed634205f6f754dbe20769a60369c614f85
ubuntu                                                                            16.04               sha256:0b1edfbffd27c935a666e233a0042ed634205f6f754dbe20769a60369c614f85   11 months ago       113MB
```



# 删除关闭了的docker容器

```shell
docker ps -a |grep Exited|awk '{print $1}'|xargs sudo /usr/bin/docker rm
```

# docker远端仓库可用api

- [https://www.jianshu.com/p/6a7b80122602](https://www.jianshu.com/p/6a7b80122602)
- **[https://docs.docker.com/registry/spec/api/#detail](https://docs.docker.com/registry/spec/api/#detail)**
- [Docker registry V2 推送镜像、拉取镜像、搜索镜像、删除镜像和垃圾回收](https://blog.csdn.net/nklinsirui/article/details/80705306#%E5%AE%98%E6%96%B9%E6%96%87%E6%A1%A3)

## 查询

```shell
curl -XGET http://127.0.0.1:5000/v2/ # 如果实现了v2 api，会返回一个{}，相当于ping
curl -XGET http://127.0.0.1:5000/v2/_catalog # 获取镜像列表
curl -XGET http://127.0.0.1:5000/v2/haifeng_test/tags/list # 获取某个镜像有多少个tag：{"name":"haifeng_test","tags":["test","latest"]}
curl -XGET http://127.0.0.1:5000/v2/haifeng_test/manifests/test # 获取某个镜像版本的详细信息，总共有：get put delete head四种请求

curl -XGET http://127.0.0.1:5000/v2/_catalog?n=2 # 获取前两个镜像，每次获取内容相同，不能分页
curl -XGET http://127.0.0.1:5000/v2/haifeng_test/tags/list?n=2 # 获取某个镜像的前两个版本，每次获取的内容相同，好像不能实现分页
curl -XGET "http://127.0.0.1:5000/v2/haifeng_test/tags/list?n=1&last=1" # 分页获取，文档说这样可以，但是测试时好像不行
```

manifests的输出类似于：

```json
{
	"schemaVersion": 1,
	"name": "haifeng_test",
	"tag": "test",
	"architecture": "amd64",
	"fsLayers": [{
			"blobSum": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
		},
		{
			"blobSum": "sha256:1b930d010525941c1d56ec53b97bd057a67ae1865eebf042686d2a2d18271ced"
		}
	],
	"history": [{
			"v1Compatibility": "{\"architecture\":\"amd64\",\"config\":{\"Hostname\":\"\",\"Domainname\":\"\",\"User\":\"\",\"AttachStdin\":false,\"AttachStdout\":false,\"AttachStderr\":false,\"Tty\":false,\"OpenStdin\":false,\"StdinOnce\":false,\"Env\":[\"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin\"],\"Cmd\":[\"/hello\"],\"ArgsEscaped\":true,\"Image\":\"sha256:a6d1aaad8ca65655449a26146699fe9d61240071f6992975be7e720f1cd42440\",\"Volumes\":null,\"WorkingDir\":\"\",\"Entrypoint\":null,\"OnBuild\":null,\"Labels\":null},\"container\":\"8e2caa5a514bb6d8b4f2a2553e9067498d261a0fd83a96aeaaf303943dff6ff9\",\"container_config\":{\"Hostname\":\"8e2caa5a514b\",\"Domainname\":\"\",\"User\":\"\",\"AttachStdin\":false,\"AttachStdout\":false,\"AttachStderr\":false,\"Tty\":false,\"OpenStdin\":false,\"StdinOnce\":false,\"Env\":[\"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin\"],\"Cmd\":[\"/bin/sh\",\"-c\",\"#(nop) \",\"CMD [\\\"/hello\\\"]\"],\"ArgsEscaped\":true,\"Image\":\"sha256:a6d1aaad8ca65655449a26146699fe9d61240071f6992975be7e720f1cd42440\",\"Volumes\":null,\"WorkingDir\":\"\",\"Entrypoint\":null,\"OnBuild\":null,\"Labels\":{}},\"created\":\"2019-01-01T01:29:27.650294696Z\",\"docker_version\":\"18.06.1-ce\",\"id\":\"9f5834b25059239faef06a9ba681db7b7c572fc0d87d2b140b10e90e50902b53\",\"os\":\"linux\",\"parent\":\"65b27d3bd74d2cf4ea3aa9e250be6c632f0a347e8abd5485345c55fa6eed0258\",\"throwaway\":true}"
		},
		{
			"v1Compatibility": "{\"id\":\"65b27d3bd74d2cf4ea3aa9e250be6c632f0a347e8abd5485345c55fa6eed0258\",\"created\":\"2019-01-01T01:29:27.416803627Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) COPY file:f77490f70ce51da25bd21bfc30cb5e1a24b2b65eb37d4af0c327ddc24f0986a6 in / \"]}}"
		}
	],
	"signatures": [{
		"header": {
			"jwk": {
				"crv": "P-256",
				"kid": "USX5:66ZV:IHQM:IILY:LA7D:BKKO:TDDZ:QWP6:6QRC:P6QO:37ZL:FZRJ",
				"kty": "EC",
				"x": "guUQpSApBOzRLJMl7XGRpRKswkDn9AsUd0t1fylckAs",
				"y": "qxZl09dYVa2Fc0nsyUnRNv5b5EJUgRvl78OiOp0pTnE"
			},
			"alg": "ES256"
		},
		"signature": "TQgUh1P2pyf1MZXTm43XY4whdW9lgwFAZS9QWnVdO1307Y8QJVEepc1DNpe2tpYKITWEUIOC-zuOgQ9j6Y6rPg",
		"protected": "eyJmb3JtYXRMZW5ndGgiOjIxMzUsImZvcm1hdFRhaWwiOiJDbjAiLCJ0aW1lIjoiMjAxOS0wMy0yNlQwMzowNjoxMVoifQ"
	}]
}
```



## 增加

- [<https://docs.docker.com/registry/spec/api/#pushing-an-image>](https://docs.docker.com/registry/spec/api/#pushing-an-image)

```shell
POST /v2/<name>/blobs/uploads/
HEAD /v2/<name>/blobs/<digest>
docker push 127.0.0.1:5000/haifeng_test:tag1
```



## 删除

- [<https://docs.docker.com/registry/spec/api/#deleting-an-image>](https://docs.docker.com/registry/spec/api/#deleting-an-image)

```shell
DELETE /v2/<name>/manifests/<reference>
```

- `-e REGISTRY_STORAGE_DELETE_ENABLED="true"`
- `docker exec -it registry  /bin/registry garbage-collect  /etc/docker/registry/config.yml`
- 完整：详见<https://blog.csdn.net/nklinsirui/article/details/80705306>

```shell
curl -XGET http://127.0.0.1:5000/v2/_catalog
curl -XGET http://127.0.0.1:5000/v2/word2vec/tags/list
docker ps|grep registry|grep /entrypoint.sh|awk '{print $1}'

curl --silent -H 'Accept: application/vnd.docker.distribution.manifest.v2+json' -X GET  http://127.0.0.1:5000/v2/word2vec/manifests/seq2bow
curl -v --silent -H 'Accept: application/vnd.docker.distribution.manifest.v2+json' -X GET  http://127.0.0.1:5000/v2/word2vec/manifests/seq2bow 2>&1 | grep Docker-Content-Digest | awk '{print ($3)}'  # 这个头一定要加，不然获取不到正确的Digest，后面删除会失败
curl -v --silent -H 'Accept: application/vnd.docker.distribution.manifest.v2+json' -X DELETE http://127.0.0.1:5000/v2/word2vec/manifests/sha256:ceebde8ef883624f19d69b395bdc719ae3a21814f929283eeede4e9c30d72175
docker exec -it registry  /bin/registry garbage-collect  /etc/docker/registry/config.yml
```

1. 在执行garbage-collect的时候，报错：TODO

```shell
2019/04/09 03:00:00 [I] [os_util.go:76] docker exec -it a4ead8733801 registry garbage-collect /etc/docker/registry/config.yml 
2019/04/09 03:00:00 [I] [os_util.go:79]  
2019/04/09 03:00:00 [W] [os_util.go:88] the input device is not a TTY
 
2019/04/09 03:00:00 [E] [docker_registry_operation.go:140] ErrorRegistryGarbageConnectAndRestart exit status 1
```



## 获取

- [<https://docs.docker.com/registry/spec/api/#pulling-an-image>](https://docs.docker.com/registry/spec/api/#pulling-an-image)

```shell
GET /v2/<name>/manifests/<reference>
HEAD /v2/<name>/manifests/<reference> # 检查是否存在
GET /v2/<name>/blobs/<digest>
docker pull 127.0.0.1:5000/busybox
```



# 私有镜像搭建

- `docker run -d -p 5000:5000 -v /opt/data/registry:/var/lib/registry  -v /data/config.yml:/etc/docker/registry/config.yml -e REGISTRY_STORAGE_DELETE_ENABLED="true"  registry`
- `docker run -d -p 5000:5000 -v /opt/data/registry:/var/lib/registry -e REGISTRY_STORAGE_DELETE_ENABLED="true"  registry`

# 坑

1. docker push推不上去，等待后保500 internal server error。查看日志发现`mkdir /var/lib/registry/docker: no such file or directory"`
   - /var/lib/registry下面什么都没有
   - 猜测是目录挂载在主机上，然后主机上的对应目录没删除了。重启registry后可以push了。
2. docker/docker/client 在PushImage失败后，尽然没有返回错误。这个库不能用
3. 默认的本地docker服务，只支持https去访问私有仓库。解决八法
   - `/usr/lib/systemd/system/docker.service`里面ExecStart=/usr/bin/dockerd 后加 --insecure-registry 本机ip:5000来开通本机的http访问。
   - 配置私有仓库使之支持https
   - 127.0.0.1默认可以http访问
4. 删除私有仓库镜像，执行内存回收后，再次推送该镜像到私有仓库，获取不到mainfest。只有重启私有仓库重新push后才可以。

   - 详见：[https://github.com/docker/distribution/issues/2270](https://github.com/docker/distribution/issues/2270) 上的bug issue，到2.7.0版本修复了，但是必须使用redis来作为缓存。
   - 临时解决方案是，删除镜像后不进行垃圾回收。然后提供一个垃圾回收的功能，与私有仓库重启相结合。
   - 见：[零时解决办法](https://github.com/docker/distribution/issues/2270#issuecomment-354433592), 也就是垃圾回收后必须重启，这样才不会出问题。

# 参考资料

- [centos7 Docker私有仓库搭建及删除镜像](https://www.cnblogs.com/Tempted/p/7768694.html)

