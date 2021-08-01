package number

import "math"

func isThree(n int) bool {
	limit := int(math.Sqrt(float64(n)))
	var res int = 2
	for i := 2;i <= limit;i++{
		mod := n % i
		divide := n / i
		if mod == 0{
			if divide == i{
				res++
			}else{
				res += 2
			}
		}
		if res > 3{
			return false
		}
	}
	return res == 3
}