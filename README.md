# codeinterview
程序员面试指南

# 目录
- algorithm
  - list
    - [01 环形链表检测](algorithm/list/ring_check/01%20环形链表检测.md)
    - [02 奇升偶降单链表重整为降序链表](algorithm/list/reorder/奇升偶降单链表重整为降序链表.md)
    - [03 链表常见操作](algorithm/list/list_operations.md)
  - stack
      - [有序栈](algorithm/stack/sorted_stack/dailyTemperatures/main.md)
  - tree
    - [01 平衡二分查找树、B树、B+树、R树](algorithm/tree/01%20平衡二分查找树、B树、B+树、R树.md)
    - heap
      - [01 堆](algorithm/tree/heap/01%20堆.md)
    - binary_tree
      - [二叉树的深度（先序 中序 后序）和广度优先遍历](algorithm/tree/binary_tree/transfer_of_binary_tree.md)
      - [二叉树的最近公共祖先](algorithm/tree/lowestCommonAncestor/note.md)
  - map
    - leetcode
      - [128 最长连续子集合的长度](algorithm/map/leetcode/128longest_consecutive_sequence/128longest_consecutive_sequence.md)
  - 数组
    - [三数之和为0](algorithm\slice\threeSum\note.md)  [两数之和为固定值](algorithm/map/leetcode/1twoSum/main.md)
    - [岛屿的最大面积](algorithm/slice/maxAreaOfIsland/try20200228/note.md)
    - [搜索旋转排序数组](algorithm/slice/search/try20200229/note.md)
    - [最长连续递增序列](algorithm/slice/findLengthOfLCIS/try20200229/note.md)
    - [最长连续序列](algorithm/slice/longestConsecutive/try20200301/note.md)
    - [数组中的第K个最大元素](algorithm/slice/findKthLargest/try20200301/note.md)
    - [第k个排列](algorithm/slice/getPermutation/try20200303/note.md)
    - [朋友圈](algorithm/slice/findCircleNum/try20200303/note.md)
    - [合并区间](algorithm/slice/merge/try20200305/note.md)
    - [接雨水](algorithm/slice/trap/try20200305/note.md)
    - [一维有序数组二维化后的二分查找](algorithm/slice/searchMatrix/try20200404/main.md)
  - 动态规划和贪心
    - [买卖股票的最佳时机](algorithm/dp_greedy/maxProfit/try20200308/main.go)
    - [最大正方形面积问题](algorithm/dp_greedy/maximalsquare/maximalsquare.md)
    - [最大连续子数组之和](algorithm/dp_greedy/maxSubArray/maxSubArray.md)
    - [三角形型二维数组的最短路径和](algorithm/dp_greedy/minimumTotal/minimumTotal.md)
    - [最长递增数组长度问题](algorithm/dp_greedy/lengthOfLIS/lengthOfLIS.md)
    - [俄罗斯套娃问题](algorithm/dp_greedy/maxEnvelopes/maxEnvelopes.md)
  - 数据结构
    - [最小栈](algorithm/data_structure/MinStack/MinStack.md)
    - [全O(1)的LRUCache](algorithm/data_structure/LRUCache/LRUCache.md)
    - [全 O(1) 的数据结构](algorithm/data_structure/AllOne/try20200321/main.go)
  - string
    - [最长的无重复字符的子串长度](algorithm/string/lengthOfLongestSubstring/lengthOfLongestSubstring.md)
    - [最长公共前缀](algorithm/string/longestCommonPrefix/try20200220\main.go)
    - [字符串的排列](algorithm/string/checkInclusion/try20200220/main.go)
    - [字符串相乘](algorithm/string/multiply/try20200221/main.go)
    - [字符串内单词间反转](algorithm/string/reverseWords/reverseWords.md)
    - [简化路径](algorithm/string/simplifyPath/try20200222/main.go)
    - [复原IP地址](algorithm/string/restoreIpAddresses/main.go)
  - 其他
    - [x 的平方根](algorithm/other/mySqrt/main.go)
    - [第二高的薪水](algorithm/other/SecondHighestSalary/main.sql)
  - 实际应用
      - [一致性哈希算法](algorithm/practice/consistent_hash/introduce_to_consistent_hash.md)
- markdown
  - [01 markdown里的连接](markdown/01%20连接.md)
