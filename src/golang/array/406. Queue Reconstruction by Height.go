package array

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0]{
			return people[i][0] > people[j][0]
		}else{
			return people[i][1] < people[j][1]
		}
	})
	var l int = len(people)
	var res [][]int
	for i := 0;i < l;i++{
		if len(res) < people[i][1]{
			res = append(res,people[i])
		}else{
			tail := append([][]int{},res[people[i][1]:]...)
			res = append(res[0:people[i][1]],people[i])
			res = append(res,tail...)
		}
	}
	return res
}