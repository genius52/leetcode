package number

import "math"

func judgeSquareSum(c int) bool {
	if c == 0{
		return true
	}
	limit := int(math.Sqrt(float64(c)))
	for i := limit;i >= 1;i--{
		rest := c - i * i
		j := int(math.Sqrt(float64(rest)))
		if j * j == rest{
			return true
		}
	}
	return false
}