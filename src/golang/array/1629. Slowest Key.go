package array

//输入：releaseTimes = [9,29,49,50], keysPressed = "cbcd"
//输出："c"
func slowestKey(releaseTimes []int, keysPressed string) byte {
	var cnt [26]int
	var l int = len(releaseTimes)
	pre_time := 0
	var max_val int = 0
	for i := 0;i < l;i++{
		cur_cnt := releaseTimes[i] - pre_time
		if cur_cnt > cnt[keysPressed[i] - 'a']{
			cnt[keysPressed[i] - 'a'] = cur_cnt
		}
		if cur_cnt > max_val{
			max_val = cur_cnt
		}
		pre_time = releaseTimes[i]
	}
	var res byte
	for i := 25;i >= 0;i--{
		if cnt[i] == max_val{
			res = byte(i) + byte('a')
			break
		}
	}
	return res
}