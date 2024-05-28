package string_issue

import "strings"

func CompressedString(word string) string {
	var ss strings.Builder
	var left int = 0
	var right int = 1
	for right < len(word) {
		if word[right] != word[right-1] || (right-left) == 9 {
			ss.WriteByte(byte(right-left) + '0')
			ss.WriteByte(word[left])
			left = right
		}
		right++
	}
	ss.WriteByte(byte(right-left) + '0')
	ss.WriteByte(word[left])
	return ss.String()
}
