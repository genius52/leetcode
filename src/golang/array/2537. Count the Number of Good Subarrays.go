package array

//nums = [3,1,4,3,2,2,4], k = 2
func CountGood(nums []int, k int) int64 {
	var l int = len(nums)
	var res int64 = 0
	var cnt map[int]int = make(map[int]int)
	var pairs int = 0
	var left int = 0
	for right := 0;right < l;right++{
		cnt[nums[right]]++
		if cnt[nums[right]] > 1{
			pairs += cnt[nums[right]] - 1
		}
		for left < right && pairs >= k{
			res += int64(l - right)
			cnt[nums[left]]--
			pairs -= cnt[nums[left]]
			left++
		}
	}
	return res
}