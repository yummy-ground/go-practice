package main

import "fmt"

func main() {
	var a int
	var b int
	var longA int64
	costPerA := 0.2
	costPerB := 0.4
	const nameA string = "Donggyu"
	const nameB string = "Seunghyun"
	costPerA = 0.1
	nameConst := "Hi"

	totalCost := costPerA + costPerB

	a = 3
	b = 4
	longA = 120000

	fmt.Println(a + b)
	fmt.Println(longA)
	fmt.Println(totalCost)
	fmt.Println(nameA)
	fmt.Println(nameA + nameB)
	fmt.Println(nameConst)
}
