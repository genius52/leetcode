package number

func minimumChairs(s string) int {
	var res int = 0
	var cur int = 0
	for _, c := range s {
		if c == 'E' {
			cur++
			if cur > res {
				res = cur
			}
		} else if c == 'L' {
			cur--
		}
	}
	return res
}
