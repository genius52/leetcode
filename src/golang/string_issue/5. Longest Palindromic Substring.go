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


//5 Longest Palindromic Substring
//Input: "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
func longestPalindrome2(s string) string {
	l := len(s)
	if l == 1{
		return s
	}
	var dp [][]bool = make([][]bool,l)
	for i := 0;i < l;i++{
		dp[i] = make([]bool,l)
		dp[i][i] = true
	}
	var res string = s[0:1]
	var max int = 0
	for i := 1;i < l;i++{
		for j := i - 1 ;j >= 0;j--{
			if s[i] == s[j] {
				if (i - j) == 1{
					dp[j][i] = true
					if (i - j + 1) > max {
						max = i - j + 1
						res = s[j : i+1]
					}
				}else{
					if dp[j + 1][i - 1] {
						dp[j][i] = true
						if (i - j + 1) > max{
							max = i - j + 1
							res = s[j:i+1]
						}
					}
				}
			}
		}
	}
	return res
}