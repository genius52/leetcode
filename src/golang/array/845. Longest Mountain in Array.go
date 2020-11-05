package array

//Input: [2,1,4,7,3,2,5]
//Output: 5
func LongestMountain(A []int) int {
	var l int = len(A)
	if l <= 1{
		return 0
	}
	var left []int = make([]int,l)
	var right []int = make([]int,l)
	left[0] = 1
	right[l - 1] = 1
	for i := 1;i < l;i++{
		if A[i] > A[i - 1]{
			left[i] = left[i - 1] + 1
		}else{
			left[i] = 1
		}
		if A[l - 1 - i] > A[l - i]{
			right[l - 1 - i] = right[l - i] + 1
		}else{
			right[l - 1 - i] = 1
		}
	}
	var res int = 0
	for i := 0;i < l;i++{
		if left[i] > 1 && right[i] > 1{
			l := left[i] + right[i] - 1
			if l > res{
				res = l
			}
		}
	}
	return res
}