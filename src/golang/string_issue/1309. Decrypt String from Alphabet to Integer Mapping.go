package string_issue

import "strconv"

func FreqAlphabets(s string) string {
	var l int = len(s)
	var res string
	var idx int = l - 1
	for idx >= 0{
		if s[idx] == '#'{
			n,_ := strconv.Atoi(s[idx - 2:idx])
			c := string('a' + n - 1)
			res = c + res
			idx -= 3
		}else{
			c := string(s[idx] - '1' + 'a')
			res = c + res
			idx--
		}
	}
	return res
}