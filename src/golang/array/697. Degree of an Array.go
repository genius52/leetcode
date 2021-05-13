package array

func findShortestSubArray(nums []int) int {
	var l int = len(nums)
	var record map[int]int = make(map[int]int)
	var max_cnt int = 0
	var max_cnt_num []int
	for _,n := range nums{
		record[n]++
		if record[n] > max_cnt{
			max_cnt = record[n]
		}
	}
	for n,cnt := range record{
		if cnt == max_cnt{
			max_cnt_num = append(max_cnt_num,n)
		}
	}
	var res int = 2147483647
	for _, n := range max_cnt_num{
		var left int = 0
		for left < l && nums[left] != n{
			left++
		}
		var right int = l - 1
		for left >= 0 && nums[right] != n{
			right--
		}
		cur_len := right - left + 1
		if cur_len < res{
			res = cur_len
		}
	}
	return res
}