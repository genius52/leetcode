package array

func numberOfBeams(bank []string) int {
	var rows int = len(bank)
	var columns int = len(bank[0])
	var res int = 0
	var pre_cnt int = 0
	for i := 0; i < rows; i++ {
		var find bool = false
		var cnt int = 0
		for j := 0; j < columns; j++ {
			if bank[i][j] == '1' {
				find = true
				cnt++
			}
		}
		if find {
			if pre_cnt != 0 {
				res += pre_cnt * cnt
				pre_cnt = cnt
			} else {
				pre_cnt = cnt
			}
		}
	}
	return res
}
