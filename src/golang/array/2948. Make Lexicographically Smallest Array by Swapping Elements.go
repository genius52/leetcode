package array

import "sort"

//func union_find(groups []int, cur int) int {
//	if groups[cur] != cur {
//		groups[cur] = union_find(groups, groups[cur])
//	}
//	return groups[cur]
//}

func LexicographicallySmallestArray(nums []int, limit int) []int {
	var l int = len(nums)
	var record [][2]int = make([][2]int, l)
	for i := 0; i < l; i++ {
		record[i] = [2]int{nums[i], i}
	}
	sort.Slice(record, func(i, j int) bool {
		return record[i][0] < record[j][0]
	})
	var groups [][][2]int = make([][][2]int, 1)
	groups[0] = append(groups[0], [2]int{record[0][0], record[0][1]})
	var idx int = 0
	for i := 1; i < l; i++ {
		if record[i][0]-record[i-1][0] <= limit {
			groups[idx] = append(groups[idx], [2]int{record[i][0], record[i][1]})
		} else {
			var add [][2]int = make([][2]int, 1)
			add[0] = [2]int{record[i][0], record[i][1]}
			groups = append(groups, add)
			idx++
		}
	}
	var res []int = make([]int, l)
	for k := 0; k < len(groups); k++ {
		sort.Slice(groups[k], func(i, j int) bool {
			return groups[k][i][0] < groups[k][j][0]
		})
		var index []int
		for _, ele := range groups[k] {
			index = append(index, ele[1])
		}
		sort.Ints(index)
		for i := 0; i < len(groups[k]); i++ {
			res[index[i]] = groups[k][i][0]
		}
	}
	return res
}
