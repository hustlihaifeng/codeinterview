# [Kubernetes 1.14: Local Persistent Volumes GA](https://kubernetes.io/blog/2019/04/04/kubernetes-1.14-local-persistent-volumes-ga/)

 local PV 在1.7被引入，1.10发布beta版，1.14发布GA版，GA代表可以用于生产环境。

# What is a Local Persistent Volume?

1. 一个local PV代表一个绑定到某一个节点的local disk。
2. 与nfs这样的remote storage相比，local PV使用相同的api来能够提供更好的性能。

# How is it different from a HostPath Volume?

1. k8s调度器会将使用local PV的pod调度到local pv所在的节点。k8s知道local pv在哪个节点。

2. HostPath volumes可以通过Persistent Volume Claim (PVC)或者directly inline in a pod definition来访问。但是local PV只能通过一个PVC来访问。这样做更安全，因为Persistent Volume是被admin管理的，组织pod直接访问宿主机的任何路径。

# What’s New With GA?
1. he ability to specify a raw block device and have Kubernetes automatically format and mount the filesystem

# Limitations of GA

1. At GA, Local Persistent Volumes do not support dynamic volume provisioning.

# How to Use a Local Persistent Volume?
1. 

# 参考资料

1. <https://kubernetes.io/blog/2019/04/04/kubernetes-1.14-local-persistent-volumes-ga/>