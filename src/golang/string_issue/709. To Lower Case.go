package string_issue

import "strings"

func toLowerCase(str string) string {
	var res strings.Builder
	for _,c := range str{
		if c >= 'A' && c <= 'Z'{
			res.WriteByte(byte(c + 32))
		}else{
			res.WriteByte(byte(c))
		}
	}
	return res.String()
}