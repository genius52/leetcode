package number

func minimumPossibleSum(n int, target int) int64 {
	var sum int64 = 0
	var used_num map[int]bool = make(map[int]bool)
	var cur_num int = 1
	for n > 0 && cur_num < target {
		forbid_num := target - cur_num
		if _, ok := used_num[forbid_num]; !ok {
			used_num[cur_num] = true
			sum += int64(cur_num)
			n--
		}
		cur_num++
	}
	for n > 0 {
		sum += int64(cur_num)
		cur_num++
		n--
	}
	return sum
}
