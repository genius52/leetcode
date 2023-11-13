package array

func minGroupsForValidAssignment(nums []int) int {
	var num_cnt map[int]int = make(map[int]int)
	for _, n := range nums {
		num_cnt[n]++
	}
	var groups []int
	var limit int = 0
	var res int = 0
	for _, cnt := range num_cnt {
		groups = append(groups, cnt)
		limit += cnt
		if cnt > res {
			res = cnt
		}
	}
	for group_size := 1; group_size < res; group_size++ {
		var cur_cnt int = 0
		for i := 0; i < len(groups); i++ {

		}
		if cur_cnt < res {
			res = cur_cnt
		}
	}
	return res
}
