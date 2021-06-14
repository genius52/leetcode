package string_issue

import "strings"

func FindReplaceString(S string, indexes []int, sources []string, targets []string) string {
	var l int = len(indexes)
	var s_len int = len(S)
	var record map[int]int = make(map[int]int)
	for i := 0;i < l;i++{
		if strings.HasPrefix(S[indexes[i]:],sources[i]){
			record[indexes[i]] = i
		}
	}
	var res string
	var visit int = 0
	for visit < s_len{
		if idx,ok := record[visit];ok{
			res += targets[idx]
			visit += len(sources[idx])
		}else{
			res += string(S[visit])
			visit++
		}
	}
	return res
}