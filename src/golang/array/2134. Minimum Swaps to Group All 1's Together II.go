package array

func MinSwaps(nums []int) int {
	var l int = len(nums)
	var first_one int = 0
	var last_one int = l - 1
	for first_one < l && nums[first_one] != 1 {
		first_one++
	}
	for last_one >= 0 && nums[last_one] != 1 {
		last_one--
	}
	var total_one int = 0
	for i := first_one; i <= last_one; i++ {
		if nums[i] == 1 {
			total_one++
		}
	}
	if total_one == 0 || total_one == last_one-first_one+1 {
		return 0
	}
	nums = append(nums, nums...)
	var cur_one_cnt int = 0
	var left int = first_one
	var right int = first_one
	for (right - left + 1) <= total_one {
		if nums[right] == 1 {
			cur_one_cnt++
		}
		right++
	}
	right--
	var res int = l
	swap_cnt := total_one - cur_one_cnt
	if swap_cnt < res {
		res = swap_cnt
	}
	for left <= last_one && right < l*2 {
		if nums[left] == 1 {
			cur_one_cnt--
		}
		if nums[right+1] == 1 {
			cur_one_cnt++
		}
		swap_cnt := total_one - cur_one_cnt
		if swap_cnt < res {
			res = swap_cnt
		}
		left++
		right++
	}
	return res
}
