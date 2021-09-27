package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Hello world")
	var goos string = runtime.GOOS
	fmt.Println("The operator system is :",goos)
	path := os.Getenv("PATH")
	fmt.Println("Path is: ",path)
}
