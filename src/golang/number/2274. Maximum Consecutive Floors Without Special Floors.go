package number

import "sort"

func MaxConsecutive(bottom int, top int, special []int) int {
	var data []int
	for _,n := range special{
		if n >= bottom && n <= top{
			data = append(data,n)
		}
	}
	data = append(data,bottom - 1)
	data = append(data,top + 1)
	sort.Ints(data)
	var l int = len(data)
	var max_val int = 0
	for i := 1;i < l;i++{
		cur := data[i] - data[i - 1] - 1
		if cur > max_val{
			max_val = cur
		}
	}
	return max_val
}