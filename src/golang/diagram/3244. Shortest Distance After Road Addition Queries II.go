package diagram

import "sort"

func ShortestDistanceAfterQueries3244(n int, queries [][]int) []int {
	var record []int = make([]int, n)
	for i := 0; i < n; i++ {
		record[i] = i
	}
	var l int = len(queries)
	var res []int = make([]int, l)
	for k := 0; k < l; k++ {
		from := queries[k][0]
		to := queries[k][1]
		//delete elements between (from,to)
		begin := sort.Search(len(record), func(i int) bool {
			return record[i] > from
		})
		end := sort.Search(len(record), func(i int) bool {
			return record[i] >= to
		})
		record = append(record[0:begin], record[end:]...)
		res[k] = len(record) - 1
	}
	return res
}
