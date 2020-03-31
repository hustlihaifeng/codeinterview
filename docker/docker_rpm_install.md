1. 起一个运行容器的基础镜像的容器。
2. 以root用户运行:`docker exec -it -u root ec33c19230ca /bin/bash`
3. `yum install yum-plugin-downloadonly`
4. 解压`tar -zxf rpms.tar.gz;cd rpms`
5. `yum install --downloadonly --downloaddir=./ nfs-utils` 将rpm包下载到当前目录。
6. `cd ../;rm -f rpms.tar.gz; tar -zcf rpms.tar.gz rpms`重新打包
7. 退出容器，`docker commit 65ccc2d5dcc6 registry01.wezhuiyi.com/library/yiplatform-package:1.1.0`
8. `docker login registry01.wezhuiyi.com`
9. `docker push registry01.wezhuiyi.com/library/yiplatform-package:1.1.0`

# TODO

1. 直接在容器内下载好后，cp到另一个基础镜像起的容器中，感觉应该会小很多。观察到用`yum-plugin-downloadonly`下次一个52M的依赖后，镜像大小增加了1G多。
2. 依赖包先安装好，可能是构建速度加快。
3. 因为docker是一层一层的，COPY当做一个命令，后面用到这个镜像当做一个命令，可能会使镜像很大。像git的添加后删除一样。一个命令是一层。
4. docker history可以查看各层大小
5. docker save/load 镜像, docker export/import 容器