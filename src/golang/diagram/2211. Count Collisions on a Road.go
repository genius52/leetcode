package diagram

//directions := "RLRSLL"
func CountCollisions(directions string) int {
	var l int = len(directions)
	var right_cnt int = 0
	var s_cnt int = 0
	var total int = 0
	for i := 0;i < l;i++{
		if directions[i] == 'R'{
			right_cnt++
			s_cnt = 0
		}else if directions[i] == 'S'{
			s_cnt = 1
			total += right_cnt
			right_cnt = 0
		}else if directions[i] == 'L'{
			if right_cnt > 0{
				total += 2
				total += right_cnt - 1
				s_cnt = 1
				right_cnt = 0
			}else if s_cnt > 0{
				total++
				//s_cnt++
			}
		}
	}
	return total
}