package string_issue

func findPermutationDifference(s string, t string) int {
	var l int = len(s)
	var char_idx1 [26]int
	var char_idx2 [26]int
	for i := 0; i < 26; i++ {
		char_idx1[i] = -1
		char_idx2[i] = -1
	}
	for i := 0; i < l; i++ {
		char_idx1[s[i]-'a'] = i
		char_idx2[t[i]-'a'] = i
	}
	var res int = 0
	for i := 0; i < 26; i++ {
		if char_idx1[i] != -1 {
			res += abs_int(char_idx1[i] - char_idx2[i])
		}
	}
	return res
}
