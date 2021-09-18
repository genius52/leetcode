package string_issue

//输入：palindrome = "abccba"
//输出："aaccba"
//解释：存在多种方法可以使 "abccba" 不是回文，例如 "zbccba", "aaccba", 和 "abacba" 。
//在所有方法中，"aaccba" 的字典序最小。
//aaaa
func check_palindrome(s string) bool{
	var left int = 0
	var right int = len(s) - 1
	for left < right{
		if s[left] != s[right]{
			return false
		}
		left++
		right--
	}
	return true
}

func BreakPalindrome(palindrome string) string {
	var l int = len(palindrome)
	if l == 1{
		return ""
	}
	for i := 0;i < l;i++{
		for j := 0;j < 26;j++{
			s := palindrome[:i] + string('a' + j) + palindrome[i + 1:]
			if s == palindrome {
				break
			}
			if !check_palindrome(s) {
				return s
			}
		}
	}
	res := palindrome[:l - 1] + "b"
	return res
}