package diagram

func countBlackBlocks(m int, n int, coordinates [][]int) []int64 {
	var record map[int64]bool = make(map[int64]bool)
	var left_up [][]int = [][]int{{-1, -1}, {-1, 0}, {0, -1}}
	var right_up [][]int = [][]int{{-1, 1}, {-1, 0}, {0, 1}}
	var left_down [][]int = [][]int{{1, -1}, {1, 0}, {0, -1}}
	var right_down [][]int = [][]int{{1, 1}, {1, 0}, {0, 1}}
	var block_dir [][][]int = [][][]int{left_up, right_up, left_down, right_down}
	var res []int64 = make([]int64, 5)
	res[0] = int64((m - 1) * (n - 1))
	for _, point := range coordinates {
		for _, dirs := range block_dir {
			var cnt int = 0
			var overflow bool = false
			for _, dir := range dirs {
				x2 := point[0] + dir[0]
				y2 := point[1] + dir[1]
				if x2 < 0 || x2 >= m || y2 < 0 || y2 >= n {
					overflow = true
					break
				}
				k := int64(x2*100000 + y2)
				if _, ok := record[k]; ok {
					cnt++
				}
			}
			if !overflow {
				res[cnt]--
				res[cnt+1]++
			}
		}
		record[int64(point[0]*100000+point[1])] = true
	}
	return res
}
