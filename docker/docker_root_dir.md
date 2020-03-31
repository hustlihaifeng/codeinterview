# docker数据目录设置与迁移

```shell
docker info|grep "Docker Root Dir"  # /var/lib/docker
mkdir -p /data/docker
systemclt stop docker
mv /var/lib/docker/* /data/docker/
locate docker.service
vim /usr/lib/systemd/system/docker.service
在ExecStart=/usr/bin/dockerd后面加 --graph=/data/docker
systemctl daemon-reload
systemclt start docker
```

# 参考资料

- [修改Docker默认镜像和容器的存储位置](https://www.cnblogs.com/bigberg/p/8057807.html) 里面的`storage-driver`如果改成和最初的`docker info`里面的结果不一样的话，数据迁移失败，新拉起的docker里面没有任何镜像和container。