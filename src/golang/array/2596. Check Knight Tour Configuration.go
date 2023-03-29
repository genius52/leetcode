package array

func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}
	var rows int = len(grid)
	var dirs [][2]int = [][2]int{[2]int{1, 2}, [2]int{-1, 2}, [2]int{1, -2}, [2]int{-1, -2}, [2]int{2, 1}, [2]int{-2, 1}, [2]int{2, -1}, [2]int{-2, -1}}
	var visit point
	visit.x = 0
	visit.y = 0
	var target int = 1
	for target < rows*rows {
		var find bool = false
		for _, dir := range dirs {
			var next point
			next.x = visit.x + dir[0]
			next.y = visit.y + dir[1]
			if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= rows {
				continue
			}
			if grid[next.x][next.y] == target {
				find = true
				visit = next
				break
			}
		}
		if !find {
			return false
		}
		target++
	}
	return true
}
