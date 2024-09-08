package array

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	x1 := coordinate1[0] - 'a'
	y1 := coordinate1[1] - '1'
	x2 := coordinate2[0] - 'a'
	y2 := coordinate2[1] - '1'
	return ((x1 + y1) & 1) == ((x2 + y2) & 1)
}
