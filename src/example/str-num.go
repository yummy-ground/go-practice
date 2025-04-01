package example

func sliceInts(ints []int, start int, end int) []int {
	//slice := make(ints, len(ints))
	return ints[start:end]
}

func sliceStrings(strs []string, start int, end int) []string {
	//slice := make([]int, len(strs))
	return strs[start:end]
}
