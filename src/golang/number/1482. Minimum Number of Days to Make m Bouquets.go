package number

//func check_minDays(bloomDay []int,l int,mid int,m int,k int,status []bool)bool{
//	var cnt int = 0
//	for i := 0;i < l;i++{
//		if bloomDay[i] <= mid{
//			status[i] = true
//			cnt++
//		}else{
//			status[i] = false
//		}
//	}
//	if cnt < m * k{
//		return false
//	}
//	var left int = 0
//	for left < l{
//
//	}
//}

func minDays(bloomDay []int, m int, k int) int {
	var l int = len(bloomDay)
	if l < m * k{
		return -1
	}
	var left int = 0
	var right int = 0
	for i := 0;i < l;i++{
		if bloomDay[i] > right{
			right = bloomDay[i]
		}
	}
	for left < right{
		mid := (left + right)/2
		flowers := 0
		cnt := 0
		for i := 0;i < l;i++{
			if bloomDay[i] <= mid{
				flowers++
				if flowers == k{
					cnt++
					flowers = 0
					if cnt >= m{
						break
					}
				}
			}else{
				flowers = 0
			}
		}
		if cnt >= m{
			right = mid
		}else{
			left = mid + 1
		}
	}
	return left
}