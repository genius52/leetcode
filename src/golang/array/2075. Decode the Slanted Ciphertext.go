package array

import "strings"

func decodeCiphertext(encodedText string, rows int) string {
	var l int = len(encodedText)
	if l == 0 || rows == 1{
		return encodedText
	}
	var res []string = make([]string,l)
	columns := l / rows
	var idx int = 0
	for i := 0;i < columns;i++{
		for j := i;j < l;j += columns + 1{
			res[idx] = string(encodedText[j])
			idx++
		}
	}
	var not_space int = l
	for i := l - 1;i >= 0;i--{
		if res[i] != "" && res[i] != " "{
			not_space = i
			break
		}
	}
	res = res[:not_space + 1]
	return strings.Join(res,"")
}