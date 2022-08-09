package array

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	var l1 int = len(items1)
	var l2 int = len(items2)
	var record map[int]int = make(map[int]int)
	for i := 0;i < l1;i++{
		record[items1[i][0]] += items1[i][1]
	}
	for i := 0;i < l2;i++{
		record[items2[i][0]] += items2[i][1]
	}
	var res [][]int
	for k,v := range record{
		res = append(res,[]int{k,v})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}