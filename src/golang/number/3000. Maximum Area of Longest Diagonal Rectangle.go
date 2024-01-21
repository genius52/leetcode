package number

func areaOfMaxDiagonal(dimensions [][]int) int {
	var max_area int = 0
	var max_line int = 0
	for _, d := range dimensions {
		line := d[0]*d[0] + d[1]*d[1]
		if line > max_line {
			max_line = line
			max_area = d[0] * d[1]
		} else if line == max_line {
			area := d[0] * d[1]
			if area > max_area {
				max_area = area
			}
		}
	}
	return max_area
}
