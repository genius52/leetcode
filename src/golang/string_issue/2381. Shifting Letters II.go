package string_issue

import "strings"

func ShiftingLetters3(s string, shifts [][]int) string {
	var l int = len(s)
	var move []int = make([]int,l)
	for i := 0;i < len(shifts);i++{
		var start int = shifts[i][0]
		var end int = shifts[i][1]
		var step int = 0
		if shifts[i][2] == 0{
			step = -1
		}else{
			step = 1
		}
		for j := start;j <= end;j++{
			move[j] += step
		}
	}
	var res strings.Builder
	for i := 0;i < l;i++{
		move[i] %= 26
		var offset int = (int(s[i]) - 'a' + move[i]) % 26
		if offset < 0{
			offset += 26
		}
		var c byte = 'a' + byte(offset)
		res.WriteString(string(c))
	}
	return res.String()
}