package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("첫번째 숫자를 입력하세요.")
	line, _ := reader.ReadString('\n') // _ 을 통해 무시
	line = strings.TrimSpace(line)
	numberA, _ := strconv.Atoi(line)

	fmt.Println("두번째 숫자를 입력하세요.")
	line, _ = reader.ReadString('\n') // _ 을 통해 무시
	line = strings.TrimSpace(line)
	numberB, _ := strconv.Atoi(line)

	fmt.Printf("입력 숫자는 %d, %d 입니다.\n", numberA, numberB)

	fmt.Println("연산자 기호를 입력하세요. (+, -, *, /, %)")
	line, _ = reader.ReadString('\n') // _ 을 통해 무시
	op := strings.TrimSpace(line)

	switch op {
	case "+":
		//goto PlusMent
		fmt.Println("<덧셈>")
		fmt.Printf("결과 = %v\n", numberA+numberB)
	case "-":
		//goto MinusMent
		fmt.Println("<뺄셈>")
		fmt.Printf("결과 = %v\n", numberA-numberB)
	case "*":
		//goto MultiMent
		fmt.Println("<곱셈>")
		fmt.Printf("결과 = %v\n", numberA*numberB)
	case "/":
		//goto DevideMent
		fmt.Println("<나눗셈>")
		fmt.Printf("결과 = %v\n", numberA/numberB)
	case "%":
		//goto ModMent
		fmt.Println("<나머지>")
		fmt.Printf("결과 = %v\n", numberA%numberB)
	default:
		fmt.Println("잘못된 연산자를 입력했습니다. 프로그램을 종료합니다.")
		goto End
	}

End:
	fmt.Println("Bye")
}
