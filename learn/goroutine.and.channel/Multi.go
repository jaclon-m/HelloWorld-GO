package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
练习二：改为多个生生产者和消费者
*/
func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)
	go pro1(ch1)
	go pro2(ch2)
	go multiCon(ch1, ch2)
	time.Sleep(1e9)
}

func pro1(ch chan<- int) {
	for {
		rand.Seed(time.Now().UnixNano())
		i := rand.Int()
		fmt.Printf("produce i: %d\n", i)
		ch <- i
	}
}

func pro2(ch chan<- int) {
	for {
		rand.Seed(time.Now().UnixNano())
		i := rand.Int()
		fmt.Printf("produce i: %d\n", i)
		ch <- i
	}
}

func multiCon(ch1, ch2 <-chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received from ch1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received from ch2: %d\n", v)
		}
	}
}
