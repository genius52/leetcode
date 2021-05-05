package diagram

func judgeCircle(moves string) bool {
	var x_cnt int = 0
	var y_cnt int = 0
	for _,m := range moves{
		if m == 'R'{
			x_cnt++
		}else if m == 'L'{
			x_cnt--
		}else if m == 'U'{
			y_cnt++
		}else if m == 'D'{
			y_cnt--
		}
	}
	return x_cnt == 0 && y_cnt == 0
}