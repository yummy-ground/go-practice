package example

var (
	start = rune(44032)
	end   = rune(55204)
)

func hasHangul(s string) bool {
	for _, r := range s {
		if start <= r && r < end {
			return true
		}
	}
	return false
}

//func main() {
//	/*
//		Go : UTF-8 사용 -> 자동적으로 Unicode 사용
//	*/
//	for _, r := range "가나다라마바사" {
//		fmt.Println(r) // Unicode 출력
//	}
//	for _, r := range "가나다라마바사" {
//		fmt.Println(string(r)) // 한글 출력
//	}
//}
