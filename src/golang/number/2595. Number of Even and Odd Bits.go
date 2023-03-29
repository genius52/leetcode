package number

func evenOddBit(n int) []int {
	var res []int = make([]int, 2)
	var is_even bool = true
	for n > 0 {
		if n|1 == n {
			if is_even {
				res[0]++
			} else {
				res[1]++
			}
		}
		is_even = !is_even
		n >>= 1
	}
	return res
}
