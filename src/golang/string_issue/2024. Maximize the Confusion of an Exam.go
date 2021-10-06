package string_issue

func MaxConsecutiveAnswers(answerKey string, k int) int{
	var l int = len(answerKey)
	var t_cnt int = 0
	var f_cnt int = 0
	var res int = 0
	var left int = 0
	var right int = 0
	for left < l && right < l{
		for right < l && (t_cnt <= k || f_cnt <= k){
			if answerKey[right] == 'T'{
				t_cnt++
			}else{
				f_cnt++
			}
			right++
		}
		if t_cnt > k && f_cnt > k{
			res = max_int(res,right - left - 1)
		}else{
			res = max_int(res,right - left)
		}
		if answerKey[left] == 'T'{
			t_cnt--
		}else{
			f_cnt--
		}
		left++
	}
	return res
}