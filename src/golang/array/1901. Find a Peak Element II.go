package array

func findPeakGrid(mat [][]int) []int {
	var rows int = len(mat)
	var columns int = len(mat[0])
	var r int = 0
	var c int = 0
	for true {
		move := false
		if r < rows-1 && mat[r][c] < mat[r+1][c] {
			move = true
			r++
		} else if r > 0 && mat[r][c] < mat[r-1][c] {
			move = true
			r--
		} else if c < columns-1 && mat[r][c] < mat[r][c+1] {
			move = true
			c++
		} else if c > 0 && mat[r][c] < mat[r][c-1] {
			move = true
			c--
		}
		if !move {
			return []int{r, c}
		}
	}
	return []int{}
}
