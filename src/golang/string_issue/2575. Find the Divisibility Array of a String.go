package string_issue

func DivisibilityArray(word string, m int) []int {
	var l int = len(word)
	var res []int = make([]int, l)
	var premod int = 0
	for i := 0; i < l; i++ {
		val := int(word[i] - '0')
		cur := (premod*10 + val) % m
		if cur == 0 {
			res[i] = 1
		}
		premod = cur
	}
	return res
}
