package diagram

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	var connect []int = make([]int,n)
	for _,r := range roads{
		connect[r[0]]++
		connect[r[1]]++
	}
	sort.Ints(connect)
	var res int64 = 0
	var score int = 1
	for i := 0;i < n;i++{
		res += int64(connect[i] * score)
		score++
	}
	return res
}