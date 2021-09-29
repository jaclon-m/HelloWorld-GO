package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
基于Channel编写一个简单的单线程生产者消费者模型
队列:队列长度10，队列元素类型为 int
生产者:每1秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者: 每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/
func main() {
	ch := make(chan int, 10)
	go product(ch)
	consumer(ch)
	fmt.Println("finish")

}

func product(ch chan<- int) {
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Int()
		fmt.Println("putting: ", n)
		ch <- n
		time.Sleep(1e9)
	}
}

func consumer(ch <-chan int) {
	for {
		time.Sleep(1e9)
		n := <-ch
		fmt.Println("getting: ", n)
	}
}
