package number

func duplicateNumbersXOR(nums []int) int {
	var num_cnt map[int]int = make(map[int]int)
	for _, n := range nums {
		num_cnt[n]++
	}
	var res int = 0
	for n, cnt := range num_cnt {
		if cnt == 2 {
			res ^= n
		}
	}
	return res
}
