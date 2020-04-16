[TOC]

# epoll为啥比select高效

1. epoll是通过读写事件触发来将有读写事件的socket存储到内核缓存list中，调用epoll只需要返回内核中list的一部分即可；select是在调用的时候进行轮询检查每一个监听的套接字。
2. select调用时需要拷贝fd_set到内核空间；epoll通过`epoll_create`,`epoll_ctl`这两个函数来事先注册进而避免了每次epoll的时候拷贝，其中`epoll_ctl`给内核中断处理程序注册一个回调函数，告诉内核，如果这个句柄的中断到了，就把它放到准备就绪list链表里。



# 参考资料

1. [select和epoll区别](<https://www.jianshu.com/p/430141f95ddb>)