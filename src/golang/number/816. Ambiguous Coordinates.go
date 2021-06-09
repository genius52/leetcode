package number

//Input: s = "(00011)"
//Output: ["(0, 0.011)","(0.001, 1)"]
//func dp_ambiguousCoordinates(S string,pos int,res []string){
//
//}

func check_ambiguousCoordinates(s string)([]string){
	var res []string
	var l int = len(s)
	if l == 1{
		res = append(res,s)
		return res
	}
	var all_zero bool = true
	for i := 0;i < l;i++{
		if s[i] != '0'{
			all_zero = false
			break
		}
	}
	if all_zero{
		return res
	}
	if s[0] != '0'{
		res = append(res,s)
	}

	for i := 1;i < l;i++{
		head := s[:i]//长度超过1时，不能以0开头
		if i > 1{
			if head[0] == '0'{
				break
			}
		}
		tail := s[i:]//不能以0结尾
		if tail[len(tail) - 1] == '0'{
			break
		}
		res = append(res,head + "." + tail)
	}
	return res
}

func ambiguousCoordinates(S string) []string {
	var l int = len(S)
	S = S[1:l - 1]
	l -= 2
	var res []string
	for i := 1;i < l;i++{
		s1 := S[:i]
		s2 := S[i:]
		val1 := check_ambiguousCoordinates(s1)
		val2 := check_ambiguousCoordinates(s2)
		for _,v1 := range val1{
			for _,v2 := range val2{
				res = append(res,"(" + v1 + ", " + v2 + ")")
			}
		}
	}
	return res
}