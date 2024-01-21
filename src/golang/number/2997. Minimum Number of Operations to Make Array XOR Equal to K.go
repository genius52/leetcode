package number

func minOperations2997(nums []int, k int) int {
	var sum int = 0
	for _, n := range nums {
		sum ^= n
	}
	var cnt int = 0
	for i := 0; i < 32; i++ {
		bit1 := sum & (1 << i)
		bit2 := k & (1 << i)
		if bit1 != bit2 {
			cnt++
		}
	}
	return cnt
}
