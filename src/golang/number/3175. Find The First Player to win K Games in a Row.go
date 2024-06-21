package number

func FindWinningPlayer(skills []int, k int) int {
	var l int = len(skills)
	if k >= l {
		var max_player int = -1
		var max_skill int = 0
		for i := 0; i < l; i++ {
			if skills[i] > max_skill {
				max_skill = skills[i]
				max_player = i
			}
		}
		return max_player
	}
	var max_skill int = skills[0]
	var win_idx int = 0
	var win_cnt int = 0
	for i := 1; i < l; i++ {
		if skills[i] > max_skill {
			win_cnt = 1
			max_skill = skills[i]
			win_idx = i
		} else {
			win_cnt++
		}
		if win_cnt >= k {
			return win_idx
		}
	}
	return win_idx
	//var q [][2]int = make([][2]int, l)
	//for i := 0; i < l; i++ {
	//	q[i] = [2]int{skills[i], i}
	//}
	//var win_cnt int = 0
	//for true {
	//	if q[0][0] > q[1][0] {
	//		win_cnt++
	//		if win_cnt >= k {
	//			return q[0][1]
	//		}
	//		q = append(q[:1], append(q[2:], q[1])...)
	//	} else {
	//		win_cnt = 1
	//		if win_cnt >= k {
	//			return q[0][1]
	//		}
	//		q = append(q[1:], q[0])
	//	}
	//}
	return -1
}
