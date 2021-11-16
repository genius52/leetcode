package string_issue

//func check_palindrome(s string) bool{
//	var left int = 0
//	var right int = len(s) - 1
//	for left < right{
//		if s[left] != s[right]{
//			return false
//		}
//		left++
//		right--
//	}
//	return true
//}

func recursive_checkPalindromeFormation(s1 string,l1 int,left int,s2 string,l2 int,right int)bool{
	if left != 0 && left >= right{
		return true
	}
	if s1[left] == s2[right]{
		return recursive_checkPalindromeFormation(s1,l1,left + 1,s2,l2,right - 1) ||
			recursive_checkPalindromeFormation(s1,l1,left + 1,s1,l1,right - 1) ||
			recursive_checkPalindromeFormation(s2,l2,left + 1,s2,l1,right - 1)
	}
	return false
}

func CheckPalindromeFormation(a string, b string) bool {
	var l int = len(a)
	if l <= 1{
		return true
	}
	return recursive_checkPalindromeFormation(a,l,0,b,l,l -1) ||
		recursive_checkPalindromeFormation(b,l,0,a,l,l -1)
}