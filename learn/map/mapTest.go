package main

import "fmt"

func main() {
	x := map[int]string{1: "one", 2: "two"}
	fmt.Println(x)
	x[12] = "1212"
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	myFuncMap := map[string]func() int{
		"funcA": func() int { return 1 },
	}
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	fmt.Println(f())
	value, exists := myMap["a"]
	if exists {
		fmt.Println(value)
	}
	for k, v := range myMap {
		println(k, v)
	}

}
