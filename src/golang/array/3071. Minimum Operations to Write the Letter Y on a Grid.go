package array

func MinimumOperationsToWriteY(grid [][]int) int {
	var l int = len(grid)
	var y_cnt [3]int
	var other_cnt [3]int
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if (i == j && i <= l/2) || (i+j == l-1 && i < l/2) || (i >= l/2 && j == l/2) {
				y_cnt[grid[i][j]]++
			} else {
				other_cnt[grid[i][j]]++
			}
		}
	}
	var res int = l * l
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if i == j {
				continue
			}
			//set all y = i,set other = j
			change_cnt := l + l/2 - y_cnt[i] + l*l - l - l/2 - other_cnt[j]
			if change_cnt < res {
				res = change_cnt
			}
		}
	}
	return res
}
