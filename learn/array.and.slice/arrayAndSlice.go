package main

import "fmt"

func main() {
	testFor()
	testSlice()
	testSlice2()
}

func testFor() {
	var arr1 [16]int
	for i,_:= range arr1{
		arr1[i] = i
	}
	fmt.Println(arr1)
}

func testSlice() {
	var slice1 []int = make([]int ,10)
	for i,_ := range slice1{
		slice1[i] = i*2
	}
	fmt.Println(slice1)
	fmt.Println(len(slice1),"\t",cap(slice1))
}

func testSlice2()  {
	mySlice := []string{"I","am","stupid","and","weak"}
	for index,value := range mySlice{
		if value == "stupid"{
			mySlice[index] = "smart"
		}
		if value == "weak"{
			mySlice[index] = "strong"
		}
	}
	fmt.Println(mySlice)
}
