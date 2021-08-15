package number

func lastStoneWeightII(stones []int) int {
	var l int = len(stones)
	if l == 1{
		return stones[0]
	}
	var sum int = 0
	for i := 0;i < l;i++{
		sum += stones[i]
	}
	var res int = 2147483647
	var record []bool = make([]bool,sum + 1)
	var cur_sum int = 0
	for i := 0;i < l;i++{
		cur_sum += stones[i]
		for j := 0;j < cur_sum;j++{
			if !record[j]{
				continue
			}
			record[cur_sum - j] = true
			diff1 := sum - (cur_sum - j) * 2
			if diff1 >= 0{
				res = min_int(res,diff1)
				if res == 0{
					return 0
				}
			}
		}
		record[cur_sum] = true
		diff2 := sum - cur_sum * 2
		if diff2 >= 0{
			res = min_int(res,diff2)
			if res == 0{
				return 0
			}
		}
	}
	return res
}