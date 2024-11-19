package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("start")
	ch := make(chan int, 10)
	go func(producer chan<- int) {
		for {
			rand.Seed(time.Now().UnixNano())
			n := rand.Int()
			fmt.Println("putting:", n)
			producer <- n
			time.Sleep(1e9)
		}

	}(ch)
	consumer(ch)
	fmt.Println("finish")
}

func consumer(con <-chan int) {
	for {
		time.Sleep(1e9)
		n := <-con
		fmt.Println("getting:", n)
	}

}
