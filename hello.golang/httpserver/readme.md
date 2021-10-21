作业
- 构建本地镜像。
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
- 将镜像推送至 Docker 官方镜像仓库。
- 通过 Docker 命令本地启动 httpserver。
- 通过 nsenter 进入容器查看 IP 配置

---
# v0.0.2
1. 通过makefile 构建
2. Dockerfile中指定编译环境
3. 运行阶段指定scratch作为基础镜像
```cgo
make push
make run

docker inspect e8fd89bee38b # 查看pid
nsenter -t 16711 -n ip a #查看网络信息
```

# v0.0.1
Dockerfile
```cgo
step0: edit Dockerfile
step1: CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 
step2: trans httpsever file to server
step3: docker build . (docker build . -t httpsever-lgs/ubuntu:v1.0 .)
```

推送到docker 官方
注册账号，登录并创建repo
```cgo
# https://blog.csdn.net/lxy___/article/details/105821141
# docker tag local-image:tagname new-repo:tagname
# docker push new-repo:tagname
docker tag httpsever-lgs/ubuntu:v1.0 jaclond/httpserver-lgs-io:v1.0
docker push jaclond/httpserver-lgs-io:v1.0
```

容器操作
```cgo
 docker run -d -p 8080:8080 --name httpserver 7b15daab0fc8
 
 nsenter -t 4003874 -n ip a
```

最佳实践
参考 https://www.jianshu.com/p/cbce69c7a52f
1. 最小化层
2. 创建缓存
3. 多行参数排序，尽量在末尾修改
4. add和copy 优先使用copy
5. 每个容器只跑一个进程
6. 避免安装不必要的包（使用.dockerignore;在最小目录运行docker build）