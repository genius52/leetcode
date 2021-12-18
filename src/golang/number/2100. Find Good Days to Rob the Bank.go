package number

//输入：security = [5,3,3,3,5,6,2], time = 2
//输出：[2,3]
//解释：
//第 2 天，我们有 security[0] >= security[1] >= security[2] <= security[3] <= security[4] 。
//第 3 天，我们有 security[1] >= security[2] >= security[3] <= security[4] <= security[5] 。
//没有其他日子符合这个条件，所以日子 2 和 3 是适合打劫银行的日子。
//O(n * 3)
func goodDaysToRobBank(security []int, time int) []int {
	var l int = len(security)
	var prefix_decrease []int = make([]int, l+1)
	for i := 1; i < l; i++ {
		if security[i] <= security[i-1] {
			prefix_decrease[i] = 1 + prefix_decrease[i-1]
		}
	}
	var suffix_increase []int = make([]int, l+1)
	for i := l - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			suffix_increase[i] = 1 + suffix_increase[i+1]
		}
	}
	var res []int
	for i := 0; i < l; i++ {
		if prefix_decrease[i] >= time && suffix_increase[i] >= time {
			res = append(res, i)
		}
	}
	return res
}

func GoodDaysToRobBank(security []int, time int) []int {
	var res []int
	var l int = len(security)
	if time*2 >= l {
		return res
	}
	if time == 0 {
		for i := 0; i < l; i++ {
			res = append(res, i)
		}
		return res
	}
	var pre_decrease_cnt int = 0
	for i := 1; i <= time; i++ {
		if security[i] <= security[i-1] {
			pre_decrease_cnt++
		}
	}
	var post_increase_cnt int = 0
	for i := time; i < time*2; i++ {
		if security[i] <= security[i+1] {
			post_increase_cnt++
		}
	}
	var idx int = time
	var left int = 0
	var right int = time * 2
	for right < l {
		if pre_decrease_cnt >= time && post_increase_cnt >= time {
			res = append(res, idx)
		}
		if right+1 >= l {
			break
		}
		if security[left] >= security[left+1] {
			pre_decrease_cnt--
		}
		if security[idx] >= security[idx+1] {
			pre_decrease_cnt++
		}
		if security[idx] <= security[idx+1] {
			post_increase_cnt--
		}
		if security[right+1] >= security[right] {
			post_increase_cnt++
		}
		left++
		right++
		idx++
	}
	return res
}
