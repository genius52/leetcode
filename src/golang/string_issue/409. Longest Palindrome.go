package string_issue

//Input: s = "abccccdd"
//Output: 7
//Explanation:
//One longest palindrome that can be built is "dccaccd", whose length is 7.
func longestPalindrome409(s string) int {
	var record [255]int
	var l int = len(s)
	for i := 0;i < l;i++{
		record[s[i]]++
	}
	var total int = 0
	var has_odd bool = false
	for i := 0;i < 255;i++{
		if (record[i] | 1) == record[i]{
			has_odd = true
		}
		total += record[i] & ^1
	}
	if has_odd{
		total++
	}
	return total
}