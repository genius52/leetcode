package number

func maxCount(banned []int, n int, maxSum int) int {
	var cnt int = 0
	var ban map[int]bool = make(map[int]bool)
	for _,n := range banned{
		ban[n] = true
	}
	var cur_sum int = 0
	for i := 1;i <= n;i++{
		if cur_sum + i > maxSum{
			break
		}
		if _,ok := ban[i];ok{
			continue
		}
		cur_sum += i
		cnt++
	}
	return cnt
}