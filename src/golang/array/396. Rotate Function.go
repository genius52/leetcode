package array

//A = [4, 3, 2, 6]
//
//F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
//F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
//F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
//F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26
//
//So the maximum value of F(0), F(1), F(2), F(3) is F(3) = 26.
func MaxRotateFunction(A []int) int {
	var l int = len(A)
	var sum int = 0
	for i := 0;i < l;i++ {
		sum += A[i]
	}
	var res int = 0
	for i := 0;i < l;i++{
		res += (A[i] * i)
	}
	var prev int = res
	for i := 1;i < l;i++{
		var cur int = prev + sum - A[l - i] * l
		if cur > res{
			res = cur
		}
		prev = cur
	}
	return res
}