# 



# 参考资料

1. 官方文档： <https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/>

   1. <https://kubernetes.io/docs/concepts/policy/resource-quotas/>

2. 如何使用：[k8s学习(十七) 配置并使用storageclass](https://blog.csdn.net/u011943534/article/details/100887530)

   1. [开启Kubernetes的抢占模式](<http://weekly.dockerone.com/article/9097>)

   ```shell
   # cat ./priority_class_high.yaml
   apiVersion: scheduling.k8s.io/v1
   kind: PriorityClass
   metadata:
     name: high-priority
   value: 1000000
   globalDefault: false
   description: "This priority class should be used for XYZ service pods only."
   # kubectl apply -f ./priority_class_high.yaml
   priorityclass.scheduling.k8s.io/high-priority created
   # kubectl get priorityclasses.scheduling.k8s.io
   NAME                      VALUE        GLOBAL-DEFAULT   AGE
   high-priority             1000000      false            8s
   system-cluster-critical   2000000000   false            28d
   system-node-critical      2000001000   false            28d
   ```

   