package number

func FindNonMinOrMax(nums []int) int {
	var min_num int = 1<<31 - 1
	var max_num int = -1
	for _, n := range nums {
		if n > max_num {
			if min_num != max_num && max_num != -1 && min_num != (1<<31-1) {
				return max_num
			}
			max_num = n
		}
		if n < min_num {
			if min_num != max_num && max_num != -1 && min_num != (1<<31-1) {
				return min_num
			}
			min_num = n
		}
		if n > min_num && n < max_num {
			return n
		}
	}
	return -1
}
