[TOC]

1. **轮询**(不需要额外配置),**最少连接**(`least_conn;`),**IP地址哈希**(`ip_hash;`将同一个ip分配到同一个后端服务中，比如方便处理会话Session)，**基于权重的负载均衡**(`weight=2`,默认是1. 可以和和基于IP地址哈希的负载均衡一起使用)
2. 配置详见：[Nginx负载均衡的4种方案配置](<https://blog.csdn.net/xiaoshanghe/article/details/54865819>)

# 参考资料

1. [Nginx负载均衡的4种方案配置](<https://blog.csdn.net/xiaoshanghe/article/details/54865819>)