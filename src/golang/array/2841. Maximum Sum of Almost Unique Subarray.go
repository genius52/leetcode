package array

func maxSum2841(nums []int, m int, k int) int64 {
	var l int = len(nums)
	var record map[int]int = make(map[int]int)
	var res int64 = 0
	var left int = 0
	var right int = 0
	var sum int64 = 0
	for right < k {
		sum += int64(nums[right])
		record[nums[right]]++
		right++
	}
	if len(record) >= m {
		res = sum
	}
	for right < l {
		record[nums[left]]--
		sum -= int64(nums[left])
		record[nums[right]]++
		sum += int64(nums[right])
		if record[nums[left]] == 0 {
			delete(record, nums[left])
		}
		if len(record) >= m && sum > res {
			res = sum
		}
		left++
		right++
	}
	return res
}
