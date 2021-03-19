package number

func isPerfectSquare(num int) bool {
	var n int = 1
	for n <= num{
		if n * n == num{
			return true
		}
		n++
	}
	return false
}