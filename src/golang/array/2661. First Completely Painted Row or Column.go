package array

func firstCompleteIndex(arr []int, mat [][]int) int {
	var rows int = len(mat)
	var columns int = len(mat[0])
	var row_cnt []int = make([]int, rows)
	var column_cnt []int = make([]int, columns)
	var val_pos [][2]int = make([][2]int, rows*columns+1)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			val_pos[mat[i][j]] = [2]int{i, j}
		}
	}
	for i := 0; i < len(arr); i++ {
		r := val_pos[arr[i]][0]
		c := val_pos[arr[i]][1]
		row_cnt[r]++
		column_cnt[c]++
		if row_cnt[r] == columns || column_cnt[c] == rows {
			return i
		}
	}
	return -1
}
