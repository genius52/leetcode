package array

func PrevPermOpt1(A []int) []int {
	var l int = len(A)
	var left int = -1
	for i := l - 1; i > 0;i--{
		if A[i-1] > A[i]{
			left = i - 1
			break
		}
	}
	if left == -1{
		return A
	}
	var right int = left + 1
	for i := left + 1; i < l; i++{
		if A[i] < A[left] && A[i] > A[right]{
			right = i
		}
	}
	A[left],A[right] = A[right],A[left]
	return A
}