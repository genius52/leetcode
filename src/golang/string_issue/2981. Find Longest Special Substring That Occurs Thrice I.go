package string_issue

func MaximumLength(s string) int {
	var l int = len(s)
	for sub_len := l - 2; sub_len > 0; sub_len-- {
		var record map[string]int = make(map[string]int)
		for left := 0; left+sub_len <= l; left++ {
			var equal bool = true
			for right := left + 1; right < left+sub_len; right++ {
				if s[right] != s[left] {
					equal = false
					break
				}
			}
			if !equal {
				continue
			}
			record[s[left:left+sub_len]]++
			if record[s[left:left+sub_len]] == 3 {
				return sub_len
			}
		}
	}
	return -1
}
