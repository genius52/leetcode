package string_issue

//s = abaabab
func make_next(s string)[]int{
	var l int = len(s)
	var next []int = make([]int,l)
	next[0] = 0
	var last_prefixlen int = 0
	for i := 1;i < l;i++{
		for last_prefixlen > 0 && s[i] != s[last_prefixlen]{
			last_prefixlen = next[last_prefixlen - 1]
		}
		if s[i] == s[last_prefixlen]{
			last_prefixlen++
		}
		next[i] = last_prefixlen
	}
	return next
}

func StrStr(haystack string, needle string) int {
	var l1 int = len(haystack)
	var l2 int = len(needle)
	if l2 == 0{
		return 0
	}
	if l1 == 0{
		return -1
	}
	if l1 < l2{
		return -1
	}
	var next []int = make_next(needle)
	var i int = 0
	var j int = 0
	for i < l1{
		for j > 0 && haystack[i] != needle[j]{
			j = next[j - 1]
		}
		if haystack[i] == needle[j]{
			j++
		}
		if j == l2{
			return i - j + 1
		}
		i++
	}
	return -1
}