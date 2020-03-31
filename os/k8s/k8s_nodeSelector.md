

1. 打标签：`kubectl label node k8snode1 disktype=ssd`

   1. 同一个对象的labels属性的key必须唯一

2. 用标签：

   ![博客02.png](https://s1.51cto.com/images/20180629/1530240564712480.png?x-oss-process=image/watermark,size_16,text_QDUxQ1RP5Y2a5a6i,color_FFFFFF,t_100,g_se,x_10,y_10,shadow_90,type_ZmFuZ3poZW5naGVpdGk=)

- service 选择 pod 就用到了 label

3. 调度：节点需要满足标签选择器中的所有标签。

   > It specifies a map of key-value pairs. For the pod to be eligible to run on a node, the node must have each of the indicated key-value pairs as labels (it can have additional labels as well).

4. Node Affinity:亲和性调度，它对应的是 `Anti-Affinity`，我们翻译成“互斥”。

- 关键点：`Gt`：label 的值大于某个值（字符串比较）
- 详见：[kubernetes 亲和性调度](<https://cizixs.com/2017/05/17/kubernetes-scheulder-affinity/>)

# 参考资料

1. [深入玩转K8S之利用Label控制Pod位置](https://blog.51cto.com/devingeng/2134064)
2. <https://kubernetes.io/docs/concepts/configuration/assign-pod-node/>
3. [kubernetes 亲和性调度](<https://cizixs.com/2017/05/17/kubernetes-scheulder-affinity/>)