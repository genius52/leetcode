package string_issue

import "strings"

//Input: "  hello world!  "
//Output: "world! hello"
func ReverseWords(s string) string {
	array := strings.Fields(s)
	l := len(array)
	var res string
	for i := l - 1;i >= 0;i--{
		if len(res) != 0{
			res += " "
		}
		res += array[i]
	}
	return res
}