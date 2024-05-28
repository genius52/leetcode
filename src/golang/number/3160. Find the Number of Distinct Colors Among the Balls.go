package number

func QueryResults(limit int, queries [][]int) []int {
	var l int = len(queries)
	var color_status map[int]int = make(map[int]int)
	var idx_status map[int]int = make(map[int]int)
	var total_colors int = 0
	var res []int = make([]int, l)
	for i := 0; i < l; i++ {
		idx := queries[i][0]
		new_color := queries[i][1]
		var old_color int = 0
		if _, ok := idx_status[idx]; ok {
			old_color = idx_status[idx]
		}
		if new_color == old_color {
			res[i] = total_colors
			continue
		}
		if old_color == 0 {
			if _, ok := color_status[new_color]; !ok {
				total_colors++
			}
			idx_status[idx] = new_color
			color_status[new_color]++
		} else {
			color_status[old_color]--
			color_status[new_color]++
			idx_status[idx] = new_color
			if color_status[old_color] == 0 {
				delete(color_status, old_color)
				total_colors--
			}
			if color_status[new_color] == 1 {
				total_colors++
			}
		}
		res[i] = total_colors
	}
	return res
}
