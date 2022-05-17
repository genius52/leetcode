package number

import (
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	var res int = 0
	var s string = strconv.Itoa(num)
	var l int = len(s)
	for i := 0;i + k <= l;i++{
		n,_ := strconv.Atoi(s[i:i + k])
		if n != 0 && num % n == 0{
			res++
		}
	}
	return res
}