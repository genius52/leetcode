package number

func FindSmallestInteger(nums []int, value int) int {
	var number_cnt []int = make([]int, value)
	for _, n := range nums {
		mod := n % value
		if mod >= 0 {
			number_cnt[mod]++
		} else {
			number_cnt[mod+value]++
		}
	}
	var min_cnt_num int = -1
	var min_cnt int = 1<<31 - 1
	for i := 0; i < value; i++ {
		if number_cnt[i] == 0 {
			return i
		}
		if number_cnt[i] < min_cnt {
			min_cnt = number_cnt[i]
			min_cnt_num = i
		}
	}
	return min_cnt_num + min_cnt*value
}
