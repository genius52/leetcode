package number

//func is_prime(n int) bool {
//	if n == 1 {
//		return false
//	}
//	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
//		if n%i == 0 {
//			return false
//		}
//	}
//	return true
//}

func check_mostFrequentPrime(mat [][]int, rows int, columns int, r int, c int, dir [2]int, pre_num int, num_cnt map[int]int) {
	if r < 0 || c < 0 || r == rows || c == columns {
		return
	}
	cur := pre_num*10 + mat[r][c]
	if cur >= 10 && is_prime(cur) {
		num_cnt[cur]++
	}
	check_mostFrequentPrime(mat, rows, columns, r+dir[0], c+dir[1], dir, cur, num_cnt)
}
func mostFrequentPrime(mat [][]int) int {
	var rows int = len(mat)
	var columns int = len(mat[0])
	var num_cnt map[int]int = make(map[int]int)
	var dirs [][2]int = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			for _, dir := range dirs {
				check_mostFrequentPrime(mat, rows, columns, i, j, dir, 0, num_cnt)
			}
		}
	}
	var max_cnt int = 0
	var max_num int = -1
	for n, cnt := range num_cnt {
		if cnt > max_cnt {
			max_cnt = cnt
			max_num = n
		} else if cnt == max_cnt {
			if n > max_num {
				max_num = n
			}
		}
	}

	return max_num
}
