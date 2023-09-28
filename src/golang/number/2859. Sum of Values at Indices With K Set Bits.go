package number

func sumIndicesWithKSetBits(nums []int, k int) int {
	var sum int = 0
	for idx, n := range nums {
		var one_bit int = 0
		for idx > 0 {
			if (idx | 1) == idx {
				one_bit++
				if one_bit > k {
					break
				}
			}
			idx >>= 1
		}
		if one_bit == k {
			sum += n
		}
	}
	return sum
}
