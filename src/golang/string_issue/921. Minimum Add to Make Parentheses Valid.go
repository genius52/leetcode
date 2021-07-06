package string_issue

//Input: s = "()))(("
//Output: 4
func minAddToMakeValid(s string) int {
	var left_cnt int = 0
	var l int = len(s)
	var res int = 0
	for i := 0;i < l;i++{
		if s[i] == '('{
			left_cnt++
		}
		if s[i] == ')'{
			left_cnt--
		}
		if left_cnt < 0{
			res++
			left_cnt = 0
		}
	}
	return res + left_cnt
}