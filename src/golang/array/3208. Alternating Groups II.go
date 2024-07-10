package array

func NumberOfAlternatingGroups2(colors []int, k int) int {
	var l int = len(colors)
	var right int = 0
	var equal_cnt int = 0
	for i := 0; i < k-1; i++ {
		if colors[right] == colors[right+1] {
			equal_cnt++
		}
		right++
	}
	var res int = 0
	if equal_cnt == 0 {
		res++
	}
	for left := 0; left < l-1; left++ {
		if colors[left] == colors[left+1] {
			equal_cnt--
		}
		if colors[(left+k-1)%l] == colors[(left+k)%l] {
			equal_cnt++
		}
		if equal_cnt == 0 {
			res++
		}
	}
	return res
}
