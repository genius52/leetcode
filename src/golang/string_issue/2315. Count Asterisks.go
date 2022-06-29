package string_issue

func countAsterisks(s string) int {
	var l int = len(s)
	var find bool = false
	var res int = 0
	for i := 0;i < l;i++{
		if s[i] == '*'{
			if !find{
				res++
			}
		}else if s[i] == '|'{
			find = !find
		}
	}
	return res
}