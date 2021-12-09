package string_issue

//Input: s = "YazaAay"
//Output: "aAa"
func LongestNiceSubstring(s string) string {
	var l int = len(s)
	if l < 2 {
		return ""
	}
	var lower [26]bool
	var upper [26]bool
	for i := 0; i < l; i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			lower[s[i]-'a'] = true
		} else {
			upper[s[i]-'A'] = true
		}
	}
	var res string
	var find_single bool = false
	for i := 0; i < l; i++ {
		var c int = 0
		if s[i] >= 'a' && s[i] <= 'z' {
			c = int(s[i] - 'a')
		} else {
			c = int(s[i] - 'A')
		}
		if lower[c] && upper[c] {
			continue
		}
		//Any string contains s[i] NEVER match
		find_single = true
		s1 := LongestNiceSubstring(s[:i])
		s2 := LongestNiceSubstring(s[i+1:])
		if len(s1) >= len(s2) {
			return s1
		} else {
			return s2
		}
	}
	if !find_single {
		return s
	}
	return res
}

//func LongestNiceSubstring(s string) string {
//	var l int = len(s)
//	var res string
//	for i := 0; i < l; i++ {
//		var upper [26]int
//		var lower [26]int
//		for j := i; j < l; j++ {
//			if s[j] >= 'a' && s[j] <= 'z' {
//				lower[s[j]-'a']++
//			} else if s[j] >= 'A' && s[j] <= 'Z' {
//				upper[s[j]-'A']++
//			}
//			var match bool = true
//			for k := 0; k < 26; k++ {
//				if (lower[k] > 0 && upper[k] == 0) || (lower[k] == 0 && upper[k] > 0) {
//					match = false
//					break
//				}
//			}
//			if match {
//				if j-i+1 > len(res) {
//					res = s[i : j+1]
//				}
//			}
//		}
//	}
//	return res
//}
