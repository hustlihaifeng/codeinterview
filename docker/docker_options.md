# 1. docker run

1.1 `--privileged=true`:<https://blog.csdn.net/halcyonbaby/article/details/43499409> 用privileged拉起的容器里面的root用户有真正的root权限。



# 2. docker push 非默认仓库

```shell
# 登录
$ docker login registry.xxxx.com
Username: xxxx
Password: 
Login Succeeded
$ docker push registry.xxxx.com/library/yiplatform-package:1.1.0
```

