package string_issue

func RemoveAlmostEqualCharacters(word string) int {
	var res int = 0
	var pre_change bool = false
	for i := 1; i < len(word); i++ {
		diff := abs_int(int(word[i]) - int(word[i-1]))
		if !pre_change && diff <= 1 {
			pre_change = true
			res++
		} else {
			pre_change = false
		}
	}
	return res
}
