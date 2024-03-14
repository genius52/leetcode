package number

import "sort"

func minimumBoxes3074(apple []int, capacity []int) int {
	var sum int = 0
	for _, a := range apple {
		sum += a
	}
	sort.Ints(capacity)
	var cnt int = 0
	for i := len(capacity) - 1; i >= 0; i-- {
		sum -= capacity[i]
		cnt++
		if sum <= 0 {
			return cnt
		}
	}
	return cnt
}
