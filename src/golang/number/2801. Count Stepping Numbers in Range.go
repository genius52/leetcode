package number

func count_countSteppingNumbers(s string, l int, cur_len int, last_num int) int {
	if cur_len > l {
		return 0
	}
	return 1
}

func countSteppingNumbers(low string, high string) int {
	var l1 int = len(low)
	var l2 int = len(high)
	return count_countSteppingNumbers(high, l2, 0, -1) - count_countSteppingNumbers(low, l1, 0, -1)
}
