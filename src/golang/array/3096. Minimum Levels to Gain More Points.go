package array

func minimumLevels(possible []int) int {
	var sum int = 0
	var l int = len(possible)
	for i := 0; i < l; i++ {
		if possible[i] == 1 {
			sum++
		} else {
			sum--
		}
	}
	var cur int = 0
	for i := 0; i < l-1; i++ {
		if possible[i] == 1 {
			cur++
		} else {
			cur--
		}
		rest := sum - cur
		if cur > rest {
			return i + 1
		}
	}
	return -1
}
