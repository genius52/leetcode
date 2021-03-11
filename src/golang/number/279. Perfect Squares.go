package number

import "math"

func NumSquares(n int) int {
	var record []int = make([]int,n + 1)
	record[0] = 0
	for i := 1;i <= n;i++{
		min := math.MaxInt32
		for j := 1;(j * j) <= i;j++{
			min = min_int_number(min,record[i - j*j] + 1)
		}
		record[i] = min
	}
	return record[n]
}