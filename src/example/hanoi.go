package main

import (
	"fmt"
)

func move(num int, from int, to int) {
	fmt.Println("원반 ", num, " : ", from, "번 기둥 -> ", to, "번 기둥")
}
func hanoi(n int, start int, end int, mid int) {
	if n == 1 {
		move(n, start, end)
	} else {
		hanoi(n-1, start, mid, end)
		move(n, start, end)
		hanoi(n-1, mid, end, start)
	}
}

func main() {
	hanoi(3, 1, 3, 2)
}
