package number

func sumOfMultiples(n int) int {
	var res int = 0
	for i := 3; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			res += i
		}
	}
	return res
}
