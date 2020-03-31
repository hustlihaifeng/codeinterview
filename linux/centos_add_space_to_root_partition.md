# 参考

1. <https://blog.csdn.net/onlysingleboy/article/details/38562283>
2. <https://stackoverflow.com/questions/26305376/resize2fs-bad-magic-number-in-super-block-while-trying-to-open>
3. 主要将1中最后一步``
4. 命令

```shell
df -h
fdisk -l
fdisk /dev/sda # n新建，p主分区，4分区号4，w保存
mkfs.ext4 /dev/sda4
reboot
vgdisplay
pvcreate /dev/sda4 
vgextend centos /dev/sda4 
lvdisplay
lvextend /dev/centos/root /dev/sda4 
xfs_growfs /dev/centos/root
df -h
```

