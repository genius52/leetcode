package diagram

import (
	"math"
)

func min_int(a,b int)int{
	if a < b {
		return a
	}else{
		return b
	}
}

func min_int_number(nums ...int)int{
	var min int = math.MaxInt32
	for _,n := range nums{
		if n < min{
			min = n
		}
	}
	return min
}

func max_int_number(nums ...int)int{
	var max int = math.MinInt32
	for _,n := range nums{
		if n > max{
			max = n
		}
	}
	return max
}

func max_int(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}