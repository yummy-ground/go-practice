package main

import "fmt"

func main() {
	a := 4
	b := 2
	if a == b {
		fmt.Println("a == b")
	}
	a /= b
	if a == b {
		fmt.Println("a/b == b")
	}

	if a == 0 {
		fmt.Println("a == 0")
	} else if a == 1 {
		fmt.Println("a == 1")
	} else {
		fmt.Println("a != 0 && a != 1")
	}

}
