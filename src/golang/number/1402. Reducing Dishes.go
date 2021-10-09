package number

import "sort"

//输入：satisfaction = [-1,-8,0,5,-9]
//输出：14
//解释：去掉第二道和最后一道菜，最大的喜爱时间系数和为 (-1*1 + 0*2 + 5*3 = 14) 。每道菜都需要花费 1 单位时间完成。
func MaxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	var l int = len(satisfaction)
	var pre_sum int = 0
	var cur_sum int = 0
	single_sum := 0
	for i := l - 1;i >= 0;i--{
		single_sum += satisfaction[i]
		cur_sum = single_sum + pre_sum
		if cur_sum < pre_sum{
			return pre_sum
		}else{
			pre_sum = cur_sum
		}
	}
	return pre_sum
}