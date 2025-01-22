package main

// package가 main이 아니면
// "package command-line-arguments is not a main package" / "package go-practice is not a main package" 에러 발생
// https://forum.golangbridge.org/t/error-package-command-line-arguments-is-not-a-main-package/25851

// build 시, "hello"이라는 이름의 파일이 있으면 "go: build output "hello" already exists and is a directory" 이라는 에러 발생

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}
