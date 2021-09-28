package _map

import "fmt"

func main() {
	x := map[int]string{1:"one",2:"two"}
	fmt.Println(x)
	x[12] = "1212"
}
