package array

import "math"

//Input: [1,0,5]
//
//Output: 3
//
//Explanation:
//1st move:    1     0 <-- 5    =>    1     1     4
//2nd move:    1 <-- 1 <-- 4    =>    2     1     3
//3rd move:    2     1 <-- 3    =>    2     2     2
func findMinMoves(machines []int) int {
	var l int = len(machines)
	var sum int = 0
	for i := 0;i < l;i++{
		sum += machines[i]
	}
	if sum % l != 0{
		return -1
	}
	var target int = sum / l
	var res int = 0
	var rest int = 0
	for i := 0;i < l;i++{
		rest += machines[i] - target
		res = max_int_number(res,int(math.Abs(float64(rest))),int(math.Abs(float64(machines[i] - target))))
	}
	return res
}