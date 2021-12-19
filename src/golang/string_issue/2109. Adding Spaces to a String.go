package string_issue

import "bytes"

func addSpaces(s string, spaces []int) string {
	var l int = len(spaces)
	var res bytes.Buffer
	var pre int = 0
	for i := 0; i < l; i++ {
		res.WriteString(s[pre:spaces[i]] + " ")
		pre = spaces[i]
	}
	res.WriteString(s[pre:])
	return res.String()
}
