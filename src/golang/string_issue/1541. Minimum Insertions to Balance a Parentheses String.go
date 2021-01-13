package string_issue

import (
	"strings"
)

//Input: s = "))())("
//Output: 3
//Explanation: Add '(' to match the first '))', Add '))' to match the last '('.
//"((()))())()()()))))"
func MinInsertions(s string) int {
	s = strings.ReplaceAll(s,"))","|")
	var l int = len(s)
	var res int = 0
	var left_cnt int = 0
	var visit int = 0
	for visit < l{
		if s[visit] == '(' {
			left_cnt++
		}else if s[visit] == '|'{
			if left_cnt > 0{
				left_cnt--
			}else{
				res++
			}
		}else if s[visit] == ')'{
			if left_cnt > 0{
				left_cnt--
				res++
			}else{
				res += 2
			}
		}
		visit++
	}
	res += left_cnt * 2
	return res
}