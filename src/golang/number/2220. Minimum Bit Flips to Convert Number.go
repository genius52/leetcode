package number

func minBitFlips(start int, goal int) int {
	var res int = 0
	for start > 0 && goal > 0{
		if ((start | 1) == start && (goal | 1) != goal) || ((start | 1) != start && (goal | 1) == goal){
			res++
		}
		start >>= 1
		goal >>= 1
	}
	for start > 0{
		if (start | 1) == start{
			res++
		}
		start >>= 1
	}
	for goal > 0{
		if (goal | 1) == goal{
			res++
		}
		goal >>= 1
	}
	return res
}