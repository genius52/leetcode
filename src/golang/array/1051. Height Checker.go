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


func HeightChecker2(heights []int) int {
	var record []int = make([]int,101)
	for _,h := range heights{
		record[h]++
	}
	var res int = 0
	var pos int = 0
	for i := 1;i <= 100;i++{
		if(record[i] == 0){
			continue
		}
		for(record[i] > 0){
			if(heights[pos] != i){
				res++
			}
			pos++
			record[i]--
		}
	}
	return res
}