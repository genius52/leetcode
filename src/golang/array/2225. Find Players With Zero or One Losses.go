package array

import "sort"

func findWinners(matches [][]int) [][]int {
	var win_cnt map[int]bool = make(map[int]bool)
	var lose_cnt map[int]int = make(map[int]int)
	for _,match := range matches{
		win_cnt[match[0]] = true
		lose_cnt[match[1]]++
	}
	var res [][]int = make([][]int,2)
	for k,_ := range win_cnt{
		if _,ok := lose_cnt[k];!ok{
			res[0] = append(res[0],k)
		}
	}
	for k,v := range lose_cnt{
		if v == 1{
			res[1] = append(res[1],k)
		}
	}
	sort.Ints(res[0])
	sort.Ints(res[1])
	return res
}