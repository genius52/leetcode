package array

import "sort"

//Input: items = [[1,2],[3,2],[2,4],[5,6],[3,5]], queries = [1,2,3,4,5,6]
//Output: [2,4,5,5,6,6]
//Input: items = [[1,2],[1,2],[1,3],[1,4]], queries = [1]
//Output: [4]
func maximumBeauty(items [][]int, queries []int) []int {
	var l int = len(queries)
	var res []int = make([]int,l)
	sort.Slice(items, func(i, j int) bool {
		if items[i][1] != items[j][1]{
			return items[i][1] > items[j][1]
		}else{
			return items[i][0] < items[j][0]
		}
	})
	for i := 0;i < l;i++{
		var find bool = false
		for j := 0;j < len(items);j++{
			if items[j][0] <= queries[i]{
				res[i] = items[j][1]
				find = true
				break
			}
		}
		if !find{
			res[i] = 0
		}
	}
	return res
}