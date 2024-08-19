package number

func winningPlayerCount(n int, pick [][]int) int {
	var player_color_cnt [][11]int = make([][11]int, n)
	for _, player_color := range pick {
		player_color_cnt[player_color[0]][player_color[1]]++
	}
	var res int = 0
	for i := 0; i < n; i++ {
		for j := 0; j <= 10; j++ {
			if player_color_cnt[i][j] > i {
				res++
				break
			}
		}
	}
	return res
}
