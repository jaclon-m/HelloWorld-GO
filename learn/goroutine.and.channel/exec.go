package main

import (
	"fmt"
	"math/rand"
)

func test() {
	fmt.Println("start")
	queue := make(chan int, 10)
	go producer(queue)
	consumer2(queue)
}

func producer(q chan<- int) {
	for {
		if len(q) < 10 {
			i := rand.Int()
			q <- i
			fmt.Printf("produce date : %s\n", i)
		}
	}
}

func consumer2(q <-chan int) {
	for {
		if len(q) > 0 {
			fmt.Printf("consumer data : %s\n", <-q)
		}
	}
}
