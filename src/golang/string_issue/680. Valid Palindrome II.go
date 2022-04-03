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

func validPalindrome2(s string) bool{
	var l int = len(s)
	var left int = 0
	var right int = l - 1
	for left < right{
		if s[left] != s[right]{
			break
		}
		left++
		right--
	}
	if left >= right{
		return true
	}
	var left2 int = left + 1
	var right2 int = right
	for left2 < right2{
		if s[left2] != s[right2]{
			break
		}
		left2++
		right2--
	}
	if left2 >= right2{
		return true
	}
	var left3 int = left
	var right3 int = right - 1
	for left3 < right3{
		if s[left3] != s[right3]{
			break
		}
		left3++
		right3--
	}
	if left3 >= right3{
		return true
	}
	return false
}