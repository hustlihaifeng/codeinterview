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
      
      service iptables status
      service ip6tables status
      service iptables stop
      service ip6tables stop
      chkconfig iptables off
      chkconfig ip6tables off
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

   ```shell
   df -h|grep ":/"  将所有的挂载umount
   systemctl stop nfs-idmap nfs-lock nfs-server rpcbind
   systemctl disable nfs-idmap nfs-lock nfs-server rpcbind
   yum remove nfs-utils
   yum remove libnfsidmap
   ```

   

5. `mkdir -p /mnt/sharednfs`

6. `mount -t nfs 192.168.120.4:/sharednfs /mnt/sharednfs`

7. 自动挂载：`vim /etc/fstab`:

```shell
nfs-server:/sharednfs   /mnt/sharednfs/ nfs defaults 0 0
```

8. 在`192.168.120.4:/sharednfs`新建一个文件，在`192.168.120.3:/mnt/sharednfs`看有没有
9. 重启测试自动挂载
10. 不重启加载配置：`exportfs -arv`

# nfs权限配置

1. 账户映射关系：
   - no_root_squash：客户端使用的是root用户时，则映射到FNS服务器的用户依然为root用户

   - root_squash：将客户端使用的是root用户时，则映射到NFS服务器的用户为NFS的匿名用户（nfsnobody）.

     - 在只配置这个的时候

     > 如果客户端所使用的用户身份不是root，而是一个普通用户，那么默认情况下在服务器端会被映射为服务端同id用户，如果不是文件owner，那么就是chmod 750中的0。

     - 默认的权限规则（root_squash）是，root用户被映射成nfsnobody（65534）用户，对于客户端机器上和NFS服务器上UID相同的用户会对应映射（haifeng<->client uid 1000<->server uid 1000<->xxxx_with_id_1000），其它非root用户被映射成nobody（99）用户。当root用户访问共享目录时是以nfsnobody用户访问共享目录的，具有什么权限看下共享目录权限便知。客户端机器上和NFS服务器上的相同UID用户，以NFS服务器上的用户访问共享目录，看戏目录权限便知。其它非root用户则映射成nobody用户，有啥权限一看便知。(见:<https://blog.csdn.net/donghaixiaolongwang/article/details/79230220>)

     

   - all_squash：选项，将所有访问NFS服务器的客户端的用户都映射为匿名用户，不管客户端使用的是什么用户。

     - 最简单，UID一一对应即可：

       ```shell
       vim /etc/exports
       
        /mount/www  10.1.8.0/24(rw)   ##共享/mount/www目录 给10.1.8.0网段机器。并且是读写的
       
       ====================
       
       useradd  USER_NAME   ##在NFS服务器上建立用户，和客户端建立的用户UID一样（必须）
       
       useradd USER_NAME  ##在客户端机器上建立用户，和NFS服务器上建立的用户UID一样（必须）
       
       chown USER_NAME   /mount/www  ##NFS服务器上的共享目录改好
       
       ====================
       
       如果还需要第二个用户，则再在客户端和NFS服务端机器上配对建立UID一样的用户。同时在NFS服务端对共享目录使用"setfacl  -m u:USER_NAME2:rwx"增加第二个用户的权限即可。
       ```

       

   - anonuid：本地的匿名用户的UID

     - /mount/www  10.1.8.0/24(rw,all_squash,anonuid=500,anongid=500). 将所有的账户映射为本地的500账户，500组。也就是不同的客户端可以映射到不同的账户。 这个已经验证。（<http://cn.linux.vbird.org/linux_server/0330nfs/0330nfs-centos4.php#What_NFS_perm>）

   - anongid：本地的匿名用户的GID

   - 用户身份重叠

     - <https://blog.51cto.com/yttitan/2406403>

       > 在使用NFS共享的过程中，有时还可能会遇到用户身份重叠的问题。所谓用户身份重叠，是指在NFS服务采用默认设置（用户身份映射选项为root_squash）时，如果在服务器端赋予某个用户对共享目录具有相应权限，而且在客户端恰好也有一个具有相同uid的用户，那么当在客户端以该用户身份访问共享时，将自动具有服务器端对应用户的权限。下面举例予以说明。

2. 读写方式：
- ro：默认选项，以只读的方式共享。
- rw：以读写的方式共享。

3. [如何确保NFS服务安全](<https://cloud.tencent.com/developer/article/1072369>)
4. 默认配置项：<https://www-uxsup.csx.cam.ac.uk/pub/doc/redhat/redhat7.3/rhl-rg-en-7.3/s1-nfs-server-config.html>

> - `ro` Read-only. Hosts mounting this filesystem will not be able to change it. To allow hosts to make changes to the filesystem, you must specify `rw` (read-write).
> - `async` Allows the server to write data to the disk when it sees fit. While this is not important if the host is accessing data as read-only, if a host is making changes to a read-write filesystem and the server crashes, data could be lost. By specifying the `sync` option, all file writes must be committed to the disk before the write request by the client is actually completed. This may lower performance.
> - `wdelay` Causes the NFS server to delay writing to the disk if it suspects another write request is imminent. This can improve performance by reducing the number of times the disk must be accessed by separate write commands, reducing write overhead. Use `no_wdelay` to turn this feature off, which only works if you are using the `sync` option.
> - `root_squash` Makes any client accesses to the exported filesystem, made as the root user on the client machine, take place as the nobody user ID. This effectively "squashes" the power of the remote root user to the lowest local user, preventing remote root users from acting as though they were the root user on the local system. Alternatively, the `no_root_squash` option turns off root squashing. To squash every remote user, including root, use the `all_squash` option. To specify the user and group IDs to use with remote users from a particular host, use the `anonuid` and `anongid` options, respectively. In this way, you can create a special user account for remote NFS users to share and specify `(anonuid=*<uid-value>*,anongid=*<gid-value>*)`, where `*<uid-value>*` is the user ID number and `*<gid-value>*` is the group ID number.

5. nfs 705 可以，704不行，706不行。也就是必须有rx权限，才能读取。其他人有rw就会变成绿色。
6. 关于id相等：server上id为1000的是haifeng，client上id为1000的是xxxx。在server上将共享目录owner设置为haifeng，在client上查看，owner会变成xxxx。进而client上的xxxx可以访问共享目录。因为id相同。
7. 客户端和服务端同uid，服务端对应user不是共享文件的owner，但是所在group可以访问共享文件夹。此时在客户端依然不能访问共享文件夹。客户端user有服务端共享文件加的同group，依然不能再客户端访问共享文件夹。
8. 不同的共享目录是不同的组的，不同的目录设置不同的anonuid，将共享目录权限设置为700，实现不同组之间通过用户来隔离。
9. 组操作
   创建组：sudo groupadd -g 1016 nfsuser
   修改组：usermod -g 组名 haifeng
   添加组：usermod -a -G 组名 用户名
10. 例子

```shell
drwxr-x--- 2 1001   1003   4.0K Dec 25 21:52 user1001
drwxr-x--- 2 1002   1003   4.0K Dec 26 09:49 user1002
```

```shell
/data/nfs/user1001 10.4.1.87(rw,all_squash,anonuid=1001,anongid=1003) 10.4.1.2(rw,anonuid=1001,anongid=1001)
/data/nfs/user1002 10.4.1.87(rw,all_squash,anonuid=1002,anongid=1003) 10.4.1.2(rw,anonuid=1002,anongid=1002)
```

- 10.4.1.2上，1001账号能访问user1001,不能访问user1002；1002账号能访问user1002，不能访问user1001。root账户能访问所有者两个目录。
- 87上，all_squash, 所以任何一个账户在user1001上被映射为1001，在user1002给映射为1002，所以都能访问。

# 参考资料

1. [How To Setup an NFS Mount on CentOS 7](<https://www.globo.tech/learning-center/setup-nfs-mount-centos-7/>)

2. [21.2.1. Mounting NFS File Systems using `/etc/fstab`](<https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/4/html/System_Administration_Guide/Network_File_System_NFS-Mounting_NFS_File_Systems.html>)

