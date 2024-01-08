package number

func hasTrailingZeros(nums []int) bool {
	var find bool = false
	for _, n := range nums {
		if n|1 != n {
			if find {
				return true
			} else {
				find = true
			}
		}
	}
	return false
}
