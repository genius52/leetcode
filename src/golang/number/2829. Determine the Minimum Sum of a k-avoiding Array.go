package number

func MinimumSum2829(n int, k int) int {
	var res int = 0
	var used_num map[int]bool = make(map[int]bool)
	var cur_num int = 1
	for n > 0 {
		if cur_num >= k {
			res += cur_num
			cur_num++
			n--
			continue
		}
		if _, ok := used_num[k-cur_num]; !ok {
			res += cur_num
			n--
			used_num[cur_num] = true
		}
		cur_num++
	}
	return res
}
