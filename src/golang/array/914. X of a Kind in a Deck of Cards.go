package array

import "math"

//914. X of a Kind in a Deck of Cards
func hasGroupsSizeX(deck []int) bool {
	var record map[int]int = make(map[int]int)
	for _,num := range deck{
		if _,ok := record[num];ok{
			record[num]++
		}else{
			record[num] = 1
		}
	}

	var min_cnt int = math.MaxInt32
	var max_cnt int = 0
	for _,v := range record{
		if v > max_cnt{
			max_cnt = v
		}
		if v < min_cnt{
			min_cnt = v
		}
	}
	if min_cnt < 2{
		return false
	}
	min_cnt = 2
	for ;min_cnt <= max_cnt;min_cnt++{
		var match bool = true
		for _,v := range record{
			if v % min_cnt != 0{
				match = false
				break
			}
		}
		if match{
			return true
		}
	}
	return false
}