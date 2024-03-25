package string_issue

import "strings"

func MinimizeStringValue(s string) string {
	var l int = len(s)
	var char_cnt [26]int
	var tag_cnt int = 0
	for i := 0; i < l; i++ {
		if s[i] != '?' {
			char_cnt[s[i]-'a']++
		} else {
			tag_cnt++
		}
	}
	var add [26]int
	for i := 0; i < l; i++ {
		if s[i] == '?' {
			var min_cnt int = 1<<31 - 1
			var min_char int = 26
			for j := 0; j < 26; j++ {
				if char_cnt[j] < min_cnt {
					min_cnt = char_cnt[j]
					min_char = j
				}
			}
			char_cnt[min_char]++
			add[min_char]++
		}
	}
	var ss strings.Builder
	for i := 0; i < l; i++ {
		if s[i] == '?' {
			for j := 0; j < 26; j++ {
				if add[j] == 0 {
					continue
				} else {
					ss.WriteByte(uint8('a' + j))
					add[j]--
					break
				}
			}
		} else {
			ss.WriteByte(s[i])
		}
	}
	return ss.String()
}
