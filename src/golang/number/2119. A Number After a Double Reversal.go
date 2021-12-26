package number

import "strconv"

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}
	s2 := strconv.Itoa(num)
	var l int = len(s2)
	if s2[l-1] == '0' {
		return false
	} else {
		return true
	}
}
