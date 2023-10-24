package array

func MaxSubarrays(nums []int) int {
	var l int = len(nums)
	var sum int = 1<<31 - 1
	for i := 0; i < l; i++ {
		sum &= nums[i]
	}
	if sum > 0 {
		return 1
	}
	var res int = 0
	var cur int = 1<<31 - 1
	for i := 0; i < l; i++ {
		if (cur & nums[i]) > sum {
			cur &= nums[i]
			continue
		}
		res++
		cur = 1<<31 - 1
	}
	return res
}
