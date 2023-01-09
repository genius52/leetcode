package number

func countDigits(num int) int {
	var res int = 0
	var n int = num
	for n > 0{
		mod := n % 10
		if mod != 0 && num % mod == 0{
			res++
		}
		n /= 10
	}

	return res
}