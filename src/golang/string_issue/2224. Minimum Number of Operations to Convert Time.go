package string_issue

import "strconv"

func convertTime(current string, correct string) int {
	h1,_ := strconv.Atoi(current[:2])
	m1,_ := strconv.Atoi(current[3:])
	start := h1 * 60 + m1
	h2,_ := strconv.Atoi(correct[:2])
	m2,_ := strconv.Atoi(correct[3:])
	end := h2 * 60 + m2
	diff := end - start
	var res int = 0
	res += diff/ 60
	diff %= 60
	res += diff/15
	diff %= 15
	res += diff / 5
	diff %= 5
	res += diff
	return res
}