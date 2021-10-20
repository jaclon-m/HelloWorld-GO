学习优秀的写法
https://github.com/idefav/idefav-httpserver/blob/master/README.md

# 如何编写 Handler

1. 在 handler 文件夹添加 handler 文件夹, 然后添加 demo.go 和 init.go 文件
2. 在 demo.go 文件添加 DemoHandler struct
3. DemoHandler 实现 HandlerMapping 接口 HandlerMapping 包含的函数说明:
    - Name() 返回 Handler 名称
    - Path() 返回 Handler 匹配路径
    - Method() 返回 Handler 可以处理的 HttpMethod 
    - Handler() 开发真正的处理逻辑
4. 在 init.go 编写 包初始化函数
```cgo
func init(){
// init
headerz:=HeaderHandler{}
handler.DefaultDispatchHandler.AddHandler(&headerz)
}
```
5. 在 server.go 新增包引用
```cgo
import (
//...
_ "idefav-httpserver/handler/headerz"
)
```
6. 启动服务

# 整体框架

1. cfg配置文件中写入常量、配置启动监听地址
2. main函数启动时
   1. 配置启动监听地址
   2. 加入handlerMapping
   3. 按规范启动，配置连接超时、请求超时等
3. 启动监听
4. 错误返回

# 其它handler

1. header
2. healthz

# makefile 

Makefile 简介
一个工程中的源文件不计其数，按其类型、功能、模块分别放在若干个目录中。makefile定义了一系列的规则来指定，哪些文件需要先编译，哪些文件需要后编译，哪些文件需要重新编译，甚至进行更复杂的功能操作(因为makefile就像一个shell脚本一样，可以执行操作系统的命令)。
makefile带来的好处就是——“自动化编译”，一但写好，只需要一个make命令，整个工程完全编译，极大的提高了软件的开发效率。make是一个命令工具，是一个及时makefile中命令的工具程序。
make工具最主要也是最基本的功能就是根据makefile文件中描述的源程序至今的相互关系来完成自动编译、维护多个源文件工程。而makefile文件需要按某种语法进行编写，文件中需要说明如何编译各个源文件并链接生成可执行文件，要求定义源文件之间的依赖关系。

TODO: 详细学习 https://www.ruanyifeng.com/blog/2015/02/make.html

## build.sh

# Dockerfile

# 执行


编译并推送镜像
```cgo
make build
```

启动
```cgo
make run
```

查看容器信息
```cgo
docker ps
```

执行下面命令查看 docker 容器详情
```cgo
docker inspect 9da42d4b92ed|grep Pid
```

进入容器网络
```cgo
nsenter -t 37437 -n
```

执行网络命令
```cgo
ifconfig
```

退出容器网络
```cgo
exit
```

再执行网络命令
```cgo
ifconfig
```
