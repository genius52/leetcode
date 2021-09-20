package array

import "sort"

func ArrayRankTransform(arr []int) []int{
	var l int = len(arr)
	var arr2 []int = make([]int,l)
	copy(arr2,arr)
	sort.Ints(arr2)
	var record map[int][]int = make(map[int][]int)
	for i := 0;i < l;i++{
		record[arr[i]] = append(record[arr[i]],i)
	}
	var rank int = 1
	for i := 0;i < l;{
		n := arr2[i]
		var dup_len int = len(record[n])
		for j := 0;j < dup_len;j++{
			arr[record[n][j]] = rank
		}
		rank++
		i += dup_len
	}
	return arr
}