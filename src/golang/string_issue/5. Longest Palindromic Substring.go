package string_issue

//Input: s = "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
func extend_longestPalindrome(s string,l int,pos1 int,pos2 int)string{
	for pos1 >= 0 && pos2 < l && s[pos1] == s[pos2]{
		pos1--
		pos2++
	}
	return s[pos1 + 1:pos2]
}

func longestPalindrome(s string) string{
	var l int = len(s)
	if l <= 1{
		return s
	}
	var res string
	var max_len int = 0
	for i := 0;i < l - 1;i++{
		ret := extend_longestPalindrome(s,l,i,i)
		var cur_len int = len(ret)
		if cur_len > max_len{
			res = ret
			max_len = cur_len
		}
		ret = extend_longestPalindrome(s,l,i,i + 1)
		cur_len = len(ret)
		if cur_len > max_len{
			res = ret
			max_len = cur_len
		}
	}
	return res
}