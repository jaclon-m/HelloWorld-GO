
# JSON编解码

Unmarshal:从string转换至struct
```cgo
func unmarshal2Struct(humanStr string) Human{
	h := Human{};
	err := json.Unmarshal([]byte(humanStr),&h)
	if err != nil {
		println(err)
	}
	return h
}
```
Marshal:从struct转换至string
```cgo
func struct2Unmarshal(h Human) string{
	h.age = 30
	undatedBytes,err := json.Marshal(&h)
	if err != nil {
		println(err)
	}
	return string(undatedBytes)
}
```

# 错误处理

error本身是一个接口

# defer panic recover

```go
func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		// go func(i int) {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("loopFunc:", i)
		// }(i)
	}
}
```
• panic:可在系统出现不可恢复错误时主动调用panic,panic会使当前线程直接crash 
• defer:保证执行并把控制权交还给接收到panic的函数调用者
• recover:函数从panic或错误场景中恢复
```go
defer func() {
    fmt.Println("defer func is called")
    if err := recover(); err != nil {
    fmt.Println(err)
    }
    }()
    panic("a panic is triggered")
```

defer、panic和recover的关系
defer延迟函数在panic发生时仍会被执行，这使得recover只能在defer函数中有效。

## 错误处理的最佳实践

- 优先使用错误返回值：Go语言提倡使用错误作为函数的返回值，而非异常机制。
- 避免滥用panic：panic只应用于不可恢复的错误，如程序的内部逻辑错误。
- 提供有用的错误信息：错误信息应尽可能清晰，包含足够的上下文。
- 使用错误包装：利用错误包装机制，保留原始错误信息，构建错误链。
- 检查错误类型：使用errors.Is和errors.As来判断和提取特定的错误信息。

# 多线程

- 协程
```cgo
go functionName()
for i := 0; i < 10; i++ {
go fmt.Println(i) }
time.Sleep(time.Second)
```

- channel
```go
ch := make(chan int) 
go func() {
    fmt.Println("hello from goroutine")
    ch <- 0 //数据写入Channel 
}() 
i := <-ch//从Channel中取数据并赋值
```
- 遍历通道缓冲区
```go
ch := make(chan int, 10) 
go func() {
    for i := 0; i < 10; i++ { 
		rand.Seed(time.Now().UnixNano())
        n := rand.Intn(10) // n will be between 0 and 10 fmt.Println("putting: ", n)
        ch <- n
    }
    close(ch) 
}()
fmt.Println("hello from main") 

for v := range ch {
    fmt.Println("receiving: ", v) 
}
```
- 单向通道
只发送通道
• var sendOnly chan<- int
只接收通道
 • var readOnly <-chan int
- 关闭通道
通道无需每次关闭，只有发送方需要关闭
```go
ch := make(chan int)
defer close(ch)
if v, notClosed := <-ch; notClosed {
    fmt.Println(v) 
}
```
select轮询通道

定时器Timer
```go
timer := time.NewTimer(time.Second) 
select {
    // check normal channel
    case <-ch:
    fmt.Println("received from ch")
    case <-timer.C:
    fmt.Println("timeout waiting from channel ch")
}
```

## 上下文Context

context.Background
• Background 通常被用于主函数、初始化以及测试中，作为一个顶层的 context，也就是说一般 我们创建的 context 都是基于 Background
context.TODO
• TODO 是在不确定使用什么 context 的时候才会使用
context.WithDeadline • 超时时间
context.WithValue
• 向 context 添加键值对
context.WithCancel
• 创建一个可取消的 context

### 如何停止子协程

```go
func main() {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)
	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
```

```go
func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		fmt.Println(c.Value("a"))
	}(ctx)
	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)
	select {
		case <-timeoutCtx.Done():
			time.Sleep(1 * time.Second)
			fmt.Println("main process exit!")
	}
	// time.Sleep(time.Second * 5)
}
```

## context

https://blog.csdn.net/qq_38428433/article/details/140649778

取消信号传递：可以用来传递取消信号，让一个正在执行的函数知道它应该提前终止。
超时控制：可以设定一个超时时间，自动取消超过执行时间的操作。
截止时间：与超时类似，但是是设定一个绝对时间点，而不是时间段。
值传递：可以安全地在请求的上下文中传递数据，避免了使用全局变量或者参数列表不断增长。

用法
context.Background
• Background 通常被用于主函数、初始化以及测试中，作为一个顶层的 context，也就是说一般 我们创建的 context 都是基于 Background
context.TODO
• TODO 是在不确定使用什么 context 的时候才会使用
context.WithDeadline • 超时时间
context.WithValue
• 向 context 添加键值对
context.WithCancel
• 创建一个可取消的 context