# 删除一个节点

 https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/#tear-down

1. 在控制节点上
    kubectl drain k8s-worker-02 --delete-local-data --force --ignore-daemonsets
    kubectl delete node k8s-worker-02
2. 在被删除的节点上
    kubeadm reset


- 删除控制节点
  kubectl drain k8s-master-01 --delete-local-data --force --ignore-daemonsets
  kubectl delete node k8s-master-01
  kubeadm reset

# 准备

## Installing kubeadm, kubelet and kubectl

1. `kubeadm`: the command to bootstrap the cluster.
2. `kubelet`: the component that runs on all of the machines in your cluster and does things like starting pods and containers.
3. `kubectl`: the command line util to talk to your cluster.

```shell
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
EOF

setenforce 0
sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config
systemctl enable --now kubelet

cat <<EOF >  /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF

systemctl daemon-reload
systemctl restart kubelet
```



# 初始化一个单控制节点集群

https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/

kubeadm init --ignore-preflight-errors=ImagePull

> I0101 15:46:04.340695   23796 version.go:96] could not fetch a Kubernetes version from the internet: unable to get URL "https://dl.k8s.io/release/stable-1.txt": Get https://dl.k8s.io/release/stable-1.txt: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
> I0101 15:46:04.340781   23796 version.go:97] falling back to the local client version: v1.14.4
> [init] Using Kubernetes version: v1.14.4
> [preflight] Running pre-flight checks
> [preflight] Pulling images required for setting up a Kubernetes cluster
> [preflight] This might take a minute or two, depending on the speed of your internet connection
> [preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
>         [WARNING ImagePull]: failed to pull image k8s.gcr.io/kube-apiserver:v1.14.4: output: Error response from daemon: Get https://k8s.gcr.io/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
> , error: exit status 1
>         [WARNING ImagePull]: failed to pull image k8s.gcr.io/kube-controller-manager:v1.14.4: output: Error response from daemon: Get https://k8s.gcr.io/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
> , error: exit status 1
>         [WARNING ImagePull]: failed to pull image k8s.gcr.io/kube-scheduler:v1.14.4: output: Error response from daemon: Get https://k8s.gcr.io/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
> , error: exit status 1
>         [WARNING ImagePull]: failed to pull image k8s.gcr.io/kube-proxy:v1.14.4: output: Error response from daemon: Get https://k8s.gcr.io/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
> , error: exit status 1
> [kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
> [kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
> [kubelet-start] Activating the kubelet service
> [certs] Using certificateDir folder "/etc/kubernetes/pki"
> [certs] Generating "front-proxy-ca" certificate and key
> [certs] Generating "front-proxy-client" certificate and key
> [certs] Generating "etcd/ca" certificate and key
> [certs] Generating "etcd/peer" certificate and key
> [certs] etcd/peer serving cert is signed for DNS names [k8s-master-01 localhost] and IPs [10.4.1.238 127.0.0.1 ::1]
> [certs] Generating "etcd/healthcheck-client" certificate and key
> [certs] Generating "etcd/server" certificate and key
> [certs] etcd/server serving cert is signed for DNS names [k8s-master-01 localhost] and IPs [10.4.1.238 127.0.0.1 ::1]
> [certs] Generating "apiserver-etcd-client" certificate and key
> [certs] Generating "ca" certificate and key
> [certs] Generating "apiserver" certificate and key
> [certs] apiserver serving cert is signed for DNS names [k8s-master-01 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.4.1.238]
> [certs] Generating "apiserver-kubelet-client" certificate and key
> [certs] Generating "sa" key and public key
> [kubeconfig] Using kubeconfig folder "/etc/kubernetes"
> [kubeconfig] Writing "admin.conf" kubeconfig file
> [kubeconfig] Writing "kubelet.conf" kubeconfig file
> [kubeconfig] Writing "controller-manager.conf" kubeconfig file
> [kubeconfig] Writing "scheduler.conf" kubeconfig file
> [control-plane] Using manifest folder "/etc/kubernetes/manifests"
> [control-plane] Creating static Pod manifest for "kube-apiserver"
> [control-plane] Creating static Pod manifest for "kube-controller-manager"
> [control-plane] Creating static Pod manifest for "kube-scheduler"
> [etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
> [wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
> [kubelet-check] Initial timeout of 40s passed.
>
> Unfortunately, an error has occurred:
>         timed out waiting for the condition
>
> This error is likely caused by:
>         - The kubelet is not running
>         - The kubelet is unhealthy due to a misconfiguration of the node in some way (required cgroups disabled)
>
> If you are on a systemd-powered system, you can try to troubleshoot the error with the following commands:
>         - 'systemctl status kubelet'
>         - 'journalctl -xeu kubelet'
>
> Additionally, a control plane component may have crashed or exited when started by the container runtime.
> To troubleshoot, list all containers using your preferred container runtimes CLI, e.g. docker.
> Here is one example how you may list all Kubernetes containers running in docker:
>         - 'docker ps -a | grep kube | grep -v pause'
>                 Once you have found the failing container, you can inspect its logs with:
>                 - 'docker logs CONTAINERID'
> error execution phase wait-control-plane: couldn't initialize a Kubernetes cluster



# k8s-国内源安装

1. <https://gist.github.com/islishude/231659cec0305ace090b933ce851994a>

```shell
kubeadm init --image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers'
```

> I0101 19:48:16.531310   32520 version.go:96] could not fetch a Kubernetes version from the internet: unable to get URL "https://dl.k8s.io/release/stable-1.txt": Get https://dl.k8s.io/release/stable-1.txt: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
> I0101 19:48:16.531386   32520 version.go:97] falling back to the local client version: v1.14.4
> [init] Using Kubernetes version: v1.14.4
> [preflight] Running pre-flight checks
> [preflight] Pulling images required for setting up a Kubernetes cluster
> [preflight] This might take a minute or two, depending on the speed of your internet connection
> [preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
> [kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
> [kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
> [kubelet-start] Activating the kubelet service
> [certs] Using certificateDir folder "/etc/kubernetes/pki"
> [certs] Generating "ca" certificate and key
> [certs] Generating "apiserver" certificate and key
> [certs] apiserver serving cert is signed for DNS names [k8s-master-01 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.4.1.238]
> [certs] Generating "apiserver-kubelet-client" certificate and key
> [certs] Generating "front-proxy-ca" certificate and key
> [certs] Generating "front-proxy-client" certificate and key
> [certs] Generating "etcd/ca" certificate and key
> [certs] Generating "etcd/server" certificate and key
> [certs] etcd/server serving cert is signed for DNS names [k8s-master-01 localhost] and IPs [10.4.1.238 127.0.0.1 ::1]
> [certs] Generating "apiserver-etcd-client" certificate and key
> [certs] Generating "etcd/peer" certificate and key
> [certs] etcd/peer serving cert is signed for DNS names [k8s-master-01 localhost] and IPs [10.4.1.238 127.0.0.1 ::1]
> [certs] Generating "etcd/healthcheck-client" certificate and key
> [certs] Generating "sa" key and public key
> [kubeconfig] Using kubeconfig folder "/etc/kubernetes"
> [kubeconfig] Writing "admin.conf" kubeconfig file
> [kubeconfig] Writing "kubelet.conf" kubeconfig file
> [kubeconfig] Writing "controller-manager.conf" kubeconfig file
> [kubeconfig] Writing "scheduler.conf" kubeconfig file
> [control-plane] Using manifest folder "/etc/kubernetes/manifests"
> [control-plane] Creating static Pod manifest for "kube-apiserver"
> [control-plane] Creating static Pod manifest for "kube-controller-manager"
> [control-plane] Creating static Pod manifest for "kube-scheduler"
> [etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
> [wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
> [apiclient] All control plane components are healthy after 15.503169 seconds
> [upload-config] storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
> [kubelet] Creating a ConfigMap "kubelet-config-1.14" in namespace kube-system with the configuration for the kubelets in the cluster
> [upload-certs] Skipping phase. Please see --experimental-upload-certs
> [mark-control-plane] Marking the node k8s-master-01 as control-plane by adding the label "node-role.kubernetes.io/master=''"
> [mark-control-plane] Marking the node k8s-master-01 as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
> [bootstrap-token] Using token: 1rtxm6.u0d806gntgaoff6e
> [bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
> [bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
> [bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
> [bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
> [bootstrap-token] creating the "cluster-info" ConfigMap in the "kube-public" namespace
> [addons] Applied essential addon: CoreDNS
> [addons] Applied essential addon: kube-proxy
>
> Your Kubernetes control-plane has initialized successfully!
>
> To start using your cluster, you need to run the following as a regular user:
>
> mkdir -p $HOME/.kube
> sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
> sudo chown $(id -u):$(id -g) $HOME/.kube/config
>
> You should now deploy a pod network to the cluster.
> Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
> https://kubernetes.io/docs/concepts/cluster-administration/addons/
>
> Then you can join any number of worker nodes by running the following on each as root:
>
> kubeadm join 10.4.1.238:6443 --token 1rtxm6.u0d806gntgaoff6e \
>  --discovery-token-ca-cert-hash sha256:bab66383e931c38342302147db284359154fde15c5838f3b65dd16ddefbd046f

# 参考资料

1. 安装指定版本的k8s工具：<https://unix.stackexchange.com/questions/151689/how-can-i-instruct-yum-to-install-a-specific-version-of-package-x>
2. k8s安装前期准备：<https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/>
3. k8s安装：<https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/>
4. centos安装k8s：<https://blog.51cto.com/wzlinux/2321767>



# k8s11master3

```shell
kubeadm reset 
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
EOF

setenforce 0
sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config
systemctl enable --now kubelet

cat <<EOF >  /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF

systemctl daemon-reload
systemctl restart kubelet

kubeadm init --image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers'
```

> I0210 16:37:21.056147   13092 version.go:96] could not fetch a Kubernetes version from the internet: unable to get URL "https://dl.k8s.io/release/stable-1.txt": Get https://dl.k8s.io/release/stable-1.txt: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
> I0210 16:37:21.056218   13092 version.go:97] falling back to the local client version: v1.14.4
> [init] Using Kubernetes version: v1.14.4
> [preflight] Running pre-flight checks
> [preflight] Pulling images required for setting up a Kubernetes cluster
> [preflight] This might take a minute or two, depending on the speed of your internet connection
> [preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
> [kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
> [kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
> [kubelet-start] Activating the kubelet service
> [certs] Using certificateDir folder "/etc/kubernetes/pki"
> [certs] Generating "ca" certificate and key
> [certs] Generating "apiserver" certificate and key
> [certs] apiserver serving cert is signed for DNS names [k8s-master-03 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.4.1.44]
> [certs] Generating "apiserver-kubelet-client" certificate and key
> [certs] Generating "front-proxy-ca" certificate and key
> [certs] Generating "front-proxy-client" certificate and key
> [certs] Generating "etcd/ca" certificate and key
> [certs] Generating "etcd/server" certificate and key
> [certs] etcd/server serving cert is signed for DNS names [k8s-master-03 localhost] and IPs [10.4.1.44 127.0.0.1 ::1]
> [certs] Generating "etcd/peer" certificate and key
> [certs] etcd/peer serving cert is signed for DNS names [k8s-master-03 localhost] and IPs [10.4.1.44 127.0.0.1 ::1]
> [certs] Generating "etcd/healthcheck-client" certificate and key
> [certs] Generating "apiserver-etcd-client" certificate and key
> [certs] Generating "sa" key and public key
> [kubeconfig] Using kubeconfig folder "/etc/kubernetes"
> [kubeconfig] Writing "admin.conf" kubeconfig file
> [kubeconfig] Writing "kubelet.conf" kubeconfig file
> [kubeconfig] Writing "controller-manager.conf" kubeconfig file
> [kubeconfig] Writing "scheduler.conf" kubeconfig file
> [control-plane] Using manifest folder "/etc/kubernetes/manifests"
> [control-plane] Creating static Pod manifest for "kube-apiserver"
> [control-plane] Creating static Pod manifest for "kube-controller-manager"
> [control-plane] Creating static Pod manifest for "kube-scheduler"
> [etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
> [wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
> [apiclient] All control plane components are healthy after 16.501590 seconds
> [upload-config] storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
> [kubelet] Creating a ConfigMap "kubelet-config-1.14" in namespace kube-system with the configuration for the kubelets in the cluster
> [upload-certs] Skipping phase. Please see --experimental-upload-certs
> [mark-control-plane] Marking the node k8s-master-03 as control-plane by adding the label "node-role.kubernetes.io/master=''"
> [mark-control-plane] Marking the node k8s-master-03 as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
> [bootstrap-token] Using token: rmwfuh.cx9jceyng3fiu0fv
> [bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
> [bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
> [bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
> [bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
> [bootstrap-token] creating the "cluster-info" ConfigMap in the "kube-public" namespace
> [addons] Applied essential addon: CoreDNS
> [addons] Applied essential addon: kube-proxy
>
> Your Kubernetes control-plane has initialized successfully!
>
> To start using your cluster, you need to run the following as a regular user:
>
>   mkdir -p $HOME/.kube
>   sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
>   sudo chown $(id -u):$(id -g) $HOME/.kube/config
>
> You should now deploy a pod network to the cluster.
> Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
>   https://kubernetes.io/docs/concepts/cluster-administration/addons/
>
> Then you can join any number of worker nodes by running the following on each as root:
>
> kubeadm join 10.4.1.44:6443 --token rmwfuh.cx9jceyng3fiu0fv \
>     --discovery-token-ca-cert-hash sha256:1dd05c9aed380d959748d12484c3dfa413352485d6ba31750f2a45f2faa08341 

```shell
kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"
```

