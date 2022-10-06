package number


func commonFactors(a int, b int) int {
	var res int = 1
	for i := 2;i <= min_int(a,b);i++{
		if a % i == 0 && b % i == 0{
			res++
		}
	}
	return res
}