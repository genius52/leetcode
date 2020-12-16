package array

//Input: hours = [9,9,6,0,6,6,9]
//Output: 3
//Explanation: The longest well-performing interval is [9,9,6].
func LongestWPI(hours []int) int {
	var l int = len(hours)
	var record map[int]int = make(map[int]int)
	var great_cnt int = 0
	var res int = 0
	for i := 0;i < l;i++{
		if hours[i] > 8{
			great_cnt++
		}else{
			great_cnt--
		}
		if _,ok := record[great_cnt];!ok{
			record[great_cnt] = i
		}
		if great_cnt > 0{
			res = i + 1
		}else{
			if pre_pos,ok := record[great_cnt - 1];ok{
				if i - pre_pos > res{
					res = i - pre_pos
				}
			}
		}
	}
	return res
}