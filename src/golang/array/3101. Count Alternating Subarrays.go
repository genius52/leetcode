package array

func countAlternatingSubarrays(nums []int) int64 {
	var l int = len(nums)
	var res int64 = int64(l)
	var left int = 0
	var right int = 1
	for left < l && right < l {
		for right < l && nums[right] != nums[right-1] {
			res += int64(right - left)
			right++
		}
		left = right
		right++
	}
	return res
}
