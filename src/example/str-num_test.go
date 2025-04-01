package example

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_sliceInts(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println(sliceInts(ints, 1, 3))
	fmt.Println(sliceInts(ints, 1, len(ints)-1))
}

func Test_sliceStrings(t *testing.T) {
	strs := []string{"가", "나", "다", "라", "미"}
	fmt.Println(sliceStrings(strs, 1, 3))
	fmt.Println(sliceStrings(strs, 1, len(strs)-1))
	appendedStrs := append(sliceStrings(strs, 1, len(strs)), []string{"바", "사"}...)
	fmt.Println(appendedStrs)
}

var (
	i   int
	k   int64
	f   float64
	s   string
	err error
)

func Test_StringToNum(t *testing.T) {
	i, err = strconv.Atoi("350")
	fmt.Println(i)
	k, err = strconv.ParseInt("100", 2, 32) // 2진수 100 -> 4
	fmt.Println(k)
	k, err = strconv.ParseInt("100", 10, 32) // 10진수 100 -> 100 (int64)
	fmt.Println(k)
	k, err = strconv.ParseInt("100", 16, 32) // 16진수 100 -> 256
	fmt.Println(k)
	k, err = strconv.ParseInt("cc7fdd", 16, 32) // 13402077
	fmt.Println(k)
	k, err = strconv.ParseInt("0xcc7fdd", 0, 32) // 13402077
	fmt.Println(k)
	f, err = strconv.ParseFloat("3.14", 64) // 3.14
	fmt.Println(f)
	s = strconv.Itoa(350)
	fmt.Println(s)
	s = strconv.FormatInt(13402077, 16) // cc7fdd
	fmt.Println(s)
}
