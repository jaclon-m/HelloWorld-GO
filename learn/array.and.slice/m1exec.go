package main

import "fmt"

func main() {
	var arr1 = []string{"I", "am", "stupid", "and", "weak"}
	var arr2 = []string{"I", "am", "smart", "and", "strong"}

	for i := range arr1 {
		arr1[i] = arr2[i]
	}
	fmt.Println(arr1)
}
