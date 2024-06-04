package string_issue

import "strings"

func ClearStars(s string) string {
	var l int = len(s)
	var cnt [26][]int
	var deleted map[int]bool = make(map[int]bool)
	for i := 0; i < l; i++ {
		if s[i] != '*' {
			cnt[s[i]-'a'] = append(cnt[s[i]-'a'], i)
		} else {
			for j := 0; j < 26; j++ {
				if len(cnt[j]) != 0 {
					deleted[cnt[j][len(cnt[j])-1]] = true
					cnt[j] = cnt[j][:len(cnt[j])-1]
					break
				}
			}
		}
	}
	var ss strings.Builder
	for i := 0; i < l; i++ {
		if s[i] != '*' {
			if _, ok := deleted[i]; !ok {
				ss.WriteByte(s[i])
			}
		}
	}
	return ss.String()
}
