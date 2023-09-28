package number

import "strings"

func maximumOddBinaryNumber(s string) string {
	var l int = len(s)
	var one_cnt int = 0
	for i := 0; i < l; i++ {
		if s[i] == '1' {
			one_cnt++
		}
	}
	var ss strings.Builder
	for i := 0; i < l-1; i++ {
		if one_cnt > 1 {
			ss.WriteString("1")
			one_cnt--
		} else {
			ss.WriteString("0")
		}
	}
	ss.WriteString("1")
	return ss.String()
}
