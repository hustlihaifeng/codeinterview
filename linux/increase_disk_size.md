# 给virtualbox虚拟机磁盘扩容

## 替换为大磁盘

### 添加磁盘
1. virtualbox在虚拟机的setting-》storage里面，新建一个虚拟硬盘，默认的`.vdi`格式可以使用时才申请磁盘空间，所以我们可以把磁盘空间上限设置的很大，比如：`1T`（所以最简单的就是一开始就将virtualbox的虚拟硬盘空间设置的很大，超过当前磁盘大小也没关系，反正是用时才申请，平时不占空间。避免后面磁盘扩容的麻烦）。此时编号0，对应/dev/sda；编号1对应/dev/sdb，依次类推。

### 设置自动挂载新磁盘
1. `fdisk -l /dev/sd`按空格键，可以展示有哪些磁盘和分区。假设我们原来的`/dev/sdb`小了，新加的大盘是`/dev/sdc`。
2. 在`/etc/fstab`中,由
```shell
/dev/sdb /data ext4 errors=remount-ro 0 0
```
改为
```shell
/dev/sdc /data ext4 errors=remount-ro 0 0
/dev/sdb /data1 ext4 errors=remount-ro 0 0
```

- 一是设置新加的`/dev/sdc`的自动挂载
- 而是将`/dev/sdc`挂载到`/data`，`/dev/sdb` 改为挂载到`/data1`

3. reboot

### 进行数据迁移
1. 直接将`/data1`下面有用的文件`mv`到`/data`下面。

### 删除老磁盘上的分区
1. `fdisk /dev/sdb`,然后执行`d`删除命令、选择删除的分区、`w`保存命令。
2. 重复第1步，直到`/dev/sdb`上的分区都被删除。
3. `TODO`：删除`/dev/sdb`，没找到好办法，直接从`virtual box`的管理界面删除该磁盘，然后将该磁盘的`.vdi`文件移动到其他目录，会导致虚拟机启动的时候，报错。后面只能恢复该`.vdi`文件，和磁盘序号。

### 总结

1. 这样一通操作下来也蛮麻烦的，最简单的，最开始就把虚拟机的虚拟硬盘大小设置为最大值，比如`2T`，即使磁盘空间本身也没有`2T`，反正是用时才申请。留足了扩展空间，避免以后扩容麻烦。