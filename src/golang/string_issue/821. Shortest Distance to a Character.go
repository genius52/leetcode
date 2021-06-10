package string_issue

import "math"

//Input: s = "loveleetcode", c = "e"
//Output: [3,2,1,0,1,0,0,1,2,2,1,0]
func shortestToChar(s string, c byte) []int {
	var l int = len(s)
	var res []int = make([]int,l)
	var visit int = 0
	var pre_c_position int = -l
	for visit < l{
		var left int = visit
		for visit < l && s[visit] != c{
			visit++
		}
		if visit == l{
			for left < visit{
				res[left] = left - pre_c_position
				left++
			}
		}else{
			for left <= visit{
				res[left] = int(math.Min(math.Abs(float64(left - pre_c_position)),math.Abs(float64(left - visit))))
				left++
			}
		}
		pre_c_position = visit
		visit = left
	}
	return res
}