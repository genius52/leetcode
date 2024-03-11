package number

import "sort"

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	var l int = len(bottomLeft)
	var pos [][4]int = make([][4]int, l) //left_x,left_y,right_x,right_y
	for i := 0; i < l; i++ {
		pos[i] = [4]int{bottomLeft[i][0], bottomLeft[i][1], topRight[i][0], topRight[i][1]}
	}
	sort.Slice(pos, func(i, j int) bool {
		if pos[i][0] == pos[j][0] {
			return pos[i][1] <= pos[j][1]
		} else {
			return pos[i][0] < pos[j][0]
		}
	})
	var res int64 = 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if pos[i][2] <= pos[j][0] || pos[i][1] >= pos[j][3] || pos[i][3] <= pos[j][1] {
				continue
			}
			dis_x := min_int(pos[i][2], pos[j][2]) - max_int(pos[i][0], pos[j][0])
			dis_y := min_int(pos[i][3], pos[j][3]) - max_int(pos[i][1], pos[j][1])
			edge := min_int(dis_x, dis_y)
			res = max_int64(res, int64(edge)*int64(edge))
		}
	}
	return res
}
