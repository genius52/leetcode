package array

func minimumIndex(nums []int) int {
	var l int = len(nums)
	var record map[int]int = make(map[int]int)
	var max_freq_num int = -1
	for i := 0; i < l; i++ {
		record[nums[i]]++
		if record[nums[i]] > l/2 {
			max_freq_num = nums[i]
		}
	}
	if max_freq_num == -1 {
		return -1
	}
	var max_freq_cnt int = record[max_freq_num]
	var left_cnt int = 0
	for i := 0; i < l; i++ {
		if nums[i] == max_freq_num {
			left_cnt++
		}
		right_cnt := max_freq_cnt - left_cnt
		if (left_cnt*2) > (i+1) && (right_cnt*2) > (l-1-i) {
			return i
		}
	}
	return -1
}
