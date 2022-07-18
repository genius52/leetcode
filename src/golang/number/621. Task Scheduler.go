package number

func leastInterval(tasks []byte, n int) int {
	var cnt [26]int
	for _,t := range tasks{
		cnt[t - 'A']++
	}
	var max_cnt int = 0
	for i := 0;i < 26;i++{
		if cnt[i] > 0{
			if cnt[i] > max_cnt{
				max_cnt = cnt[i]
			}
		}
	}
	var times int = 0
	for i := 0;i < 26;i++{
		if cnt[i] == max_cnt{
			times++
		}
	}
	var l int = len(tasks)
	return max_int(l,(max_cnt - 1) * (n + 1) + times)
}