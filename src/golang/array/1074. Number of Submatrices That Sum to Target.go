package array


//Input: matrix = [[1,-1],[-1,1]], target = 0
//Output: 5
func NumSubmatrixSumTarget(matrix [][]int, target int) int {
	var rows int = len(matrix)
	var columns int = len(matrix[0])
	var prefix [][]int = make([][]int,rows + 1)//prefix[i][j] = sum from 0,0 to i - 1,j - 1
	for i := 0;i <= rows;i++{
		prefix[i] = make([]int,columns + 1)
	}
	for i := 1;i <= rows;i++{
		for j := 1;j <= columns;j++{
			prefix[i][j] = matrix[i - 1][j - 1] + prefix[i - 1][j] + prefix[i][j - 1] - prefix[i - 1][j - 1]
		}
	}
	var res int = 0
	for i := 0;i <= rows;i++{
		for j := 0;j <= columns;j++{
			for r := i + 1;r <= rows;r++{
				for c := j + 1;c <= columns;c++{
					area := prefix[r][c] - prefix[r][j] - prefix[i][c] + prefix[i][j]
					if area == target{
						res++
					}
				}
			}
		}
	}
	return res
}