package number

func sumOfTheDigitsOfHarshadNumber(x int) int {
	var sum int = 0
	var n int = x
	for n > 0 {
		m := n % 10
		sum += m
		n /= 10
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}
