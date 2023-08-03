package array

func CountCompleteSubarrays(nums []int) int {
	var record map[int]int = make(map[int]int)
	var l int = len(nums)
	for i := 0; i < l; i++ {
		record[nums[i]]++
	}
	var total_diff_cnt int = len(record)
	var res int = 0
	for i := 0; i < l; i++ {
		var cur map[int]int = make(map[int]int)
		for j := i; j < l; j++ {
			cur[nums[j]]++
			if len(cur) == total_diff_cnt {
				res += l - j
				break
			}
		}
	}
	return res
}
