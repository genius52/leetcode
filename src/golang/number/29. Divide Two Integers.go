package number
import "math"

func Divide(dividend int, divisor int) int {
	if(dividend == math.MinInt32 && divisor == -1) {
		return math.MaxInt32
	}
	var dividend_is_neg bool = dividend < 0
	var divisor_is_neg bool = divisor < 0
	var result_is_neg bool = (!dividend_is_neg && divisor_is_neg) || (dividend_is_neg && !divisor_is_neg)
	if(result_is_neg){
		if(dividend_is_neg){
			dividend = -dividend
		}
		if(divisor_is_neg){
			divisor = -divisor
		}
	}
	if(dividend_is_neg && divisor_is_neg){
		dividend = -dividend
		divisor = -divisor
	}

	var res int = 0
	for(dividend > 0){
		var d int = divisor
		var n int = 1
		for(d <= dividend){
			d = d << 1
			n = n << 1
		}
		d = d >> 1
		n = n >> 1
		dividend -= d
		res = res | n
	}
	if(result_is_neg){
		return -res
	}
	return res
}