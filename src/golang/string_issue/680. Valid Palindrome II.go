package string_issue

//Input: s = "abca"
//Output: true
//Explanation: You could delete the character 'c'.
func dfs_validPalindrome(s string,left int,right int,allow_del bool)bool{
	for left < right{
		if s[left] == s[right]{
			left++
			right--
		}else{
			if allow_del{
				return dfs_validPalindrome(s,left + 1,right,false) || dfs_validPalindrome(s,left,right - 1,false)
			}else{
				return false
			}
		}
	}
	return true
}

func validPalindrome(s string) bool {
	var l int = len(s)
	return dfs_validPalindrome(s,0,l - 1,true)
}