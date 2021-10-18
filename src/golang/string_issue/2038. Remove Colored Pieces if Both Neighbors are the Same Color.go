package string_issue

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