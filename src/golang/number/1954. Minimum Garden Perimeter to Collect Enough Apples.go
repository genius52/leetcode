package number

func MinimumPerimeter(neededApples int64) int64 {
	var l int64 = 2
	var pre int64 = 12
	for pre < neededApples{
		l += 2
		add := 4 *((l/2 + l/2 + l/2 - 1) * l/4 + (l/2 + 1 + l/2 + l/2) * l/4)
		pre += add
	}
	return l * 4
}