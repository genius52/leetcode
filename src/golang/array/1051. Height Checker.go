package array

import "sort"

func HeightChecker(heights []int) int {
	var l int = len(heights)
	var sorted_heights []int = make([]int,l)
	for i,v := range heights{
		sorted_heights[i] = v
	}
	sort.Ints(sorted_heights)
	var res int = 0
	for i := 0;i < l;i++{
		if(heights[i] != sorted_heights[i]){
			res++
		}
	}
	return res
}