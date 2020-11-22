package string_issue

import "strings"

//Input: n = 5, k = 73
//Output: "aaszz"
//Input: n = 3, k = 27
//Output: "aay"
func GetSmallestString(n int, k int) string {
	var res []string = make([]string,n)
	var index int = n - 1
	for n > 0{
		if k - n + 1 >= 26{
			res[index] = "z"
			k -= 26
		}else{
			cur := k - n + 1
			res[index] = string('a' + cur - 1)
			k -= cur
		}
		n--
		index--
	}
	var s string = strings.Join(res,"")
	return s
}