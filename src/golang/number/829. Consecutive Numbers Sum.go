package number

import "math"

//x,x + 1,x + 2.... x + k - 1 = n,k个数字的等差数列
// kx + k * (k - 1)/2 = n
// x = (n -  (k * (k-1) /2)) / k
func ConsecutiveNumbersSum(n int) int {
	var res int = 1
	for k := 2;k <= int(math.Sqrt(float64(n * 2))) ;k++{
		x := (n - (k * (k-1) / 2))/ k
		if x < 0{
			break
		}
		mod := (n - (k * (k-1) / 2))% k
		if x >= 0 && mod == 0{
			res++
		}
	}
	return res
}