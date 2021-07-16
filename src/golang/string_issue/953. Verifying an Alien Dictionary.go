package string_issue

func my_cmp(s1 string,s2 string,dict [26]int)bool{
	var l1 int = len(s1)
	var l2 int = len(s2)
	var l int = l1
	if l2 < l1{
		l = l2
	}
	for i := 0;i < l;i++{
		c1 := dict[s1[i] - 'a']
		c2 := dict[s2[i] - 'a']
		if c1 > c2{
			return false
		}else if c1 < c2{
			return true
		}
	}
	if l1 > l2{
		return false
	}
	return true
}

func IsAlienSorted(words []string, order string) bool {
	var dict [26]int
	for i := 0;i < 26;i++{
		dict[order[i] - 'a'] = i
	}
	var l int = len(words)
	for i := 0;i < l - 1;i++{
		if !my_cmp(words[i],words[i + 1],dict){
			return false
		}
	}
	return true
}