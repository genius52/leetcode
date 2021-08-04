package string_issue

func removeOuterParentheses(S string) string {
	var res string
	var open int = 0
	for _,c := range S{
		if c == '(' {
			if open > 0{
				res += "("
			}
			open++
		}else if c == ')' {
			if open > 1{
				res += ")"
			}
			open--
		}
	}
	return res
}