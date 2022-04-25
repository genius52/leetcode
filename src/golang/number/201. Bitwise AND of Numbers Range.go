package number

import "math"

//Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.
//
//Example 1:
//
//Input: [5,7]
//Output: 4
func rangeBitwiseAnd(left int, right int) int {
	if right >= left * 2{
		return 0
	}
	var res = math.MaxInt32
	for left <= right{
		res &= left
		left++
	}
	return res
}

func RangeBitwiseAnd(left int, right int) int{
	if left == right{
		return left
	}
	var mask int = 1 << 31
	var res int = 0
	for mask > 0 || (mask & left == 0 && mask & right == 0){
		if ((mask | left) == left && (mask | right) != right) || ((mask | left) != left && (mask | right) == right) {
			break
		}
		if (mask | left) == left && (mask | right) == right{
			res |= mask
		}

		mask >>= 1
	}
	return res
}