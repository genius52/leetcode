package number

import "sort"

func topKFrequent(nums []int, k int) []int {
	var record map[int]int = make(map[int]int)
	for _,n := range nums{
		record[n]++
	}
	var freq_num [][]int
	for k,v := range record{
		freq_num = append(freq_num,[]int{k,v})
	}
	sort.Slice(freq_num, func(i, j int) bool {
		return freq_num[i][1] > freq_num[j][1]
	})
	var res []int = make([]int,k)
	for i := 0;i < k;i++{
		res[i] = freq_num[i][0]
	}
	return res
}