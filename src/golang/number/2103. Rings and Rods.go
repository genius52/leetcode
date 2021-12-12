package number

func countPoints(rings string) int {
	var l int = len(rings)
	var record [10][3]bool
	var res int = 0
	for i := 0; i < l; i += 2 {
		n := int(rings[i+1] - '0')
		color := 0
		if rings[i] == 'G' {
			color = 1
		} else if rings[i] == 'B' {
			color = 2
		}
		record[n][color] = true
	}
	for i := 0; i < 10; i++ {
		if record[i][0] && record[i][1] && record[i][2] {
			res++
		}
	}
	return res
}
