package string_issue

import (
	"sort"
	"strings"
)

func sortVowels(s string) string {
	var l int = len(s)
	var letters []uint8
	var res []uint8 = make([]uint8, l)
	for i := 0; i < l; i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			letters = append(letters, s[i])
		} else {
			res[i] = s[i]
		}
	}
	var l2 int = len(letters)
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	var j int = 0
	for i := 0; i < l2; i++ {
		for j < l && res[j] != 0 {
			j++
		}
		res[j] = letters[i]
	}
	var ss strings.Builder
	for i := 0; i < l; i++ {
		ss.WriteByte(res[i])
	}
	return ss.String()
}
