package number

func lemonadeChange(bills []int) bool{
	var five_cnt int = 0
	var ten_cnt int = 0
	var l int = len(bills)
	for i := 0;i < l;i++{
		if bills[i] == 5{
			five_cnt++
		}else if bills[i] == 10{
			if five_cnt == 0{
				return false
			}
			five_cnt--
			ten_cnt++
		}else if bills[i] == 20{
			if ten_cnt > 0 {
				if five_cnt == 0{
					return false
				}
				ten_cnt--
				five_cnt--
			}else if five_cnt >= 3{
				five_cnt -= 3
			}else{
				return false
			}
		}
	}
	return true
}