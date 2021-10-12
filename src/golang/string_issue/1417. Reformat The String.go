package string_issue

//Input: s = "a0b1c2"
//Output: "0a1b2c"
func reformat(s string) string {
	var dig string
	var char string
	for _,c := range s{
		if c >= '0' && c <= '9'{
			dig += string(c)
		}else{
			char += string(c)
		}
	}
	var l1 int = len(dig)
	var l2 int = len(char)
	if (l1 - l2) > 1 || (l2 - l1) > 1{
		return ""
	}
	var res string
	var i,j int = 0,0
	for i < l1 && j < l2{
		res += string(dig[i]) + string(char[i])
		i++
		j++
	}
	if l1 > l2{
		res += string(dig[l1 - 1])
	}
	if l1 < l2{
		res = string(char[l2 - 1]) + res
	}
	return res
}