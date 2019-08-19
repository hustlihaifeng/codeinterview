1. 参考 <https://www.unixmen.com/install-configure-ftp-server-centos-7/>  <https://linuxize.com/post/how-to-setup-ftp-server-with-vsftpd-on-centos-7/> 
2. 关于账号：按照 <https://linuxize.com/post/how-to-setup-ftp-server-with-vsftpd-on-centos-7/#2-enabling-uploads> 错误的将账号`haifeng`加入 `/etc/vsftpd/user_list`，导致登录`WinScp`时报错:

```shell
验证失败。
连接失败
Permission denied.
```

将`haifeng`从 `/etc/vsftpd/user_list`中移除，就可以登录了。

3. 关于`winscp`链接协议：要选择`FTP`,选择`SFTP`和`SCP`都报错。
4. 关于端口：默认是21，`vsftpd.conf`的`listen_port`会改变这个监听端口，链接时应该链接`listen_port`指定的端口。
5. 关于登录工具：`WinScp`可以登录，直接在windows某个文件夹上面的地址中输入：`ftp://user:password@ip:port/`也能登录查看，建议使用`WinScp`，省事。在`WinScp`可以的时候，在浏览器中用`url`登录报错了，么有详细去定位什么错误。

6. 上述过程的配置：

   ```shell
   anonymous_enable=NO
   local_enable=YES
   write_enable=YES
   local_umask=022
   dirmessage_enable=YES
   xferlog_enable=YES
   xferlog_std_format=YES
   chroot_local_user=YES
   chroot_list_enable=NO
   chroot_list_file=/etc/vsftpd/chroot_list
   listen=YES
   listen_ipv6=NO
   
   pam_service_name=vsftpd
   userlist_enable=YES
   tcp_wrappers=YES
   listen_port=xx
   pasv_min_port=xx
   pasv_max_port=xx
   pasv_enable=YES
   pasv_address=xx.xx.xx.xx
   pasv_promiscuous=YES
   allow_writeable_chroot=YES
   ```

   

   ```shell
   # firewall-cmd --state        
   not running
   ```

   与上面两个链接中的配置比对这个看。暂时不深究。

