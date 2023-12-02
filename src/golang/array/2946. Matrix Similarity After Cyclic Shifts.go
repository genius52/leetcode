package array

func areSimilar(mat [][]int, k int) bool {
	var rows int = len(mat)
	var columns int = len(mat[0])
	k %= columns
	if k == 0 {
		return true
	}
	for i := 0; i < rows; i += 2 {
		for j := 0; j < columns; j++ {
			if mat[i][j] != mat[i][(j+k)%columns] {
				return false
			}
		}
	}
	for i := 1; i < rows; i += 2 {
		for j := 0; j < columns; j++ {
			if mat[i][j] != mat[i][(j+columns-k)%columns] {
				return false
			}
		}
	}
	return true
}
