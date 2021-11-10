package diagram

//Input: s = "aabaa", cost = [1,2,3,4,1]
//Output: 2
func minCost(s string, cost []int) int{
	var l int = len(s)
	var res int = 0
	last_cost := cost[0]
	for i := 1;i < l;i++{
		if s[i] == s[i - 1]{
			if cost[i] > last_cost{
				res += last_cost
				last_cost = cost[i]
			}else{
				res += cost[i]
			}
		}else{
			last_cost = cost[i]
		}
	}
	return res
}

func MinCost(s string, cost []int) int {
	var l int = len(s)
	if l <= 1{
		return 0
	}
	var res int = 0
	var min_dup int = cost[0]
	var visit int = 1
	for visit < l{
		if s[visit] == s[visit - 1]{
			if cost[visit] > min_dup{
				res += min_dup
				min_dup = cost[visit]
			}else{
				res += cost[visit]
			}
		}else{
			min_dup = cost[visit]
		}
		visit++
	}
	return res
}