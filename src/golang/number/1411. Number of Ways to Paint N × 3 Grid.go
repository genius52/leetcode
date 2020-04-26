package number
//Input: n = 1
//Output: 12
func NumOfWays(n int) int {
	var diff3 []int = make([]int,n)
	var same2 []int = make([]int,n)
	diff3[0] = 6
	same2[0] = 6
	for i := 1;i < n;i++{
		diff3[i] = (2 * diff3[i - 1] + 2 * same2[i - 1])% (1e9 + 7)
		same2[i] = (2 * diff3[i - 1] + 3 * same2[i - 1])% (1e9 + 7)
	}
	return (int)(diff3[n - 1] + same2[n - 1])% (1e9 + 7)
}
