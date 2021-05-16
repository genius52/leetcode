package string_issue

import (
	"strconv"
	"strings"
)

//Input: s = "Myself2 Me1 I4 and3"
//Output: "Me Myself and I"
//Explanation: Sort the words in s to their original positions "Me1 Myself2 and3 I4", then remove the numbers.
func sortSentence(s string) string {
	var origin []string = strings.Split(s," ")
	var l int = len(origin)
	var res []string = make([]string,l)
	for i := 0;i < l;i++{
		var s_len int = len(origin[i])
		j := s_len - 1
		for j >= 0{
			if origin[i][j] < '0' || origin[i][j] > '9'{
				break
			}
			j--
		}
		var s2 string = origin[i][:j + 1]
		index,err := strconv.Atoi(origin[i][j + 1:])
		if err == nil{
			res[index - 1] = s2
		}
	}
	return strings.Join(res," ")
}