package number

func isPowerOfThree(n int) bool {
	for n > 1{
		if n % 3 != 0{
			return false
		}
		n = n / 3
	}
	return n == 1
}

func isPowerOfThree2(n int) bool {
	return n > 0 && (1162261467 % n) == 0
}
