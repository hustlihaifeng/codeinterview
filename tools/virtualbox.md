# virtualbox虚拟磁盘空间太大

## 问题

1. 按照[Virtual Box : How to Increase Disk Size - Windows](<https://www.youtube.com/watch?v=7Aqx-VHv2_k>) 去修改虚拟磁盘大小，报错："[Resize medium operation for this format is not implemented yet](https://stackoverflow.com/questions/50772789/how-to-fix-an-error-resize-medium-operation-for-this-format-is-not-implemented)", 无法通过直接改磁盘大小来实现。

## 解决办法
1. 将`/data`中确实有用的文件备份到其他盘。（其他盘没有空间的话，新建一块磁盘）
2. 修改`/etc/fstab`，将`/data`的启东时挂载注释掉。
3. 关闭虚拟机，在虚拟机的`Settings->Storage->Controller:SATA`中选中`/data`对应的磁盘，右键"Remove Attachment"。
4. 在virtualbox主界面的`File->Virtual Media Manager->Hard Disks`, 选中目标磁盘，然后点击上方菜单的Remove。这样会删除旧的虚拟硬盘
5. 选择目标虚拟机，在`Settings->Storage->Controller:SATA`新加一块想要大小的虚拟硬盘。
6. 启动虚拟机：`ls /dev/sd*`，可以看到新加的磁盘`/dev/sdb`，执行`mkfs.ext4 /dev/sdb`将新加的盘设置为`ext4`格式。
7. `ls -tlrha /dev/disk/by-uuid/`可以得到`/dev/sdb`的uuid。然后将`/etc/fstab`中`/data`的uuid替换为新的uuid，解除注释。重启虚拟机。
8. 将原来的文件还原到`/data`目录下。