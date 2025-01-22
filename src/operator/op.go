package main

import "fmt"

func main() {
	a := 4 // 0100
	b := 2 // 0010

	// %v : value || variable
	fmt.Printf("a + b (%d + %d) = %v\n", a, b, a+b) // 6
	fmt.Printf("a - b (%d - %d) = %v\n", a, b, a-b) // 2
	fmt.Printf("a + b (%d % %d) = %v\n", a, b, a%b) // 0
	fmt.Printf("a - b (%d - %d) = %v\n", a, b, a/b) // 2

	/*
		Bit Operation
	*/
	// AND -> 0&0 = 0 / 0&1 = 0 / 1&0 = 0 / 1&1 = 1
	// a = 0100
	// b = 0010
	//     0000
	fmt.Printf("a & b (%d & %d) = %v\n", a, b, a&b) // 0

	// OR -> 0|0 = 0 / 0|1 = 1 / 1|0 = 1 / 1|1 = 1
	// a = 0100
	// b = 0010
	//     0110
	fmt.Printf("a | b (%d | %d) = %v\n", a, b, a|b) // 6

	// XOR -> 0^0 = 0 / 0^1 = 1 / 1^0 = 1 / 1^1 = 0
	// a = 0100
	// b = 0010
	//     0110
	fmt.Printf("a ^ b (%d ^ %d) = %v\n", a, b, a^b) // 6

	// NOT XOR NOT
	// a = 1011
	// b = 1101
	//     0110
	fmt.Printf("^a ^ ^b ((^%d) ^ (^%d) => (%d) ^ (%d)) = %v\n", a, b, ^a, ^b, ^a^^b) // 6
}