- redis
  - [redis数据结构的常见用法和使用场景](redis/common_usage_and_target_situation_of_redis_data_structure.md)
  - [01 redis持久化方案](redis/01%20redis持久化方案.md)
  - [redis事物与ACID](redis/02%20redis事物与ACID.md)
  - [TODO:redis_cluster](redis/redis_cluster.md)
- mysql
  - [01 InnoDB和MyISAM](mysql/01%20InnoDB和MyISAM.md)
  - [02 MySQL大表优化](mysql/02%20MySQL大表优化.md)
  - [03 DDL DML 事物和delete truncate drop](mysql/03%20DDL%20DML%20and%20transaction.md)
- hadoop
  - hdfs
    - [hdfs架构原理和基本操作](hadoop/hdfs/hdfs_basic.md)
    - [搭建CDH](hadoop/hdfs/CDH_setup.md)
    - [hdfs加密-1](hadoop/hdfs/hdfs_encrypt.md)
    - [hdfs加密-2](hadoop/hdfs/hdfs_encrypt_from_simon.md)
    - [hdfs配置项](hadoop/hdfs/hdfs_config.md)
- oracle
  - [01 oracle环境搭建与常见问题](oracle/oracle_basic.md)
  - [02 beego orm从mysql切换到oracle](oracle/migrate_from_beego_orm_of_mysql_to_oracle.md)
- golang
  - sync
    - [01 等待所有goroutine结束](golang/sync/01%20等待所有goroutine结束.md)
    - [02 原子操作和并发安全的map实现](golang/sync/02_atomic.md)
    - [sync.Map](golang/sync/map/Map/main.go)
    - [锁](golang/sync/lock.md)
  - channel
    - [01 阻塞式channel和非阻塞式channel](golang/channel/01%20阻塞式channel和非阻塞式channel.md)
  - slice
    - [01 切片和数组](golang/slice/01%20切片和数组.md)
  - io
    - bufio
      - [ReadLine长行截断问题](golang/io/bufio/ReadLine.md)
  - rpc
    - [golang_rpc框架调研](golang/rpc/golang_rpc框架调研.md)
  - websocket
    - [websocket注意事项和官方样例](golang/websocket/websocket.md)
  - reflect
    - [reflect](golang/reflect/reflect.md)
  - beego
    - [beego配置模块](golang/beego/config_of_beego.md)
  - cache
    - [TODO:beego cache](golang/cache/cache.md)
  - gc
    - [TODO：Go GC 20 问](golang/gc/learning_go_gc_20_question.md)
  - 依赖管理
    - [go module](golang/module/introduction_go_module.md)
    - [go vendor](golang/vendor/govendor_and_git_submodule.md)
  - net
    - [改golang源码后如何重新编译](golang/net/http/httputil/httputil.md)
  - toolbox
    - [TODO:定时任务](golang/toolbox/toolbox.md)
  - routine
    - [协程 线程 进程 go协程池](golang/routine/routine_thread_process.md)
- python
  - [win10 python2.75 安装pip报错](python/install.md)
  - [TODO:python_to_xls](python/python_to_xls.md)
  - [python_use_mysql](python/python_use_mysql.md)
  - [python项目源码分析](python/source_code_analysis_of_data_analysis.md)
  - [venv](python/venv.md)
  - [flask](python/flask.md)
- ml
  - cs229stanford
    - [class01](ml/cs229stanford/class01.md)
    - [02_gradient_descent](ml/cs229stanford/02_gradient_descent.md)
    - [class02linear_regression_gradient_descent_normal_equation](ml/cs229stanford/class02linear_regression_gradient_descent_normal_equation.md)
    - [class03underfitting_and_overfitting](ml/cs229stanford/class03underfitting_and_overfitting.md)
