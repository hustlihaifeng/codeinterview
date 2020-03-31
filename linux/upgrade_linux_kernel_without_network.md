# 在线升级方案

## 1. [获取新的iso包作为repo源来升级](http://www.unixarena.com/2015/12/how-to-patch-update-rhel-7-without-internet-connection.html/)

## 2. 下载好依赖后升级

1. [在虚拟机中准备好老版本环境，更新源为国内的比较快的源](TODO)
2. 使用downloadonly来获取新kernel的所有依赖包，详见[How To Download A RPM Package With All Dependencies In CentOS](https://www.ostechnix.com/download-rpm-package-dependencies-centos/). (一个问题时，安装downloadonly时安装的表可能没有在这个列表，需要重试下)
3. 将这些包安装：有依赖关系的包放到通一行来安装，这样rpm就不会显示依赖冲突了。使用`rpm -Uvh`,`U`代表upgrade/install，会删除旧版本包。
4. 安装kernel，使用`rpm -ivh`,`i`指install，不会删除老的kernel
5. 重启

