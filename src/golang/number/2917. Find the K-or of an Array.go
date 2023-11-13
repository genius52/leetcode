package number

func findKOr(nums []int, k int) int {
	var l int = len(nums)
	var bit_cnt [32]int
	for i := 0; i < l; i++ {
		n := nums[i]
		var offset int = 0
		for n > 0 {
			if n|1 == n {
				bit_cnt[offset]++
			}
			n >>= 1
			offset++
		}
	}
	var res int = 0
	for i := 0; i < 32; i++ {
		if bit_cnt[i] >= k {
			res |= 1 << i
		}
	}
	return res
}
