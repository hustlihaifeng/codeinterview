# 环境准备

1. 在virtualbox中安装一台centos虚拟机
2. clone原虚拟机
   1. 原虚拟机-》右键-》clone
   2. 新虚拟机-》右键-》setting-》network-》去掉adapter 2-》勾选adapter 3-》选择hostonly网络
   3. 启动新虚拟机-》`cd /etc/sysconfig/network-scripts/`-》`vimdiff ifcfg-enp0s3 ifcfg-enp0s9`-》除了`NAME`、`UUID`、`DEVICE`之外，都将新的改为和老的一样。
3. 启动新老机器，得到`ip`分别为`192.168.120.3`、`192.168.120.4`，ssh连接进去。

# 搭建nfs系统

## server

1. 在`/etc/hosts`中添加：

```shell
192.168.120.3 client
192.168.120.4 nfs-server
```

2. 关闭防火墙

   1. ```shell
      systemctl stop firewalld
      systemctl disable firewalld
      ```

   2. `vi /etc/selinux/config`

   ```shell
   SELINUX=disabled
   ```

3. `yum -y install nfs-utils`

4. ```shell
   systemctl enable rpcbind
   systemctl enable nfs-server
   systemctl enable nfs-lock
   systemctl enable nfs-idmap
   systemctl start rpcbind
   systemctl start nfs-server
   systemctl start nfs-lock
   systemctl start nfs-idmap
   ```

5. `mkdir /sharednfs`

6. `vim /etc/exports`:

```shell
/sharednfs client(rw,sync,no_root_squash,no_subtree_check)
```

7. ```shell
   exportfs -a
   systemctl restart nfs-server
   ```

8. 重启测试

## client

1. 在`/etc/hosts`中添加：

```shell
192.168.120.3 client
192.168.120.4 nfs-server
```

2. 关闭防火墙

   1. ```shell
      systemctl stop firewalld
      systemctl disable firewalld
      ```

   2. `vi /etc/selinux/config`

   ```shell
   SELINUX=disabled
   ```

3. `yum -y install nfs-utils`

4. ```shell
   systemctl enable rpcbind
   systemctl enable nfs-server
   systemctl enable nfs-lock
   systemctl enable nfs-idmap
   systemctl start rpcbind
   systemctl start nfs-server
   systemctl start nfs-lock
   systemctl start nfs-idmap
   ```

5. `mkdir -p /mnt/sharednfs`

6. `mount -t nfs 192.168.120.4:/sharednfs /mnt/sharednfs`

7. 自动挂载：`vim /etc/fstab`:

```shell
nfs-server:/sharednfs   /mnt/sharednfs/ nfs defaults 0 0
```

8. 在`192.168.120.4:/sharednfs`新建一个文件，在`192.168.120.3:/mnt/sharednfs`看有没有
9. 重启测试自动挂载

# 参考资料

1. [How To Setup an NFS Mount on CentOS 7](<https://www.globo.tech/learning-center/setup-nfs-mount-centos-7/>)

2. [21.2.1. Mounting NFS File Systems using `/etc/fstab`](<https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/4/html/System_Administration_Guide/Network_File_System_NFS-Mounting_NFS_File_Systems.html>)

