package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	a := 10
	//if
	if a < 10 {
		fmt.Printf("a<10")
	} else if a > 10 {
		fmt.Printf("a>10")
	}
	if v := a + 10; v > 10 {
		fmt.Printf("vvvv")
	}
	//for
	sum := 0
	for sum < 1000 {
		sum += sum
	}
}
