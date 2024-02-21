package number

func MaximumLength(nums []int) int {
	var num_cnt map[int]int = make(map[int]int)
	var one_cnt int = 0
	for _, n := range nums {
		if n == 1 {
			one_cnt++
			continue
		}
		num_cnt[n]++
	}
	var l int = len(nums)
	var res int = max_int(1, (one_cnt-1)|1)
	for n, cnt := range num_cnt {
		if cnt == 1 {
			continue
		}
		var cur_len int = 2
		for cur_len <= l {
			n *= n
			if cnt2, ok := num_cnt[n]; ok {
				if cnt2 < 2 {
					cur_len++
					break
				} else {
					cur_len += 2
				}
			} else {
				cur_len--
				break
			}
		}
		res = max_int(res, cur_len)
	}
	return res
}
