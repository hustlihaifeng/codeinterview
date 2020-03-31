1. nfs v4.1 挂载有可能出现一直等待服务端相应的问题。4.0不会，也不会。
2. https://github.com/kubernetes/kubernetes/issues/24721#issuecomment-277741470  这里说可以通过修改/etc/nfsmount.conf来指定mount nfs的默认版本。

```shell
[ NFSMount_Global_Options ]
Defaultvers=3
Nfsvers=3
```

3. `netstat -anp|grep "2049"`可以看到nfs的网络请求状况，2049是v4的服务端端口。执行

```shell
systemctl enable rpcbind
systemctl enable nfs-server
systemctl enable nfs-lock
systemctl enable nfs-idmap
systemctl start rpcbind
systemctl start nfs-server
systemctl start nfs-lock
systemctl start nfs-idmap
```

后，才会起服务占用2049端口。

4. 卸载nfs：

```shell
df -h|grep ":/"  将所有的挂载umount
systemctl stop nfs-idmap nfs-lock nfs-server rpcbind
systemctl disable nfs-idmap nfs-lock nfs-server rpcbind
yum remove nfs-utils
yum remove libnfsidmap
```

