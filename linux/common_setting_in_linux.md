# 1. 设置命令行前缀为文件夹名称而非完整路径

```shell
echo $PS1
# 完整路径
export PS1='export PS1='${debian_chroot:+($debian_chroot)}\u@\h:\w\$ ''
# 最后一个文件夹名称
export PS1='export PS1='${debian_chroot:+($debian_chroot)}\u@\h:\W\$ ''
```
