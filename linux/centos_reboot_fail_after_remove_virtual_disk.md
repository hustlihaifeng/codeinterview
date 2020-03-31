# 背景

win10系统，virtualbox里面安装centos7。后面centos7空间不够，添加了一块vdi虚拟硬盘，自动挂载，设备号是`/dev/sdb`。后面空间有不够，添加另一块虚拟硬盘，设备号是`/dev/sdc`。将第一块虚拟虚拟硬盘上的文件拷贝到第二块虚拟硬盘后，使用`fdisk`删除`/dev/sdb`上的分区，然后关闭虚拟机，在该虚拟机的setting里面删除第一块硬盘。然后启动虚拟机，发现启动失败。等了好久之后，进入一个特殊模式。

# 解决办法

进入特殊模式后，`ls /dev/sd`再按`tab`键，发现原来的`/dev/sdc`竟然变成了`/dev/sdb`，但是`/etc/fstab`里面的自动挂载行还是`/dev/sdc`，所以启动失败。将自动挂载行改为`/dev/sdb`后，启动成功。

# 分析

1. `/dev/sda` `/dev/sdb`只是一个编号，标志第几块硬盘。与virtual box的setting里面的`stat port 0` `stat port 1`没有对应关系。
2. **自动挂载时使用`UUID`，而不要使用设备号(linux系统给的编号)。`UUID`通过`blkid`来查看**。

```shell
# blkid 
/dev/block/253:0: UUID="728f9c9d-6a5c-4618-90cf-05ebdd0626d9" TYPE="xfs" 
/dev/block/8:2: UUID="hDST5O-x0x2-Yop6-xI58-dquB-QZxM-MgDuGn" TYPE="LVM2_member" 
/dev/sda1: UUID="6b3dded3-53a8-4fce-88b8-77131f1dc88b" TYPE="xfs" 
/dev/sdb: UUID="d61bf336-ef93-4587-a551-5c6d26890cae" TYPE="ext4" 
/dev/mapper/centos-swap: UUID="a95e6e69-b508-4906-aad9-b8e1f24c03dd" TYPE="swap"
```

将`/etc/fstab`里面的`/dev/sdb /data ext4 defaults 0 0`改为`UUID=d61bf336-ef93-4587-a551-5c6d26890cae /data ext4 defaults 0 0`

3. `virtualbox`动态获取空间的虚拟硬盘，可以一开始就设置最大空间`2T`，即使宿主机空间只有`100G`，反正是使用时才获取，避免以后扩容麻烦。 