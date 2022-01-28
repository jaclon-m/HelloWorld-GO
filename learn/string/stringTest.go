package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"unicode/utf8"
)

type Human struct {
	name string
	age  int
}

func main() {
	strLen()
	strLen2()
	testStringConv()
}

/*func testJson() {
	humanStr := "{ \"age\":\"23\"}"
	var obj interface{}
	err := json.Unmarshal([]byte(humanStr),&obj)
	objMap,ok := obj.(map[string]interface{})
	for k,v := range objMap{
		switch value := v.(type) {
		case string:
			fmt.Printf("type of %s is string,value is %v\n", k, value)
		case interface{}:
			fmt.Printf("type of %s is interface{},value is %v\n", k, value)
		default:
			fmt.Println("wrong")
		}
	}
}*/
/**
Unmarshal:从string转换至struct
*/
func unmarshal2Struct(humanStr string) Human {
	h := Human{}
	err := json.Unmarshal([]byte(humanStr), &h)
	if err != nil {
		println(err)
	}
	return h
}

/**
Marshal:从struct转换至string
*/
func struct2Unmarshal(h Human) string {
	h.age = 30
	undatedBytes, err := json.Marshal(&h)
	if err != nil {
		println(err)
	}
	return string(undatedBytes)
}

func testStringConv() {
	var orig string = "ABC"
	// var an int
	var newS string
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// anInt, err = strconv.Atoi(origStr)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
func strLen() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
}

func strLen2() {
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("str1 byte is %d\n", len(str1))
	fmt.Printf("str1 character is %d\n", utf8.RuneCountInString(str1))
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("str2 byte is %d\n", len(str2))
	fmt.Printf("str2 character is %d\n", utf8.RuneCountInString(str2))
}
