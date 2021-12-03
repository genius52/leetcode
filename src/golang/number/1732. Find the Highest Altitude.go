package number

func largestAltitude(gain []int) int {
	var height int = 0
	var res int = 0
	for _, n := range gain {
		height += n
		if height > res {
			res = height
		}
	}
	return res
}
