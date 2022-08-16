package number

import "strconv"

func countSpecialNumbers(n int) int {
	var l int = len(strconv.Itoa(n))
	var pre []int = make([]int,l)
	pre[0] = 9
	var choice int = 9
	for i := 1;i < l;i++{
		pre[i] = pre[i - 1] * choice
		choice--
	}
	var res int = 0
	
	return res
}