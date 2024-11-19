package main

import "fmt"

func testDefer() {
	fmt.Println("hello ...")
	defer func() {
		fmt.Println("into defer...")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("a panic is triggered")
}
