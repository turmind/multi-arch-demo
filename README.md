# MULTI-ARCH-DEMO

## package

### linux arm

```linux
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GO111MODULE=on go build -a -o bin/echo-server_arm
```

### linux indel

```linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o bin/echo-server_indel
```

## 使用方法

### 执行方法

```linux
chmod +x echo-server
./echo-server -n 2
```

### 相关参数

```linux
-n 监听端口数据，从4000开始，默认为2
```

## reference

[arm ami节省成本](https://docs.aws.amazon.com/zh_cn/eks/latest/userguide/eks-optimized-ami.html#arm-ami)
[multi-arch build](https://salesjobinfo.com/multi-arch-container-images-for-docker-and-kubernetes/)
[modern multi-arch](https://docs.docker.com/desktop/multi-arch/)
[old school multi-arch](https://aws.amazon.com/cn/blogs/containers/introducing-multi-architecture-container-images-for-amazon-ecr/)
