package array


func MaximizeWin(prizePositions []int, k int) int{
	var l int = len(prizePositions)
	var left_max_cnt []int = make([]int,l + 1)//left_max_cnt[i] = i的左边最多能覆盖多少奖品
	var res int = 0
	var left int = 0
	var right int = 0
	for left < l && right < l{
		for left < l && prizePositions[right] - prizePositions[left] > k{//第二条线段
			left++
		}
		cur_gap := right - left + 1
		left_max_cnt[right + 1] = max_int(left_max_cnt[right],cur_gap)
		res = max_int(res,left_max_cnt[left] + cur_gap)
		right++
	}
	return res
}