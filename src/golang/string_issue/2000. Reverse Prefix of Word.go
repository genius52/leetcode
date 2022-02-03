package string_issue

import "strings"

func reversePrefix2(word string, ch byte) string{
	var idx int = -1
	var l int = len(word)
	for i := 0;i < l;i++{
		if word[i] == ch{
			idx = i
			break
		}
	}
	if idx == -1{
		return word
	}
	var sb strings.Builder
	for i := idx;i >= 0;i--{
		sb.WriteByte(word[i])
	}
	sb.WriteString(word[idx + 1:])
	return sb.String()
}

func reversePrefix(word string, ch byte) string {
	var idx int = -1
	var l int = len(word)
	for i := 0;i < l;i++{
		if word[i] == ch{
			idx = i
			break
		}
	}
	if idx == -1{
		return word
	}
	var res string
	for i := idx;i >= 0;i--{
		res += string(word[i])
	}
	res += word[idx + 1:]
	return res
}