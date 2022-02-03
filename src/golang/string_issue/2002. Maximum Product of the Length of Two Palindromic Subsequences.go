package string_issue

func count(s1 string,s2 string)int{
	var l1 int = len(s1)
	var l2 int = len(s2)
	for i := 0;i < l1 /2;i++{
		if s1[i] != s1[l1 - 1 - i]{
			return 0
		}
	}
	for i := 0;i < l2 /2;i++{
		if s2[i] != s2[l2 - 1 - i]{
			return 0
		}
	}
	return l1 * l2
}

func dp_maxProduct2002(s string,l int,idx int,s1 string,s2 string)int{
	if idx >= l{
		return count(s1,s2)
	}
	res1 := dp_maxProduct2002(s,l,idx + 1,s1 + string(s[idx]),s2)
	res2 := dp_maxProduct2002(s,l,idx + 1,s1,s2 + string(s[idx]))
	res3 := dp_maxProduct2002(s,l,idx + 1,s1,s2)
	return max_int_number(res1,res2,res3)
}

func maxProduct2002(s string) int {
	return dp_maxProduct2002(s,len(s),0,"","")
}