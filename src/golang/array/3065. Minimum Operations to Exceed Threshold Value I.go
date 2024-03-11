package array

func minOperations3046(nums []int, k int) int {
	var res int = 0
	for _, n := range nums {
		if n < k {
			res++
		}
	}
	return res
}
