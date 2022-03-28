package string_issue

import "strings"

func reverseWords(s string) string {
	var l int = len(s)
	var left int = 0
	var sb strings.Builder
	for left < l{
		var visit int = left
		for visit != l && s[visit] != ' '{
			visit++
		}
		var right int = visit - 1
		for i := right;i >= left;i--{
			sb.WriteByte(s[i])
		}
		if visit != l{
			sb.WriteString(" ")
		}
		left = visit + 1
	}
	return sb.String()
}