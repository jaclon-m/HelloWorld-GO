package main

import "fmt"

type Person interface {
	GetName() string
	GetAge() int
}

type Student struct {
	Name string
	Age  int
}

func (s Student) GetName() string {
	fmt.Println("name:", s.Name)
	return s.Name
}

func (s Student) GetAge() int {
	fmt.Println("age:", s.Age)
	return s.Age
}

func testAAA() {
	var per Person
	var stu Student
	stu.Name = "zhangsan"
	stu.Age = 12
	per = stu
	per.GetName()
	per.GetAge()
}
