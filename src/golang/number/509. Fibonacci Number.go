package number

//F(0) = 0, F(1) = 1
//F(n) = F(n - 1) + F(n - 2), for n > 1.
func fib(N int) int {
	if N <= 1{
		return N
	}
	var last1 int = 0
	var last2 int = 1
	var cur int = 0
	for i := 1;i < N;i++{
		cur = last1 + last2
		last1 = last2
		last2 = cur
	}
	return cur
}