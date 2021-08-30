package string_issue

//Input: s = "(ed(et(oc))el)"
//Output: "leetcode"
func reverse_string(s string) string {
	s1 := []rune(s)
	for i := 0; i < len(s1)/2; i++ {
		tmp := s1[i]
		s1[i] = s1[len(s1)-1-i]
		s1[len(s1)-1-i] = tmp
	}
	return string(s1)
}
func recur_reverseParentheses(s string,l int,start int)(string,int){
	var res string
	var visit int = start
	for visit < l{
		if s[visit] == '('{
			sub,pos := recur_reverseParentheses(s,l,visit + 1)
			visit = pos + 1
			res += reverse_string(sub)
		}else if s[visit] == ')'{
			break
		}else{
			res += string(s[visit])
			visit++
		}
	}
	return res,visit
}

func ReverseParentheses(s string) string {
	res,_ := recur_reverseParentheses(s,len(s),0)
	return res
}