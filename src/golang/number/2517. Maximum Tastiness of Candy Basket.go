package number

import "sort"

func MaximumTastiness(price []int, k int) int {
	var l int = len(price)
	sort.Ints(price)
	var low int = 0
	var high int = price[l - 1] - price[0]
	var res int = low
	for low <= high{
		var mid int = (low + high)/2 //假设甜蜜度是mid，计算有多少类
		var cnt int = 1
		var pre int = price[0]
		var meet bool = false
		for i := 1;i < l;i++{
			if price[i] - pre >= mid{
				cnt++
				if cnt >= k{
					meet = true
					break
				}
				pre = price[i]
			}
		}
		if meet{
			low = mid + 1
			res = mid
		}else{
			high = mid - 1
		}
	}
	return res
}