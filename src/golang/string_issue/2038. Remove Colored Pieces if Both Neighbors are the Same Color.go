package string_issue

func winnerOfGame2(colors string) bool{
	var l int = len(colors)
	var a_cnt int = 0
	for i := 1;i < l - 1;i++{
		if colors[i - 1] == 'A' && colors[i] == 'A' && colors[i + 1] == 'A'{
			a_cnt++
		}
		if colors[i - 1] == 'B' && colors[i] == 'B' && colors[i + 1] == 'B'{
			a_cnt--
		}
	}
	return a_cnt > 0
}

func winnerOfGame(colors string) bool {
	var l int = len(colors)
	var a_cnt int = 0
	var b_cnt int = 0
	var left int = 0
	for left < l{
		var right int = left + 1
		for right < l && colors[right] == colors[left]{
			right++
		}
		cur_len := right - left
		if cur_len >= 3{
			if colors[left] == 'A'{
				a_cnt += cur_len - 2
			}else{
				b_cnt += cur_len - 2
			}
		}
		left = right
	}
	return a_cnt > b_cnt
}