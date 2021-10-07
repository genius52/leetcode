package number

func check_sumFourDivisors(n int)int{
	var div int = -1
	for i := 2;i * i <= n;i++{
		if n % i == 0{
			if div != -1{
				return 0
			}else{
				div = i
			}
		}
	}
	if div == -1 || n / div == div{
		return 0
	}
	return 1 + n + div + n /div
}

func sumFourDivisors(nums []int) int {
	var res int = 0
	for _,n := range nums{
		res += check_sumFourDivisors(n)
	}
	return res
}