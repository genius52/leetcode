package number

func BestClosingTime(customers string) int {
	var l int = len(customers)
	var prefix_y []int = make([]int,l + 1) //prefix_y[i], customers[i - 1]之前y的个数
	for i := 0;i < l;i++{
		if customers[i] == 'Y'{
			prefix_y[i + 1] = 1 + prefix_y[i]
		}else{
			prefix_y[i + 1] = prefix_y[i]
		}
	}
	var min_cost int = 1 << 31 - 1
	var idx int = -1
	for i := 0;i <= l;i++{
		//在开门期间，如果某一个小时没有顾客到达，代价增加 1 。
		//在关门期间，如果某一个小时有顾客到达，代价增加 1 。
		//在i关门，0 -- (i - 1）开门
		cur_cost := i - prefix_y[i] + prefix_y[l] - prefix_y[i]
		if cur_cost < min_cost{
			min_cost = cur_cost
			idx = i
		}
	}
	return idx
}