# 包引用和依赖管理

- 早期所有依赖包都要放到gopath下，多项目会出现版本冲突
- vendor
  - 每个项目中新建vendor目录，放依赖
  - 改vendor后版本升级后需要murge
- gomod
  - 从1.6版本开始，替换掉gopkg，基本统一
  - 切换mod开启模式：切换 mod 开启模式:export GO111MODULE=on/off/auto
  
用依赖管理工具的目的？
- 防止篡改（gomod，好的方式是向社区提murge）
- 版本管理

gomod使用
- gomod 会帮助将依赖从github下载下来，并记录在gomod文件中
- gomod 默认放在pkg目录下，因此拉取别人的项目后都需要通过go mod init ,gomod tidy命令重新拉取依赖
- 替代方法： 运行 go mod vendor命令，经依赖放到vendor目录下
- 自定义版本：在require中指定版本；
- replace: 场景一：替换版本；场景二：自定义包的位置
- 如果没有网络的情况： go mod vendor并禁用go mod

![yldWRLbIPSgJp43](https://i.loli.net/2021/09/25/yldWRLbIPSgJp43.png)

GOPROXY 和 GOPRIVATE
某些私有代码仓库是goproxy.cn无法连接的，因此需要设置GOPRIVATE来声明私有代码仓库
```go
GOPRIVATE=*.corp.example.com
GOPROXY=proxy.example.com
GONOPROXY=myrepo.corp.example.com
```

# Makefile

Makefile 简介
一个工程中的源文件不计其数，按其类型、功能、模块分别放在若干个目录中。makefile定义了一系列的规则来指定，哪些文件需要先编译，哪些文件需要后编译，哪些文件需要重新编译，甚至进行更复杂的功能操作(因为makefile就像一个shell脚本一样，可以执行操作系统的命令)。
makefile带来的好处就是——“自动化编译”，一但写好，只需要一个make命令，整个工程完全编译，极大的提高了软件的开发效率。make是一个命令工具，是一个及时makefile中命令的工具程序。
make工具最主要也是最基本的功能就是根据makefile文件中描述的源程序至今的相互关系来完成自动编译、维护多个源文件工程。而makefile文件需要按某种语法进行编写，文件中需要说明如何编译各个源文件并链接生成可执行文件，要求定义源文件之间的依赖关系。
详细学习 https://www.ruanyifeng.com/blog/2015/02/make.html

# http server


Go语言将协程与fd资源绑定
• 一个 socket fd 与一个协程绑定
• 当 socket fd 未就绪时，将对应协程设置为 Gwaiting 状态，将 CPU 时间片让给其他协程
• Go 语言 runtime 调度器进行调度唤醒协程时，检查 fd 是否就绪，如果就绪则将协程置为 Grunnable 并加入执行队列 
• 协程被调度后处理 fd 数据

# 调试

gdb - dlv
glog - klog


