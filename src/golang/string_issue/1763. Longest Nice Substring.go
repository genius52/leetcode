package string_issue

//Input: s = "YazaAay"
//Output: "aAa"
func LongestNiceSubstring(s string) string {
	var l int = len(s)
	var res string
	for i := 0;i < l;i++{
		var upper [26]int
		var lower [26]int
		for j := i;j < l;j++ {
			if s[j] >= 'a' && s[j] <= 'z' {
				lower[s[j]-'a']++
			} else if s[j] >= 'A' && s[j] <= 'Z' {
				upper[s[j]-'A']++
			}
			var match bool = true
			for k := 0; k < 26; k++ {
				if (lower[k] > 0 && upper[k] == 0) || (lower[k] == 0 && upper[k] > 0) {
					match = false
					break
				}
			}
			if match {
				if j-i+1 > len(res) {
					res = s[i:j + 1]
				}
			}
		}
	}
	return res
}