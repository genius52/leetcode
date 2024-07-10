package string_issue

import "strings"

func getEncryptedString(s string, k int) string {
	var l int = len(s)
	k = k % l
	var ss strings.Builder
	for i := 0; i < l; i++ {
		ss.WriteByte(s[(i+k)%l])
	}
	return ss.String()
}
