package example

import (
	"fmt"
	"testing"
)

func Test_StringReadAndWrite(t *testing.T) {
	var lines []string
	if err := readLineFrom("bill\ntom\njane\nrony\ndonggyu", &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines) // [bill tom jane rony donggyu]

	/*
		bill
		tom
		jane
		rony
		donggyu
	*/
	if err := writeLineFrom(lines); err != nil {
		fmt.Println(err)
	}

}

func Test_FileReadAndWrite(t *testing.T) {
	lines := []string{
		"abc@naver.com", "abc@gmail.com", "bang2brew@gmail.com",
	}
	if err := writeFile("test.txt", lines); err != nil {
		fmt.Println(err)
	}

	var readLines []string
	if err := readFileFrom("test.txt", &readLines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines) // [abc@naver.com abc@gmail.com bang2brew@gmail.com]
}
