package array

func MaxSubarrayLength(nums []int, k int) int {
	var l int = len(nums)
	var res int = 1
	var left int = 0
	var right int = 0
	var record map[int]int = make(map[int]int)
	for left < l && right < l {
		var over bool = false
		for right < l {
			record[nums[right]]++
			if record[nums[right]] > k {
				right++
				over = true
				break
			}
			right++
		}
		var cur_len int = 0
		if over {
			cur_len = right - left - 1
		} else {
			cur_len = right - left
		}
		if cur_len > res {
			res = cur_len
		}
		for left < right {
			record[nums[left]]--
			if record[nums[left]] == k {
				left++
				break
			}
			left++
		}
	}
	return res
}
