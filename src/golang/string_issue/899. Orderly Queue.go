package string_issue

import (
	"sort"
	"strings"
)

type sortChar []string

func (s sortChar) Less(i, j int) bool {
	if s[i] < s[j]{
		return true
	}
	return false
}

func (s sortChar) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortChar) Len() int {
	return len(s)
}

func cmp_string(s1 string,s2 string,l int)bool{
	for i := 0;i < l;i++{
		if s1[i] > s2[i]{
			return true
		}else if s1[i] < s2[i]{
			return false
		}else{
			continue
		}
	}
	return false
}

func OrderlyQueue(s string, k int) string {
	if k > 1{
		var data []string = strings.Split(s,"")
		sort.Sort(sortChar(data))
		return strings.Join(data,"")
	}else{
		var l int = len(s)
		var res string = s
		for i := 1;i < l;i++{
			s2 := s[i:] + s[:i]
			if cmp_string(res,s2,l){
				res = s2
			}
		}
		return res
	}
}