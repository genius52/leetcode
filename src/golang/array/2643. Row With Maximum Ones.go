package array

func rowAndMaximumOnes(mat [][]int) []int {
	var rows int = len(mat)
	var columns int = len(mat[0])
	var res []int = make([]int, 2)
	for i := 0; i < rows; i++ {
		var one_cnt int = 0
		for j := 0; j < columns; j++ {
			if mat[i][j] == 1 {
				one_cnt++
			}
		}
		if one_cnt > res[1] {
			res[0] = i
			res[1] = one_cnt
		}
	}
	return res
}
