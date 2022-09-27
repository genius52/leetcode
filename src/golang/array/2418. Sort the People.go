package array

import "sort"

func sortPeople(names []string, heights []int) []string {
	var record [][2]int
	for i := 0;i < len(names);i++{
		record = append(record,[2]int{heights[i],i})
	}
	sort.Slice(record, func(i, j int) bool {
		return record[i][0] > record[j][0]
	})
	var res []string
	for i := 0;i < len(names);i++{
		res = append(res,names[record[i][1]])
	}
	return res
}