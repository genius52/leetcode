package number

import (
	"math"
)

//Input: low = 1000, high = 13000
//Output: [1234,2345,3456,4567,5678,6789,12345]
func SequentialDigits(low int, high int) []int {
	var res []int
	var bit_len_low int = int(math.Log10(float64(low))) + 1
	var bit_len_high int = int(math.Log10(float64(high))) + 1
	var n int = low
	var start_low int = 0
	for n > 0{
		start_low = n % 10
		n = n / 10
	}
	for l := bit_len_low;l <= bit_len_high;l++{
		var begin int = 1
		if l == bit_len_low {
			begin = start_low
		}
		for start := begin;start + l <= 10;start++{
			var num int = 0
			var next int = start
			for j := 0;j < l;j++{
				num *= 10
				num += next
				next++
			}
			if num < low {
				continue
			}
			if num > high{
				return res
			}
			res = append(res,num)
		}
	}
	return res
}