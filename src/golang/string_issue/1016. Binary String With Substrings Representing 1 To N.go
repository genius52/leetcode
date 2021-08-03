package string_issue

import "strings"

func QueryString(S string, N int) bool {
	var l int = len(S)
	for i := N;i >= 1;i--{
		var s string
		var num int = i
		for num > 0{
			if (num | 1) == num{
				s = "1" + s
			}else{
				s = "0" + s
			}
			num >>= 1
		}
		cur_len := len(s)
		if cur_len > l{
			return false
		}
		if !strings.Contains(S,s){
			return false
		}
	}
	return true
}