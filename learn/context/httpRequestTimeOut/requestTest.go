package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func randomSleepAtMost2s() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	// 生成 0 到 2000 之间的随机整数（毫秒）
	randomMillis := rand.Intn(2000)
	// 转换为 time.Duration 类型，并乘以 time.Millisecond
	sleepDuration := time.Duration(randomMillis) * time.Millisecond
	// 随机睡眠
	fmt.Println("sleeping for", sleepDuration)
	time.Sleep(sleepDuration)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 超时时间1s
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	// 创建一个模拟正常处理完成的通道
	done := make(chan struct{})

	// 模拟异步处理逻辑
	go func() {
		// 模拟耗时操作
		// - 当随机睡眠超过1s时，会触发 ctx.Done()，取消请求
		// - 当随机睡眠不超过1s时，则会正常处理请求
		randomSleepAtMost2s()
		fmt.Println("request processed")
		close(done) // 处理完成，关闭通道
	}()

	// 模拟耗时操作
	select {
	case <-done:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("request processed successfully"))
	case <-ctx.Done():
		fmt.Println("request cancelled")
		http.Error(w, "request cancelled", http.StatusRequestTimeout)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
