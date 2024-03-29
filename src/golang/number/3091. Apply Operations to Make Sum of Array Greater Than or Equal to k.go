package number

func MinOperations3091(k int) int {
	if k == 1 {
		return 0
	}
	var min_cnt int = k
	for add_cnt := 1; add_cnt <= k-1; add_cnt++ {
		cur := (k-1-add_cnt)/(1+add_cnt) + add_cnt
		if (k-1-add_cnt)%(1+add_cnt) != 0 {
			cur++
		}
		if cur < min_cnt {
			min_cnt = cur
		}
	}
	return min_cnt
}
