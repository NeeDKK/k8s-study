## 前言
### 1.前期准备(windows版)
> 1.安装docker desktop
> 
> 2.安装kubectl(https://dl.k8s.io/release/v1.25.0/bin/windows/amd64/kubectl.exe),放入C:\Windows
> 
> 3.安装kind(https://github.com/kubernetes-sigs/kind/releases/),下载后将后缀名改为.exe,放入C:\Windows
> 
> 4.准备dockerhub仓库
### 2.通过kind初始化集群
```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraPortMappings:
    - containerPort: 30080
      hostPort: 30080
      listenAddress: "0.0.0.0"
      protocol: tcp
- role: worker
- role: worker
- role: worker
```
kind create cluster --config=kind-cluster.yaml
### 3.构建容器在k8s-cluster运行
- [ ] 通过dockerfile构建docker镜像
> docker build -t needkk/node-echo:v1.0.1 .
- [ ] 将构建的docker镜像push到docker hub仓库
> docker push needkk/node-echo:v1.0.1
- [ ] 编写manifest文件创建pod，deployment和service
