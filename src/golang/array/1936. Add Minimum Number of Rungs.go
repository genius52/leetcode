package array

func AddRungs(rungs []int, dist int) int {
	var l int = len(rungs)
	var cur int = 0
	var end int = rungs[l - 1]
	var res int = 0
	for i := 0;i < l - 1;{
		if (cur + dist) >= end{
			break
		}
		if (cur + dist) < rungs[i]{
			res++
			cur = cur + dist
			if cur >= rungs[i]{
				i++
			}
		}else{
			cur = rungs[i]
			i++
		}
	}
	if (cur + dist) < end{
		res += (end - cur) / dist
		if (end - cur) % dist == 0{
			res--
		}
	}
	return res
}