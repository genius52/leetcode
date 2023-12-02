package string_issue

func BeautifulSubstrings(s string, k int) int {
	var l int = len(s)
	var res int = 0
	var record map[int][]int = make(map[int][]int)
	record[0] = append(record[0], -1)
	var vowels int = 0
	for i := 0; i < l; i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			vowels++
		}
		not_vowels := i + 1 - vowels
		diff := vowels - not_vowels
		if indices, ok := record[diff]; ok {
			for _, idx := range indices {
				vowels := (i - idx + 1) / 2
				if vowels*vowels%k == 0 {
					res++
				}
			}
		}
		record[diff] = append(record[diff], i)
	}
	return res
}
