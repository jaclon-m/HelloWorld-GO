
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
TODO：不同语言的error处理方式

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

# 多线程

- 协程
```cgo
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