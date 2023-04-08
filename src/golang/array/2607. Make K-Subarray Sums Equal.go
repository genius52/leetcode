package array

import "sort"

func MakeSubKSumEqual(arr []int, k int) int64 {
	var l int = len(arr)
	var res int64 = 0
	var visited []bool = make([]bool, l)
	//arr[i] == arr[i + k] == arr[(i + 2 * k) % l] ....
	for i := 0; i < l; i++ {
		if visited[i] {
			continue
		}
		var data []int = []int{arr[i]}
		visited[i] = true
		var j int = (i + k) % l
		for j != i {
			if !visited[j] {
				data = append(data, arr[j])
				visited[j] = true
			}
			j = (j + k) % l
		}
		sort.Ints(data)
		var target int = 0
		var l2 int = len(data)
		if l2 == 0 {
			continue
		}
		if (l2 | 1) == l2 {
			target = data[l2/2]
		} else {
			target = (data[l2/2-1] + data[l2/2]) / 2
		}
		for m := 0; m < l2; m++ {
			res += abs_int64(int64(data[m] - target))
		}
	}
	return res
}
