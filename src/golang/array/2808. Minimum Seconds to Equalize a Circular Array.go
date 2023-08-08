package array

func MinimumSeconds(nums []int) int {
	var l int = len(nums)
	var num_idx map[int][]int = make(map[int][]int)
	for i := 0; i < l; i++ {
		num_idx[nums[i]] = append(num_idx[nums[i]], i)
	}
	var res int = l - 1
	for _, idx := range num_idx {
		if len(idx) == 1 {
			cnt := (l - 1) / 2
			if cnt < res {
				res = cnt
			}
			continue
		}
		cur_len := len(idx)
		var max_cnt int = (l - idx[cur_len-1] + idx[0]) / 2
		for i := 1; i < cur_len; i++ {
			cnt := (idx[i] - idx[i-1]) / 2
			if cnt > max_cnt {
				max_cnt = cnt
			}
		}
		if max_cnt < res {
			res = max_cnt
		}
	}
	return res
}
