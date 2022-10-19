package number

import (
	"math"
	"strconv"
)

//172 + 271 = 443 so we return true.
func reverse_sumOfNumberAndReverse(n int)int{
	var n2 int = 0
	for n > 0{
		n2 *= 10
		n2 += n % 10
		n /= 10
	}
	return n2
}

func SumOfNumberAndReverse(num int) bool {
	var l int = len(strconv.Itoa(num))
	var base int = int(math.Pow10(l - 1))
	var start int = min_int(num /2,base)
	var end int = min_int( base * 5,num)
	for start <= end{
		var sum int = start + reverse_sumOfNumberAndReverse(start)
		if sum == num{
			return true
		}
		start++
	}
	return false
}