package number

func maximumEnergy(energy []int, k int) int {
	var l int = len(energy)
	var res int = -(1 << 31)
	for right := l - 1; right >= l-k && right >= 0; right-- {
		var sum int = 0
		for left := right; left >= 0; left -= k {
			sum += energy[left]
			if sum > res {
				res = sum
			}
		}
	}
	return res
}
