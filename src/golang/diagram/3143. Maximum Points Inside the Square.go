package diagram

import "sort"

func maxPointsInsideSquare(points [][]int, s string) int {
	var l int = len(points)
	var point_key [][3]int = make([][3]int, l)
	for i := 0; i < l; i++ {
		point_key[i] = [3]int{points[i][0], points[i][1], int(s[i] - 'a')}
	}
	//按照变长从小到大排序
	sort.Slice(point_key, func(i, j int) bool {
		return max_int(abs_int(point_key[i][0]), abs_int(point_key[i][1])) < max_int(abs_int(point_key[j][0]), abs_int(point_key[j][1]))
	})
	var visited [26]bool
	for i := 0; i < l; i++ {
		if visited[point_key[i][2]] {
			for j := i - 1; j >= 0; j-- {
				cur_len := max_int(abs_int(point_key[i][0]), abs_int(point_key[i][1]))
				pre_len := max_int(abs_int(point_key[j][0]), abs_int(point_key[j][1]))
				if pre_len != cur_len {
					return j + 1
				}
			}
			return 0
		}
		visited[point_key[i][2]] = true
	}
	return l
}
