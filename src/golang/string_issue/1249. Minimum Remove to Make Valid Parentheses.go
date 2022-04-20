package string_issue

import "strings"

//func MinRemoveToMakeValid(s string) string{
//	var res string
//	var left int = 0
//	var l int = len(s)
//	for i := 0;i < l;i++{
//		if s[i] == '('{
//			left++
//		}else if s[i] == ')'{
//			left--
//		}
//		if left < 0{
//			left = 0
//		}else{
//			res += string(s[i])
//		}
//	}
//	if left == 0{
//		return res
//	}
//	var res2 string
//	for i := len(res) - 1;i >= 0;i--{
//		if res[i] == '(' && left > 0{
//			left--
//		}else{
//			res2 = string(res[i]) + res2
//		}
//	}
//	return res2
//}

// "())()((("
func reverse(s string) string {
	s1 := []rune(s)
	for i := 0; i < len(s1)/2; i++ {
		tmp := s1[i]
		s1[i] = s1[len(s1)-1-i]
		s1[len(s1)-1-i] = tmp
	}
	return string(s1)
}

func MinRemoveToMakeValid(s string) string {
	var res strings.Builder
	var start_tag int = 0
	for i := 0; i < len(s);i++{
		if s[i] == '('{
			start_tag++
		}else if s[i] == ')'{
			if start_tag <= 0{
				continue
			}else{
				start_tag--
			}
		}
		res.WriteString(string(s[i]))
	}
	var res2 string
	s2 := res.String()
	i := len(s2) - 1
	for ; i >= 0 && start_tag > 0;i--{
		if s2[i] == '('{
			start_tag--
			continue
		}
		res2 += string(s2[i])
	}
	res3 := s2[0:i+1] + reverse(res2)
	return res3
}