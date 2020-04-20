package number

import "math"

//Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.
//
//Example 1:
//
//Input: [5,7]
//Output: 4
func rangeBitwiseAnd(m int, n int) int {
	if n >= m * 2{
		return 0
	}
	var res = math.MaxInt32
	for m <= n{
		res &= m
		m++
	}
	return res
}