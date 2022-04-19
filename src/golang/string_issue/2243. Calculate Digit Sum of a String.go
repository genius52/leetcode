package string_issue

import (
	"strconv"
	"strings"
)

func digitSum(s string, k int) string {
	var l int = len(s)
	for l > k{
		var cur strings.Builder
		for i := 0;i < l;i += k{
			var sum int = 0
			for j := 0;j < k && i + j < l;j++{
				n := s[i + j] - '0'
				sum += int(n)
			}
			cur.WriteString(strconv.Itoa(sum))
		}
		s = cur.String()
		l = len(s)
	}
	return s
}