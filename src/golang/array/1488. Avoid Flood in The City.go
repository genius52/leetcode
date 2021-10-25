package array

func AvoidFlood(rains []int) []int{
	var record map[int]int = make(map[int]int)
	var free []int
	var l int = len(rains)
	var res []int = make([]int,l)
	for i := 0;i < l;i++{
		if rains[i] == 0{
			free = append(free,i)
			res[i] = 1
		}else{
			if pre_idx,ok := record[rains[i]];ok{
				if len(free) == 0{
					return []int{}
				}else{
					var visit int = len(free) - 1
					for visit >= 0 && free[visit] > pre_idx{
						visit--
					}
					if visit == len(free) - 1{
						return []int{}
					}
					free_idx := free[visit + 1]
					free = append(free[:visit + 1],free[visit + 2:]...)
					record[rains[i]] = i
					res[free_idx] = rains[i]
					res[i] = -1
				}
			}else{
				record[rains[i]] = i
				res[i] = -1
			}
		}
	}
	return res
}