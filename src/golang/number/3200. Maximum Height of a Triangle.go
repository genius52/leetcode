package number

func MaxHeightOfTriangle(red int, blue int) int {
	var n1 int = red
	var n2 int = blue
	var levels int = 1
	var use1 bool = true
	for n1 >= 0 && n2 >= 0 {
		if use1 {
			n1 -= levels
			if n2 < levels+1 {
				break
			}
		} else {
			n2 -= levels
			if n1 < levels+1 {
				break
			}
		}
		if n1 == 0 && n2 == 0 {
			break
		}
		use1 = !use1
		levels++
	}
	var res int = levels
	levels = 1
	n1 = blue
	n2 = red
	use1 = true
	for n1 >= 0 && n2 >= 0 {
		if use1 {
			n1 -= levels
			if n2 < levels+1 {
				break
			}
		} else {
			n2 -= levels
			if n1 < levels+1 {
				break
			}
		}
		if n1 == 0 && n2 == 0 {
			break
		}
		use1 = !use1
		levels++
	}
	res = max_int(res, levels)
	return res
}
