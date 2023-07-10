package string_issue

import (
	"math"
	"strconv"
)

func dfs_minimumBeautifulSubstrings(s string, l int, idx int) int {
	if idx == l {
		return 0
	}
	if s[idx] == '0' {
		return -1
	}
	var res int = 1<<31 - 1
	for i := idx + 1; i <= l; i++ {
		n, _ := strconv.ParseInt(s[idx:i], 2, 0)
		r := int64(math.Log10(float64(n)) / math.Log10(5.0))
		if math.Pow(5.0, float64(r)) == float64(n) {
			cur := dfs_minimumBeautifulSubstrings(s, l, i)
			if cur != -1 {
				res = min_int(res, 1+cur)
			}
		}
	}
	if res == 1<<31-1 {
		return -1
	}
	return res
}

func MinimumBeautifulSubstrings(s string) int {
	var l int = len(s)
	//if s[0] == '0' {
	//	return -1
	//}
	return dfs_minimumBeautifulSubstrings(s, l, 0)
}
