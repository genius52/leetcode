package number

import "sort"

func maxIncreasingGroups(usageLimits []int) int {
	var l int = len(usageLimits)
	sort.Ints(usageLimits)
	var columns int = 0
	var total int = 0
	//一列一列构建，usageLimits[i]构建第i列
	for i := 0; i < l; i++ {
		total += usageLimits[i]
		//如果total大于等于columns+1，可以构建第"columns+1"列
		//
		if total >= ((columns+1)*(columns+2))/2 {
			columns++
		}
	}
	return columns
}
