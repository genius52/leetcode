package string_issue

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	var record map[string]int = make(map[string]int)
	data := strings.Split(s1 + " " + s2," ")
	for _,d := range data{
		if _,ok := record[d];ok{
			record[d]++
		}else{
			record[d] = 1
		}
	}
	var res []string
	for r,cnt := range record{
		if cnt > 1{
			continue
		}
		res = append(res,r)
	}
	return res
}