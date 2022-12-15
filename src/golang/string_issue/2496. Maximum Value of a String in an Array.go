package string_issue

import "strconv"

func maximumValue(strs []string) int {
	var res int = 0
	for _,s := range strs{
		n,err := strconv.Atoi(s)
		if err == nil{
			if n > res{
				res = n
			}
		}else{
			n2 := len(s)
			if n2 > res{
				res = n2
			}
		}
	}
	return res
}