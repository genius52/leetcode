package array

func numTeams(rating []int) int {
	var l int = len(rating)
	var dp_less []int = make([]int,l)
	var dp_bigger []int = make([]int,l)
	for i := l - 2;i >= 0;i--{
		var less_cnt int = 0
		var bigger_cnt int = 0
		for j := i + 1;j < l;j++{
			if rating[i] < rating[j]{
				less_cnt++
			}
			if rating[i] > rating[j]{
				bigger_cnt++
			}
		}
		dp_less[i] = less_cnt
		dp_bigger[i] = bigger_cnt
	}
	var res int = 0
	for i := 0;i < l - 2;i++{
		for j := i + 1;j < l - 1;j++{
			if rating[i] < rating[j]{
				res += dp_less[j]
			}
			if rating[i] > rating[j]{
				res += dp_bigger[j]
			}
		}
	}
	return res
}