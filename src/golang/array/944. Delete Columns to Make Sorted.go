package array

func minDeletionSize(A []string) int {
	var rows int = len(A)
	var columns int = len(A[0])
	var res int = 0
	for j := 0;j < columns;j++{
		for i := 1;i < rows;i++{
			if A[i][j] < A[i - 1][j]{
				res++
				break
			}
		}
	}
	return res
}