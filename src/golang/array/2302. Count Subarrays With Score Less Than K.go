package array

func CountSubarrays(nums []int, k int64) int64 {
	var l int = len(nums)
	var prefix []int64 = make([]int64,l + 1)
	for i := 0;i < l;i++{
		prefix[i + 1] = prefix[i] + int64(nums[i])
	}
	var res int64 = 0
	var left int = 0
	var right int = 0
	for left < l{
		for right <= l && (prefix[right] - prefix[left]) * int64(right - left) < k{
			right++
		}
		res += int64(right - left - 1)
		left++
	}
	return res
}