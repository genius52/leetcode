package number

import "sort"

func EarliestFullBloom(plantTime []int, growTime []int) int {
	var l int = len(plantTime)
	var record [][]int = make([][]int, l)
	for i := 0; i < l; i++ {
		record[i] = []int{plantTime[i], growTime[i]}
	}
	sort.Slice(record, func(i, j int) bool {
		//duration1 := max_int(record[i][0]+record[i][1], record[i][0]+record[j][0]+record[j][1])
		//duration2 := max_int(record[j][0]+record[j][1], record[j][0]+record[i][0]+record[i][1])
		//return duration1 < duration2
		return record[i][1] > record[j][1]
	})
	var available int = 0
	var last int = 0
	for i := 0; i < l; i++ {
		available += record[i][0]
		last = max_int(last, available+record[i][1])
	}
	return last
}
