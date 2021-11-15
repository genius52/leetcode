package string_issue

func maxDepth(s string) int {
	var l int = len(s)
	var cnt int = 0
	var res int = 0
	for i := 0;i < l;i++{
		if s[i] == '('{
			cnt++
			if cnt > res{
				res = cnt
			}
		}else if s[i] == ')'{
			cnt--
		}
	}
	return res
}