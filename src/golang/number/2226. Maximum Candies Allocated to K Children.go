package number

func check_maximumCandies(candies []int,l int,each_cnt int, k int64)bool{
	var people_cnt int64 = 0
	for i := 0;i < l;i++{
		people_cnt += int64(candies[i] / each_cnt)
	}
	return people_cnt >= k
}

func MaximumCandies(candies []int, k int64) int {
	var l int = len(candies)
	var low int = 0
	var high int = 0
	var sum int = 0
	for i := 0;i < l;i++{
		if candies[i] > high{
			high = candies[i]
		}
		sum += candies[i]
	}
	if int64(sum) < k{
		return 0
	}
	var res int = low
	for low <= high{
		mid := (low + high)/2
		if mid == 0{
			mid++
		}
		ret := check_maximumCandies(candies,l,mid,k)
		if ret{
			low = mid + 1
			res = mid
		}else{
			high = mid - 1
			res = mid - 1
		}
	}
	return res
}