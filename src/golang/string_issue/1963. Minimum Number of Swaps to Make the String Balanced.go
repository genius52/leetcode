package string_issue

func minSwaps1963(s string) int {
	var l int = len(s)
	var left_cnt int = 0
	var right_cnt int = 0
	var cnt int = 0
	for i := 0;i < l;i++{
		if s[i] == '['{
			left_cnt++
		}else if s[i] == ']'{
			right_cnt++
		}
		if right_cnt > left_cnt{
			cnt++
			left_cnt++
			right_cnt--
		}
	}
	return cnt
}