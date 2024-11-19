package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "key", "val")
	go func(c context.Context) {
		fmt.Println(c.Value("key"))
	}(ctx)

	timeOutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeOutCtx)

	select {
	case <-timeOutCtx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit")
	}

}
