package main

import "fmt"

func main() {
	// Multiply this function changes reply:
	n := 0
	reply := &n
	fmt.Println("Multiply:", *reply) // Multiply: 0
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50

	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

// Multiply this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}