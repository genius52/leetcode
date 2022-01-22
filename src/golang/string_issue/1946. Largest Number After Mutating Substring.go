package string_issue

import (
	"strconv"
	"strings"
)

func maximumNumber(num string, change []int) string {
	var l int = len(num)
	var res strings.Builder
	var can_replace bool = true
	var replaced bool = false
	for i := 0;i < l;i++{
		n := num[i] - '0'
		if can_replace && change[n] > int(n){
			res.WriteString(strconv.Itoa(change[n]))
			//res += strconv.Itoa(change[n])
			replaced = true
		}else if change[n] == int(n) {
			//res += string(num[i])
			res.WriteString(string(num[i]))
		}else{
			//res += string(num[i])
			res.WriteString(string(num[i]))
			if replaced && can_replace{
				can_replace = false
			}
		}
	}
	return res.String()
}