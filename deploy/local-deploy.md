# 通过kind在本机部署分布式集群

## 准备yaml
- kind特有的
```yaml
# kind-config.yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane # master
    image: kindest/node:v1.27.0
    #在本地映射端口，让外部访问集群
    extraPortMappings: 
      - containerPort: 30080
        hostPort: 80
        protocol: TCP
  - role: worker
    image: kindest/node:v1.27.0
  - role: worker
    image: kindest/node:v1.27.0
networking:
  disableDefaultCNI: false
  podSubnet: "10.244.0.0/16" # Pod 运行时会从这个子网中分配 IP
  serviceSubnet: "10.96.0.0/12" # 意思同上
  kubeProxyMode: "iptables"

```
## MySQL部署
```bash
# deploy & check
kubectl apply -f mysql-pv-pvc.yaml

kubectl apply -f mysql-init-configmap.yaml
kubectl get configmap mysql-init-sql -o yaml

kubectl apply -f mysql-statefulset.yaml
kubectl get statefulset domtok-mysql -o yaml
kubectl get svc domtok-mysql -o yaml
```

为了调试方便，这里svc的启动是以nodePort方式启动的
```bash
➜  deploy git:(deploy) ✗ kubectl get svc domtok-mysql -o wide        

NAME           TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE   SELECTOR
domtok-mysql   NodePort   10.96.252.96   <none>        3306:32563/TCP   79m   app=domtok-mysql

```
这里可以得知，MySQL 监听在 **集群中任意节点的** IP:32563(这是启用NodePort才有的)
所以我们可以通过任意一个节点的ip去连入
```bash
➜  deploy git:(deploy) ✗ kubectl get nodes -o wide

NAME                   STATUS   ROLES           AGE   VERSION   INTERNAL-IP   EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION     CONTAINER-RUNTIME
domtok-control-plane   Ready    control-plane   14h   v1.25.3   172.18.0.5    <none>        Ubuntu 22.04.1 LTS   6.8.0-47-generic   containerd://1.6.9
domtok-worker          Ready    <none>          14h   v1.25.3   172.18.0.3    <none>        Ubuntu 22.04.1 LTS   6.8.0-47-generic   containerd://1.6.9
domtok-worker2         Ready    <none>          14h   v1.25.3   172.18.0.4    <none>        Ubuntu 22.04.1 LTS   6.8.0-47-generic   containerd://1.6.9

mysql -h 192.168.1.100 -P 32563 -u domtok -p

```

## redis
```bash
kubectl apply -f redis.yaml
kubectl get pv,pvc,sts,svc,pod

kubectl apply -f etcd-configmap.yaml
kubectl apply -f etcd.yaml
kubectl exec -it domtok-etcd-0 -- bash
etcdctl get /config
```
