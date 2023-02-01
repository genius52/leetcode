package number

import "strconv"

func alternateDigitSum(n int) int {
	var sign int = 1
	var res int = 0
	var s string = strconv.Itoa(n)
	for i := 0;i < len(s);i++{
		res += int(s[i] - '0') * sign
		sign = -sign
	}
	return res
}