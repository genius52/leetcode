package number

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	if sx == fx && sy == fy {
		return (t % 2) == 0
	}
	x_dis := abs_int(sx - fx)
	y_dis := abs_int(sy - fy)
	total_steps := max_int(x_dis, y_dis)
	if t < total_steps {
		return false
	}
	return true
}
