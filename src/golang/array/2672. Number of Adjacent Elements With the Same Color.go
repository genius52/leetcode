package array

func colorTheArray(n int, queries [][]int) []int {
	var l int = len(queries)
	var res []int = make([]int, l)
	var pos_color []int = make([]int, n)
	var same_color_cnt int = 0
	for i := 0; i < l; i++ {
		pos := queries[i][0]
		new_color := queries[i][1]
		old_color := pos_color[pos]
		left := pos - 1
		right := pos + 1
		if old_color == 0 {
			if left >= 0 {
				if pos_color[left] == new_color {
					same_color_cnt++
				}
			}
			if right < n {
				if pos_color[right] == new_color {
					same_color_cnt++
				}
			}
		} else {
			if left >= 0 {
				if pos_color[left] == old_color {
					same_color_cnt--
				}
				if pos_color[left] == new_color {
					same_color_cnt++
				}
			}
			if right < n {
				if pos_color[right] == old_color {
					same_color_cnt--
				}
				if pos_color[right] == new_color {
					same_color_cnt++
				}
			}
		}
		pos_color[pos] = new_color
		res[i] = same_color_cnt
	}
	return res
}
