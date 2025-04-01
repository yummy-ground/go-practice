package example

import (
	"fmt"
	"testing"
)

func Test_Map(t *testing.T) {
	m := map[string]int{}

	m["1"] = 1
	fmt.Println(m)
	value, ok := m["1"]
	fmt.Println(value)
	fmt.Println(ok)

	_, ok = m["2"]
	fmt.Println(ok)
}
