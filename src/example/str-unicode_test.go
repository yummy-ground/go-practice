package example

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	testHasHangul()
}

func testHasHangul() {
	fmt.Println(hasHangul("Let's study Golang!"))
	fmt.Println(hasHangul("Let's study Go랭!"))
	fmt.Println(hasHangul("Golang을 공부하자!"))
}
