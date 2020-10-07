package array

//Input: rowSum = [3,8], colSum = [4,7]
//Output: [[3,0],
//         [1,7]]
//Explanation:
//0th row: 3 + 0 = 0 == rowSum[0]
//1st row: 1 + 7 = 8 == rowSum[1]
//0th column: 3 + 1 = 4 == colSum[0]
//1st column: 0 + 7 = 7 == colSum[1]
//The row and column sums match, and all matrix elements are non-negative.
//Another possible matrix is: [[1,2],
//                             [3,5]]
func RestoreMatrix(rowSum []int, colSum []int) [][]int {
	var rows int = len(rowSum)
	var columns int = len(colSum)
	var res [][]int = make([][]int,rows)
	for i := 0;i < rows;i++{
		res[i] = make([]int,columns)
	}
	for i := 0;i < rows;i++{
		for j := 0;j < columns;j++{
			var val int = 0
			if rowSum[i] < colSum[j]{
				val = rowSum[i]
			}else{
				val = colSum[j]
			}
			rowSum[i] -= val
			colSum[j] -= val
			res[i][j] = val
		}
	}
	return res
}