package number


//1,2,2,3,4
//1,2,3,4,5
//nums = [1,3,4,2,2]
func FindDuplicate(nums []int) int{
	var l int = len(nums)
	var res int = 0
	for i := 0;i < 32;i++{
		var index_one_cnt int = 0
		var number_one_cnt int = 0
		var mask int = 1 << i
		for i := 0;i < l;i++{
			if (i & mask) != 0{
				index_one_cnt++
			}
			if (nums[i] & mask) != 0{
				number_one_cnt++
			}
		}
		if number_one_cnt > index_one_cnt{
			res |= mask
		}
	}
	return res
}


//287
func findDuplicate(nums []int) int {
	min := 0
	max := len(nums)-1
	for min < max{
		mid := min + (max - min)/2
		cnt := 0
		for _,v := range nums{
			if v <= mid{
				cnt++
			}
		}
		if cnt <= mid{
			min = mid + 1
		}else{
			max = mid
		}
	}
	return min
}