- linux
  - [linux进程间通信](linux/communication%20between%20processes%20in%20linux.md)
  - [netstat](linux/netstat.md)
  - [tcp time_wait问题](linux/time_wait_of_tcp.md)
  - [tcp粘包问题](linux/tcp粘包问题.md)
  - [tcp发送速度控制](linux/speed_control_of_tcp.md)
  - [linux常见资源监控命令](linux/resource%20monitor%20of%20linux.md)
  - [给virtualbox虚拟机替换大硬盘](linux/increase_disk_size.md)
  - [linux常见设置](linux/common_setting_in_linux.md)
  - [一次ftp出错](linux/setup_vsftp_server_in_centos.md)
  - [address already in use](linux/address_already_in_use.md)
  - [expect](linux/expect.md)
  - [ftp](linux/setup_vsftp_server_in_centos.md)
  - [搭建nfs系统](linux/setup_nfs_system.md)
    - [nfsv4.1_mount_hang](linux/nfsv4.1_mount_hang.md)
  - [centos_add_space_to_root_partition](linux/centos_add_space_to_root_partition.md)
  - [centos_reboot_fail_after_remove_virtual_disk](linux/centos_reboot_fail_after_remove_virtual_disk.md)
  - [diff_and_patch_in_linux](linux/diff_and_patch_in_linux.md)
  - [upgrade_linux_kernel_without_network](linux/upgrade_linux_kernel_without_network.md)
  - [文件 文件描述符 文件句柄](linux/file_descriptor.md)
  - [守护进程的原理以及go程序的特殊性和解决办法](linux/daemon_process.md)
- windows
  - [windows远程连接](windows/remote_desktop_of_windows.md)
- docker

  - [docker笔记](docker/docker%20note.md)
  - [docker-compose](docker/docker-compose/basic_of_docker_compose.md)
  - [docker常见操作](docker/docker_options.md)
  - [docker数据目录设置与迁移](docker/docker_root_dir.md)
  - [如何在镜像中安装rpm依赖](docker/docker_rpm_install.md)
  - [dockerfile](docker/dockerfile.md)
    - [dockerfile多阶段构建](docker/multi_state_dockerfile.md)
  - [docker搭建mysq_ redis_registry_python服务](docker/docker_setup_mysql_redis_and_so_on.md)
  - [docker隔离机制](docker/docker_isolation.md)
- os
  - k8s
    - [k8s笔记](os/k8s/k8s_note.md)
    - [k8s优先级](os/k8s/PriorityClass_of_k8s.md)
    - [TODO:k8s容器类型](os/k8s/container_type_of_k8s.md)
    - [k8s_nodeSelector](os/k8s/k8s_nodeSelector.md)
    - [TODO:local_persistent_volumes_of_k8s](os/k8s/local_persistent_volumes_of_k8s.md)
    - [安装k8s](os/k8s/setup_k8s.md)
    - [k8s网络](os/k8s/k8s_network.md)
- 技巧
  - [how to code](skills/how_to_code.md)
  - [how to learn](skills/how_to_learn.md)
  - [如何准备项目经历面试](skills/how_to_prepare_for_previous_project.md)
  - [常见问题](skills/common_question.md)
  - [如何避坑](attention.md)
  - [常见面试题](skills/common_question.md)
- tools
  - [vim](tools/vim.md)
  - [tmux](tools/tmux.md)
  - [git](tools/git.md)
  - [processon](tools/processon.md)
  - [virtualbox](tools/virtualbox.md)
  - [chrome](tools/chrome.md)
- network
  - [SO_REUSEADDR和SO_REUSEPORT](network/SO_REUSEADDR_AND_SO_REUSEPORT.md)
  - [RTP](network/RTP.md)
  - [RTO：TCP超时重传时间计算](network/RTT_RTO.md)
  - [TCP三次握手和四次挥手](network/tcp_say_hello_and_goodbye.md)
  - [多路复用：epoll和select](network/epoll_select.md)
  - [socket关闭时的read、wirte行为和抓包工具](network/unix_network_program/read_write_close_shutdown_and_fiddler.md)
- 安全
  - [常见安全算法分类和https](safe/common_safe_algorithm.md)
- 设计模式
  - [常见设计模式及其应用场景](design_pattern/common_design_pattern_and_its_problem.md)
- 架构设计
  - [如何设计一个百万级用户的抽奖系统](architecture/如何设计一个百万级用户的抽奖系统.md)
  - [微服务](architecture/micro_service.md)
  - [长链接和短连接](architecture/long_connection_and_short_connection.md)
  - [负载均衡](architecture/常见负载均衡方案.md)
  - [分布式事物](architecture/distributed_consensus_problem.md)
  - [12306的难点和解决办法](architecture/problem_and_skill_in_12306.md)
  - [服务热更新](architecture/go_server_hot_upgrade.md)
  - [常用概念](architecture/annotions.md)
  - [https性能优化](architecture/https_speed_up.md)
  - [分布式锁如何实现](architecture/distributed_lock.md)
- nginx
  - [nginx负载均衡的常见方案](nginx/common_nginx_load_banlance_algorithm.md)
