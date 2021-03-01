package string_issue

//Input: s = "aacecaaa"
//Output: "aaacecaaa"
//Example 2:
//
//Input: s = "abcd"
//Output: "dcbabcd"
func ShortestPalindrome(s string) string {
	var l int = len(s)
	if l <= 1{
		return s
	}
	var end int = l - 1
	for end >= 0{
		var sub string = s[:end + 1]
		if isPalindrome(sub){
			break
		}
		end--
	}
	var res string
	for i := l - 1;i > end;i--{
		res += string(s[i])
	}
	return res + s
}