package number

import (
	"strconv"
)

func IsStrictlyPalindromic(n int) bool {
	for i := 2;i <= n - 2;i++{
		var s string
		var num int = n
		for num > 0{
			mod := num % i
			num /= i
			s += strconv.Itoa(mod)
		}
		var equal bool = true
		var l int = len(s)
		for j := 0;j < l;j++{
			if s[j] != s[l - 1 - j]{
				equal = false
				break
			}
		}
		if !equal{
			return false
		}
	}
	return true
}