package number

import "math"

func CountTriples(n int) int {
	var res int = 0
	for i := 1;i <= n;i++{
		for j := i - 1;j > 0;j--{
			a :=  int(math.Sqrt(float64(i * i - j * j)))
			if (a * a + j * j) == (i * i){
				res++
			}
		}
	}
	return res
}