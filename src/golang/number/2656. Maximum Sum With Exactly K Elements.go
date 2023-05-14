package number

// 5 + 6 + 7
func maximizeSum(nums []int, k int) int {
	var max_val int = 0
	for _, n := range nums {
		if n > max_val {
			max_val = n
		}
	}
	return (max_val + max_val + k - 1) * k / 2
}
