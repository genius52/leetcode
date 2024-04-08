package string_issue

import "strings"

func GetSmallestString3106(s string, k int) string {
	var l int = len(s)
	var ss strings.Builder
	for i := 0; i < l; i++ {
		if k == 0 {
			ss.WriteByte(s[i])
		} else {
			cost := min_int(int(s[i]-'a'), int('a'+26-s[i])) //改成 a 的花费
			if k >= cost {
				ss.WriteByte('a')
				k -= cost
			} else {
				ss.WriteByte(s[i] - uint8(k))
				k = 0
			}
		}
	}
	return ss.String()
}
