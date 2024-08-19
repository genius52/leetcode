package number

func canAliceWin(nums []int) bool {
	var sum int = 0
	var single_sum int = 0
	var double_sum int = 0
	for _, n := range nums {
		sum += n
		if n >= 1 && n <= 9 {
			single_sum += n
		}
		if n >= 10 && n <= 99 {
			double_sum += n
		}
	}
	if single_sum > sum-single_sum {
		return true
	}
	if double_sum > sum-double_sum {
		return true
	}
	return false
}
