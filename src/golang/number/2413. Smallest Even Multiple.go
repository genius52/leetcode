package number

func smallestEvenMultiple(n int) int {
	if n | 1 != n{
		return n
	}
	return 2 * n
}