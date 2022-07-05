package string_issue

func max_occur(cnt [26]int,s string,left int,right int)int{
	var max_cnt int = 0
	for i := left;i <= right;i++{
		if cnt[s[i] - 'A'] > max_cnt{
			max_cnt = cnt[s[i] - 'A']
		}
	}
	return max_cnt
}

func CharacterReplacement(s string, k int) int {
	var l int = len(s)
	var max_len int = 0
	var left int = 0
	var right int = 0
	var cnt [26]int
	for left < l && right < l{
		cnt[s[right] - 'A']++
		if right - left + 1 <= k{
			if right - left + 1 > max_len{
				max_len = right - left + 1
			}
			right++
			continue
		}
		max_occur_cnt := max_occur(cnt,s,left,right)
		if right - left + 1 - max_occur_cnt <= k{
			var cur_len int = right - left + 1
			if cur_len > max_len{
				max_len = cur_len
			}
		}else{
			cnt[s[left] - 'A']--
			left++
		}
		right++
	}
	return max_len
}