package array

func maxDistance(colors []int) int {
	var l int = len(colors)
	var res int = 0
	for i := 0;i < l;i++{
		if colors[i] != colors[l - 1]{
			res = l - i
			break
		}
	}
	for i := l - 1;i > 0;i--{
		if colors[0] != colors[i]{
			if i > res{
				res = i
			}
			break
		}
	}
	return res
}