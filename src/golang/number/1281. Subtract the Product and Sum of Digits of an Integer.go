package number

func subtractProductAndSum(n int) int {
	var sum int = 0
	var times int = 1
	for n > 0{
		val := n % 10
		sum += val
		times *= val
		n = n/10
	}
	return times - sum
}