package number

func check_minimumTime(time []int,total_time int64,totalTrips int64)bool{
	var trips_cnt int64 = 0
	for _,t := range time{
		trips_cnt += total_time / int64(t)
		if trips_cnt >= totalTrips{
			return true
		}
	}
	return false
}

func MinimumTime(time []int, totalTrips int) int64 {
	var min_t int = 2147483647
	for _,t := range time{
		if t < min_t{
			min_t = t
		}
	}
	var left int64 = 1
	var right int64 = int64(min_t) * int64(totalTrips)
	var res int64 = 1
	for left < right{
		mid := (left + right)/2
		ret := check_minimumTime(time,mid,int64(totalTrips))
		if ret{
			right = mid
			res = mid
		}else{
			left = mid + 1
			res = mid + 1
		}
	}
	return res
}