package number

func furthestDistanceFromOrigin(moves string) int {
	var cnt int = 0
	var dis int = 0
	for _, c := range moves {
		if c == 'L' {
			dis++
		} else if c == 'R' {
			dis--
		} else {
			cnt++
		}
	}
	return abs_int(dis) + cnt
}
